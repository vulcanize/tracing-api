package test

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/rpc"
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

func getGraphFrames(hash string) ([]frame, error) {
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
	if err := db.Select(&tmp, `SELECT "id", "from", "to", "input", "output" FROM sgd1.frame WHERE "hash"=$1`, hash); err != nil {
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

func testTxTraceGraph(hash string) func(t *testing.T) {
	return func(t *testing.T) {
		calls, err := callTracingAPI("http://127.0.0.1:8083", hash)
		if err != nil {
			t.Fatal(err)
			return
		}

		graphFrames, err := getGraphFrames(hash)
		if err != nil {
			t.Fatal(err)
			return
		}

		t.Log("--------------------------------------")
		t.Logf("hash: %s", hash)
		t.Log("--------------------------------------")
		t.Logf("TracingAPI: %+v", calls)
		t.Log("--------------------------------------")
		t.Logf("TheGraph: %+v", graphFrames)
		t.Log("--------------------------------------")

		if len(calls.Frames) == 0 {
			t.Error("tracing-api frames are empty")
			return
		}

		if len(graphFrames) == 0 {
			t.Error("thegraph frames are empty")
			return
		}

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
				t.Fatalf("Method [%s] doesn't exist", graphFrame.id)
			}
			frame, exist := sig2frame[sig]
			if !exist {
				t.Fatalf("Method [%s] doesn't exist in frame map", graphFrame.id)
			}

			t.Log("Compare frames")
			t.Logf("TracingAPI frame: %+v", frame)
			t.Logf("TheGraph frame: %+v", graphFrame)

			t.Logf("  Compare 'from': TheGraph[%s], TracingAPI[%s]", graphFrame.from.Hex(), frame.From.Hex())
			if !bytes.Equal(graphFrame.from.Bytes(), frame.From.Bytes()) {
				t.Fatalf("Bad 'FROM'. Want: %s, Got: %s", graphFrame.from.Hex(), frame.From.Hex())
			}

			t.Logf("  Compare 'to': TheGraph[%s], TracingAPI[%s]", graphFrame.to.Hex(), frame.To.Hex())
			if !bytes.Equal(frame.To.Bytes(), graphFrame.to.Bytes()) {
				t.Fatalf("Bad 'TO'. Want: %s, Got: %s", graphFrame.to.Hex(), frame.To.Hex())
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

				t.Logf("  Compare inputs: TheGraph[%+v], TracingAPI[%+v]", graphInputs, callInputs)
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

				t.Logf("  Compare outputs: TheGraph[%+v], TracingAPI[%+v]", graphOutputs, callOutputs)
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

				t.Logf("  Compare inputs: TheGraph[%+v], TracingAPI[%+v]", graphInputs, callInputs)
				if cmp, err := compareKeyFromMapAsStr("key", callInputs, graphInputs); err != nil || cmp != 0 {
					t.Errorf("Bad values in input args: [%#v][%#v] %v", callInputs["key"], graphInputs["key"], err)
				}

				if cmp, err := compareKeyFromMapAsStr("_value", callInputs, graphInputs); err != nil || cmp != 0 {
					t.Errorf("Bad values in input args: [%#v][%#v] %v", callInputs["key"], graphInputs["key"], err)
				}
			}
			t.Log("--------------------------------------")
		}
	}
}

func callDebugTraceTransaction(hash string, url string, config *eth.TraceConfig) ([]byte, error) {
	client, err := rpc.DialContext(context.Background(), url)
	var result interface{}
	configArg := map[string]interface{}{}

	if config != nil {
		if config.Tracer != nil {
			configArg["tracer"] = config.Tracer
		} else {
			configArg["disableStorage"] = config.DisableStorage
			configArg["disableMemory"] = config.DisableMemory
			configArg["disableStack"] = config.DisableStack
		}
	}

	err = client.Call(&result, "debug_traceTransaction", hash, configArg)
	if err != nil {
		return nil, err
	}

	byteResult, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return byteResult, err
}

