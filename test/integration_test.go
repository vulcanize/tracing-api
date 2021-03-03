package test

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres driver
	"github.com/vulcanize/ipld-eth-indexer/pkg/postgres"
)

var (
	pkey *ecdsa.PrivateKey
	auth *bind.TransactOpts
)

func init() {
	pkey, _ = crypto.HexToECDSA("d91499da14f0a5f0dc3c924bc8068340e3be0466c01017f34b90cee9ab73fb36")
	auth = bind.NewKeyedTransactor(pkey)
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
func callTracingAPI(path string, hash string) (string, error) {
	resp, err := http.Post(path, "application/json", strings.NewReader(fmt.Sprintf(
		`{
			"jsonrpc": "2.0",
			"id": 0,
			"method": "debug_txTraceGraph",
			"params": ["%s"]
		}`,
		hash,
	)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
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

	frames, err := getGraphFrames()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("hash: %s", hash)
	t.Log("-------------------")
	t.Logf("Tracing: %+v", calls)
	t.Log("-------------------")
	t.Logf("DGraph: %+v", frames)
}
