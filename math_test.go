package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestAdditionHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/add?a=5&b=7", nil)
	w := httptest.NewRecorder()

	AdditionHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Beklenen status 200, ama geldi: %d", resp.StatusCode)
	}

	expected := `{"result": 12}`
	buf := make([]byte, len(expected))
	_, err := resp.Body.Read(buf)
	if err != nil {
		t.Errorf("Body okunamadı: %v", err)
	}

	if string(buf) != expected {
		t.Errorf("Beklenen body %s, ama geldi: %s", expected, string(buf))
	}
}
