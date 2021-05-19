package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	ID         string
	Responce   string
	StatusCode int
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("id")
	if key == "42" {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status":200, "resp":{"user":42}}`)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500,"err":"db_error"}`)
	}
}

func TestGetUser(t *testing.T) {
	cases := []TestCase{
		TestCase{
			ID:         "42",
			Responce:   `{"status":200, "resp":{"user":42}}`,
			StatusCode: http.StatusOK,
		},
		TestCase{
			ID:         "500",
			Responce:   `{"status":500,"err":"db_error"}`,
			StatusCode: http.StatusInternalServerError,
		},
	}

	for caseNum, item := range cases {
		url := "http://example.com/api/user?id=" + item.ID
		req := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()

		GetUser(w, req)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, excepted %d", caseNum, w.Code, item.StatusCode)
		}
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		bodyStr := string(body)
		if bodyStr != item.Responce {
			t.Errorf("[%d] wrong responce: got %+v, excepted %+v", caseNum, bodyStr, item.Responce)
		}
	}
}