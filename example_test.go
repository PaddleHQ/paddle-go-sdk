package paddle_test

import (
	"embed"
	"net/http"
	"net/http/httptest"
)

type stubPath string

const (
	event_types              stubPath = "testdata/event_types.json"
	events                   stubPath = "testdata/events.json"
	transaction              stubPath = "testdata/transaction.json"
	transactions             stubPath = "testdata/transactions.json"
	transactionsPaginatedPg1 stubPath = "testdata/transactions_paginated_pg1.json"
	transactionsPaginatedPg2 stubPath = "testdata/transactions_paginated_pg2.json"
	priceCreatedEvent        stubPath = "testdata/price_created.json"
	simulation               stubPath = "testdata/simulation.json"
	simulations              stubPath = "testdata/simulations.json"
	simulationRun            stubPath = "testdata/simulation_run.json"
	simulationRunWithEvents  stubPath = "testdata/simulation_run_with_events.json"
)

//go:embed testdata
var exampleData embed.FS

type stub struct {
	paths []stubPath
}

type mockServerResponse struct {
	stub *stub
	body *[]byte
}

func mockServerForExample(res mockServerResponse) *httptest.Server {
	call := 0
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var body []byte

		if res.stub != nil && len(res.stub.paths) > 0 {
			stub := res.stub.paths[0]
			if len(res.stub.paths) > 1 {
				stub = res.stub.paths[call]
			}
			c, err := exampleData.ReadFile(string(stub))
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			body = c
		}

		if res.body != nil {
			body = *res.body
		}

		w.Write(body)
		call++
	}))

	return s
}
