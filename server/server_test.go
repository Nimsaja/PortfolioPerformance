package main

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/Nimsaja/PortfolioPerformance/store"
// )

// var (
// 	server = httptest.NewServer(handler())
// )

// func init() {
// 	os.Setenv("RUN_IN_CLOUD", "NotSet")
// }

// func TestForceCall(t *testing.T) {
// 	r := httptest.NewRequest("GET", "http://localhost:8080/portfolio/forcecall", nil)
// 	w := httptest.NewRecorder()
// 	loadHistData(w, r)

// 	// check status code
// 	resp := w.Result()
// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("expected status ok (200), but is: %v", resp.StatusCode)
// 	}

// 	// check result
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	s := ""
// 	err := json.Unmarshal(body, &s)
// 	if err != nil {
// 		t.Errorf("No err expected: %v", err)
// 	}

// 	if !strings.HasPrefix(s, "Successfully") {
// 		t.Errorf("Expected a success string, got %v", s)
// 	}
// }

// func TestGetTableData(t *testing.T) {
// 	r := httptest.NewRequest("GET", "http://localhost:8080/portfolio/table", nil)
// 	w := httptest.NewRecorder()
// 	getTableData(w, r)

// 	// check status code
// 	resp := w.Result()
// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("expected status ok (200), but is: %v", resp.StatusCode)
// 	}

// 	// results
// 	data := make([]store.Data, 0)
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	err := json.Unmarshal(body, &data)
// 	if err != nil {
// 		t.Errorf("No err expected: %v", err)
// 	}

// 	// check size of data
// 	if len(data) == 0 {
// 		t.Errorf("data size expected to be not empty, but get: %v", len(data))
// 	}

// 	//check if we can get some data
// 	l := len(data)
// 	tStart := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
// 	if !time.Unix(int64(data[0].Time), 0).After(tStart) {
// 		t.Errorf("Expected saved time to be newer than %v, got %v", 2017, time.Unix(int64(data[0].Time), 0))
// 	}
// 	if !data[l-1].TimeHuman.After(tStart) {
// 		t.Errorf("Expected saved time to be newer than %v, got %v", 2017, data[l-1])
// 	}
// 	if data[0].Value == 0 {
// 		t.Errorf("Expected to get some values for data value, got %v", data[0].Value)
// 	}
// 	if data[l-1].Diff == 0 {
// 		t.Errorf("Expected to get some values for data diff, got %v", data[l-1].Value)
// 	}
// }
