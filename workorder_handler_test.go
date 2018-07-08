package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func getWorkordersHandlerTest(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/workorders")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	expectedResult := []Workorder{}
	var result []Workorder
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}

	if len(expectedResult) != len(result) {
		t.Errorf("Expected %v, got %v", len(expectedResult), len(result))
	}
}

func createWorkordersHandlerTest(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	workorder := Workorder{
		OBJNR:       "B450",
		CreatedAt:   time.Now().Local(),
		Description: "Testing my routes",
		Adress:      "Pilevallsvägen 18",
		Start:       "Vet inte vad detta är",
		Status:      "Vet inte vad detta är heller",
		Invoice:     "FAKTURERAT",
	}

	requestBytes, _ := json.Marshal(workorder)
	requestReader := bytes.NewReader(requestBytes)

	resp, err := http.Post(mockServer.URL+"/workorders", "application/json", requestReader)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be %d, but got %d", http.StatusOK, resp.StatusCode)
	}
}
