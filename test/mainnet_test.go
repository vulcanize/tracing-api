package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func getDataFromTheGraph() (map[string]map[string][]struct {
	ID    string
	Key   string
	Value string
}, error) {
	res, err := http.Post("https://api.thegraph.com/subgraphs/name/ramilexe/blocknumstorage", "application/json", strings.NewReader(`
		{
			"query": "{\n  syncs(first: 5) {\n    id\n    key\n    value\n  }\n  storageSets(first: 5) {\n    id\n    key\n    value\n  }\n}\n",
			"variables": null
		}
	`))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]map[string][]struct {
		ID    string
		Key   string
		Value string
	}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func TestMainnet(t *testing.T) {
	t.Log("Get data from api.thegraph.com")
	theGraphData, err := getDataFromTheGraph()
	if err != nil {
		t.Fatalf("  error: %s", err)
	}
	if theGraphData == nil || theGraphData["data"] == nil {
		t.Fatalf("  data must not be empty")
	}
	if theGraphData["data"]["storageSets"] == nil || theGraphData["data"]["syncs"] == nil {
		t.Fatalf("  'storageSets' and 'syncs' must not be empty")
	}
	t.Log("  ok")
}
