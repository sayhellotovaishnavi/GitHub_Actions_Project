package main

import (
    "io/ioutil"
    "net/http/httptest"
    "strconv"
    "testing"
)

func TestAdditionHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/addition?a=3&b=5", nil)
    w := httptest.NewRecorder()
    additionHandler(w, req)
    resp := w.Result()
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    result, _ := strconv.Atoi(string(body))
    expected := 8

    if result != expected {
        t.Errorf("Expected %d, but got %d", expected, result)
    }
}
