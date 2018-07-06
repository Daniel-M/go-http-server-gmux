package main

import (
	"encoding/json"
	"log"
	//"io"
	"io/ioutil"
	//"net/http"
	"net/http/httptest"
	//"os"
	"testing"
)

// TestGetMethod tests that the Server sends the correct response for the
// GET query
func TestGetMethod(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:5555/", nil)
	w := httptest.NewRecorder()

	HomeGET(w, req)

	res := w.Result()
	bodyByte, _ := ioutil.ReadAll(res.Body)

	var bodyJSON []Response

	err := json.Unmarshal(bodyByte, &bodyJSON)

	if err != nil {
		log.Fatal("Something happened parsing the response:", err)
	}

	if bodyJSON[0].Message != "Hello GET" {
		t.Fatal("Home Handler sent the wrong response message", string(bodyByte))
	}

	if bodyJSON[0].Code != 200 {
		t.Fatal("Home Handler sent the wrong response code", string(bodyByte))
	}

}

func TestPostMethod(t *testing.T) {
	req := httptest.NewRequest("POST", "http://localhost:5555/", nil)
	w := httptest.NewRecorder()

	HomePOST(w, req)

	res := w.Result()
	bodyByte, _ := ioutil.ReadAll(res.Body)

	var bodyJSON []Response

	err := json.Unmarshal(bodyByte, &bodyJSON)

	if err != nil {
		log.Fatal("Something happened parsing the response:", err)
	}

	if bodyJSON[0].Message != "Hello POST" {
		t.Fatal("Home Handler sent the wrong response message", string(bodyByte))
	}

	if bodyJSON[0].Code != 200 {
		t.Fatal("Home Handler sent the wrong response code", string(bodyByte))
	}

}
