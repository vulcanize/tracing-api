package test

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres driver
	"github.com/vulcanize/ipld-eth-indexer/pkg/postgres"
	"github.com/vulcanize/tracing-api/pkg/cache"
	"github.com/vulcanize/tracing-api/pkg/eth/tracer"
)

var (
	pkey    *ecdsa.PrivateKey
	auth    *bind.TransactOpts
	syncABI abi.ABI
	storABI abi.ABI
)

func init() {
	pkey, _ = crypto.HexToECDSA("d91499da14f0a5f0dc3c924bc8068340e3be0466c01017f34b90cee9ab73fb36")
	auth = bind.NewKeyedTransactor(pkey)
	syncABI, _ = abi.JSON(strings.NewReader(`[{"constant":true,"inputs":[],"name":"store","outputs":[{"internalType":"contract UintStorage","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"string","name":"key","type":"string"}],"name":"sync","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":true,"stateMutability":"payable","type":"function"}]`))
	storABI, _ = abi.JSON(strings.NewReader(`[{"constant":true,"inputs":[{"internalType":"string","name":"key","type":"string"}],"name":"get","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"string","name":"key","type":"string"},{"internalType":"uint256","name":"_value","type":"uint256"}],"name":"set","outputs":[],"payable":true,"stateMutability":"payable","type":"function"}]`))
}

// callTx create transaction, docker-compose: contracts
func callTx(path string) (string, error) {
	res, err := http.Get(path)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var hexes []string
	if err := decoder.Decode(&hexes); err != nil {
		return "", err
	}
	if len(hexes) == 0 {
		return "", errors.New("no data")
	}
	return hexes[0], nil
}

// callTracingAPI get callstack for given tx hash
func callTracingAPI(path string, hash string) (*cache.TxTraceGraph, error) {
	res, err := http.Post(path, "application/json", strings.NewReader(fmt.Sprintf(
		`{
			"jsonrpc": "2.0",
			"id": 0,
			"method": "debug_txTraceGraph",
			"params": ["%s"]
		}`,
		hash,
	)))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var tmp struct {
		Err    error               `json:"error,omitempty"`
		Result *cache.TxTraceGraph `json:"result,omitempty"`
	}
	if err := decoder.Decode(&tmp); err != nil {
		return nil, err
	}
	if tmp.Err != nil {
		return nil, err
	}

	return tmp.Result, nil
}

type frame struct {
	id     string
	from   common.Address
	to     common.Address
	input  string
	output string
}

func getGraphFrames() ([]frame, error) {
	db, err := sqlx.Connect("postgres", postgres.DbConnectionString(postgres.Config{
		Name:     "thegraph",
		Hostname: "127.0.0.1",
		Port:     5434,
		User:     "postgres",
		Password: "pwd",
	}))
	if err != nil {
		return nil, postgres.ErrDBConnectionFailed(err)
	}

	tmp := []struct {
		ID     string
		From   []byte
		To     []byte
		Input  string
		Output string
	}{}
	if err := db.Select(&tmp, `SELECT "id", "from", "to", "input", "output" FROM sgd1.frame`); err != nil {
		return nil, err
	}
	data := make([]frame, len(tmp))
	for i := range tmp {
		data[i] = frame{
			id:     tmp[i].ID,
			from:   common.BytesToAddress(tmp[i].From),
			to:     common.BytesToAddress(tmp[i].To),
			input:  tmp[i].Input,
			output: tmp[i].Output,
		}
	}

	return data, nil
}

func TestMain(t *testing.T) {
	hash, err := callTx("http://127.0.0.1:3000")
	if err != nil {
		t.Error(err)
		return
	}

	calls, err := callTracingAPI("http://127.0.0.1:8083", hash)
	if err != nil {
		t.Error(err)
		return
	}

	graphFrames, err := getGraphFrames()
	if err != nil {
		t.Error(err)
		return
	}

	if len(calls.Frames) > len(graphFrames) {
		t.Errorf("tracing-api callstack (%d) less then callstack from thegraph (%d)", len(calls.Frames), len(graphFrames))
		return
	}

	name2sig := map[string]string{
		"sync": "0xfd620be1",
		"set":  "0x8a42ebe9",
	}

	sig2frame := map[string]tracer.Frame{}
	for _, frame := range calls.Frames {
		sig := fmt.Sprintf("%#x", []byte(frame.Input)[0:4])
		sig2frame[sig] = frame
	}

	for _, graphFrame := range graphFrames {
		sig, exist := name2sig[graphFrame.id]
		if !exist {
			t.Errorf("Method [%s] doesn't exist", graphFrame.id)
			return
		}
		frame, exist := sig2frame[sig]
		if !exist {
			t.Errorf("Method [%s] doesn't exist in frame map", graphFrame.id)
			return
		}
		if graphFrame.id == "sync" {
			_, err := syncABI.MethodById([]byte(frame.Input)[0:4])
			if err != nil {
				t.Errorf("Can't find method[sync] by sig[%#x]", []byte(frame.Input)[0:4])
			}
		}
		if graphFrame.id == "set" {
			_, err := storABI.MethodById([]byte(frame.Input)[0:4])
			if err != nil {
				t.Errorf("Can't find method[set] by sig[%#x]", []byte(frame.Input)[0:4])
			}
		}
	}

	t.Logf("hash: %s", hash)
	t.Log("-------------------")
	t.Logf("Tracing: %+v", calls)
	t.Log("-------------------")
	t.Logf("DGraph: %+v", graphFrames)
}