func testTraceTransaction(hash string) func(t *testing.T) {
	return func(t *testing.T) {

		t.Logf("call tracing-api debug_traceTransaction")
		rawTracingApiData, err := callDebugTraceTransaction(hash, "http://127.0.0.1:8083", nil)
		if err != nil {
			t.Fatalf("    error: %s", err)
		}
		t.Logf("    done")

		t.Logf("call geth debug_traceTransaction")
		rawGethData, err := callDebugTraceTransaction(hash, "http://127.0.0.1:8545", nil)
		if err != nil {
			t.Fatalf("    error: %s", err)
		}
		t.Logf("    done")

		if !bytes.Equal(rawTracingApiData, rawGethData) {
			t.Error("bad tracing api data")
		}
	}
}

func testTraceTransactionWithLogConfig(hash string) func(t *testing.T) {
	return func(t *testing.T) {

		config := &eth.TraceConfig{
			LogConfig: &vm.LogConfig{
				DisableMemory:     true,
				DisableStack:      true,
				DisableStorage:    true,
				DisableReturnData: true,
			},
		}

		t.Logf("call tracing-api debug_traceTransaction")
		rawTracingApiData, err := callDebugTraceTransaction(hash, "http://127.0.0.1:8083", config)
		if err != nil {
			t.Fatalf("    error: %s", err)
		}
		t.Logf("    done")

		t.Logf("call geth debug_traceTransaction")
		rawGethData, err := callDebugTraceTransaction(hash, "http://127.0.0.1:8545", config)
		if err != nil {
			t.Fatalf("    error: %s", err)
		}
		t.Logf("    done")

		if !bytes.Equal(rawTracingApiData, rawGethData) {
			t.Error("bad tracing api data")
		}
	}
}

func testTraceTransactionWithTracer(hash string) func(t *testing.T) {
	return func(t *testing.T) {

		tracer := `{
			data: [], 
			fault: function(log) {}, 
			step: function(log) { 
				var op = log.op.toString();
				if(op == 'CALL') {
					var off = (op == 'DELEGATECALL' || op == 'STATICCALL' ? 0 : 1);

					var inOff = log.stack.peek(2 + off).valueOf();
					var inEnd = inOff + log.stack.peek(3 + off).valueOf();

					this.data.push({
						to: toHex(toAddress(log.stack.peek(1).toString(16))),
						input: toHex(log.memory.slice(inOff, inEnd)),
					}); 
				}
					
			}, 
			result: function() { return this.data; }
		}`

		config := &eth.TraceConfig{
			Tracer: &tracer,
		}

		t.Logf("call tracing-api debug_traceTransaction")
		rawTracingApiData, err := callDebugTraceTransaction(hash, "http://127.0.0.1:8083", config)
		if err != nil {
			t.Fatalf("    error: %s", err)
		}
		t.Logf("    done")

		t.Logf("call geth debug_traceTransaction")
		rawGethData, err := callDebugTraceTransaction(hash, "http://127.0.0.1:8545", config)
		if err != nil {
			t.Fatalf("    error: %s", err)
		}
		t.Logf("    done")

		if !bytes.Equal(rawTracingApiData, rawGethData) {
			t.Error("bad tracing api data")
		}
	}
}

func TestMain(t *testing.T) {
	t.Log("Generate transaction")
	hash, err := callTx("http://127.0.0.1:3000")
	if err != nil {
		t.Fatalf("  error: %s", err)
	}
	t.Logf("  hash: %s", hash)

	time.Sleep(2 * time.Second)

	t.Run("test TxTraceGraph", testTxTraceGraph(hash))
	t.Run("test TraceTransaction", testTraceTransaction(hash))
	t.Run("test TraceTransaction with params", testTraceTransactionWithLogConfig(hash))
	t.Run("test TraceTransaction with tracer", testTraceTransactionWithTracer(hash))
}
