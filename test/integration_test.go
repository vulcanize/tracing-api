package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vulcanize/tracing-api/test/contracts"
)

const pkeyHEX = "d91499da14f0a5f0dc3c924bc8068340e3be0466c01017f34b90cee9ab73fb36"

func TestMain(t *testing.T) {
	pk, err := crypto.HexToECDSA(pkeyHEX)
	if err != nil {
		t.Error(err)
	}
	auth := bind.NewKeyedTransactor(pk)

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		t.Error(err)
	}

	addr, deployTx, contract, err := contracts.DeployContracts(auth, client)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Deployed: %s", addr.Hex())
	t.Logf("Transaction: %s", deployTx.Hash().Hex())
	time.Sleep(12 * time.Second)

	callTx, err := contract.Watch(auth)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Call: %s", callTx.Hash().Hex())
	time.Sleep(12 * time.Second)

	resp, err := http.Post("127.0.0.1:8083", "application/json", strings.NewReader(fmt.Sprintf(
		`{
      jsonrpc: "2.0",
      id: 0,
      method: "debug_txTraceGraph",
      params: ["%s"]
    }`,
		callTx.Hash().Hex(),
	)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Tracing response: %s", string(body))
}
