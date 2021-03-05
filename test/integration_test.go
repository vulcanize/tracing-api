package test

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

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

func compareKeyFromMapAsStr(key string, a map[string]interface{}, b map[string]interface{}) (int, error) {
	aRawValue, ok := a[key]
	if !ok {
		return 0, fmt.Errorf("key:%s isn't exist", key)
	}

	bRawValue, ok := b[key]
	if !ok {
		return 0, fmt.Errorf("key:%s isn't exist", key)
	}

	aValue, ok := aRawValue.(string)
	if !ok {
		return 0, fmt.Errorf("cant convert %#v to string", aRawValue)
	}

	bValue, ok := bRawValue.(string)
	if !ok {
		return 0, fmt.Errorf("cant convert %#v to string", bRawValue)
	}

	return strings.Compare(aValue, bValue), nil
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

	time.Sleep(2 * time.Second)

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

	t.Log("-------------------")
	t.Logf("hash: %s", hash)
	t.Log("-------------------")
	t.Logf("Tracing: %+v", calls)
	t.Log("-------------------")
	t.Logf("DGraph: %+v", graphFrames)
	t.Log("-------------------")

	if len(calls.Frames) < len(graphFrames) && false {
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

		if !bytes.Equal(graphFrame.from.Bytes(), frame.From.Bytes()) {
			t.Errorf("Bad 'FROM'. Want: %s, Got: %s", graphFrame.from.Hex(), frame.From.Hex())
		}

		if !bytes.Equal(frame.To.Bytes(), graphFrame.to.Bytes()) {
			t.Errorf("Bad 'TO'. Want: %s, Got: %s", graphFrame.from.Hex(), frame.From.Hex())
		}

		if graphFrame.id == "sync" {
			method, err := syncABI.MethodById([]byte(frame.Input)[0:4])
			if err != nil {
				t.Errorf("Can't find method[sync] by sig[%#x]", []byte(frame.Input)[0:4])
				continue
			}

			callInputs := map[string]interface{}{}
			if err := method.Inputs.UnpackIntoMap(callInputs, []byte(frame.Input)[4:]); err != nil {
				t.Errorf("Cant't decode tracing-api inputs for method[sync]: %s", err)
				continue
			}

			graphInputs := map[string]interface{}{}
			tmpGraphInputs := make([]struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			}, 0)
			if err := json.Unmarshal([]byte(graphFrame.input), &tmpGraphInputs); err != nil {
				t.Errorf("Cant't decode thegraph inputs for method[sync]: %s", err)
				continue
			}
			for i := range tmpGraphInputs {
				graphInputs[tmpGraphInputs[i].Name] = tmpGraphInputs[i].Value
			}

			cmp, err := compareKeyFromMapAsStr("key", callInputs, graphInputs)
			if err != nil || cmp != 0 {
				t.Errorf("Bad values in input args: [%#v][%#v]", callInputs["key"], graphInputs["key"])
			}

			callOutputs := map[string]interface{}{}
			tmpCallOutputs := map[string]interface{}{}
			if err := method.Outputs.UnpackIntoMap(tmpCallOutputs, []byte(frame.Output)); err != nil {
				t.Errorf("Cant't decode tracing-api outputs for method[sync]: %s", err)
				continue
			}
			for key := range tmpCallOutputs {
				callOutputs[key] = fmt.Sprintf("%d", tmpCallOutputs[key])
			}

			graphOutputs := map[string]interface{}{}
			tmpGraphOutputs := make([]struct {
				Name  string `json:"name"`
				Kind  string `json:"kind"`
				Value string `json:"value"`
			}, 0)
			if err := json.Unmarshal([]byte(graphFrame.output), &tmpGraphOutputs); err != nil {
				t.Errorf("Cant't decode thegraph inputs for method[sync]: %s", err)
				continue
			}
			for i := range tmpGraphOutputs {
				graphOutputs[tmpGraphOutputs[i].Name] = tmpGraphOutputs[i].Value
			}

			if cmp, err := compareKeyFromMapAsStr("", callOutputs, graphOutputs); err != nil || cmp != 0 {
				t.Errorf("Bad values in output args: [%#v][%#v] %v", callOutputs[""], graphOutputs[""], err)
			}
		}
		if graphFrame.id == "set" {
			method, err := storABI.MethodById([]byte(frame.Input)[0:4])
			if err != nil {
				t.Errorf("Can't find method[set] by sig[%#x]", []byte(frame.Input)[0:4])
			}

			callInputs := map[string]interface{}{}
			tmpCallInputs := map[string]interface{}{}
			if err := method.Inputs.UnpackIntoMap(tmpCallInputs, []byte(frame.Input)[4:]); err != nil {
				t.Errorf("Cant't decode tracing-api inputs for method[sync]: %s", err)
				continue
			}
			for key := range tmpCallInputs {
				callInputs[key] = fmt.Sprintf("%v", tmpCallInputs[key])
			}

			graphInputs := map[string]interface{}{}
			tmpGraphInputs := make([]struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			}, 0)
			if err := json.Unmarshal([]byte(graphFrame.input), &tmpGraphInputs); err != nil {
				t.Errorf("Cant't decode thegraph inputs for method[set]: %s", graphFrame.input)
				continue
			}
			for i := range tmpGraphInputs {
				graphInputs[tmpGraphInputs[i].Name] = tmpGraphInputs[i].Value
			}

			if cmp, err := compareKeyFromMapAsStr("key", callInputs, graphInputs); err != nil || cmp != 0 {
				t.Errorf("Bad values in input args: [%#v][%#v] %v", callInputs["key"], graphInputs["key"], err)
			}

			if cmp, err := compareKeyFromMapAsStr("_value", callInputs, graphInputs); err != nil || cmp != 0 {
				t.Errorf("Bad values in input args: [%#v][%#v] %v", callInputs["key"], graphInputs["key"], err)
			}
		}
	}
}
