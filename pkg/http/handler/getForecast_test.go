package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetForecastHandlerHandler1 checks if the request is valid
func TestGetForecastHandlerHandler1(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/forecast", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Co-Ordinates", "38.2527° N, 105.7585° W")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetForecastHandler)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusOK)
	}
}

// TestGetForecastHandlerHandler2 checks status 400 for missing comma in header co-ordinates
func TestGetForecastHandlerHandler2(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/forecast", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Co-Ordinates", "38.2527° N 85.7585° W")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetForecastHandler)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusBadRequest)
	}
}

// TestGetForecastHandlerHandler3 checks status 400 for missing direction
func TestGetForecastHandlerHandler3(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/forecast", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Co-Ordinates", "38.2527° , 85.7585° W")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetForecastHandler)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusBadRequest)
	}
}

// TestGetForecastHandlerHandler4 checks status 400 for bad direction
func TestGetForecastHandlerHandler4(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/forecast", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Co-Ordinates", "38.2527° F, 85.7585° W")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetForecastHandler)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusBadRequest)
	}
}

/*
// TestGetForecastHandlerHandler5 returns status 400 for invalid offer
func TestGetForecastHandlerHandler5(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/forecast", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Co-Ordinates", "38.2527° N, 85.7585° W")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetForecastHandler)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusBadRequest)
	}
}
*/

// TestGetForecastHandlerHandler6 returns status 400 for missing minutes char
func TestGetForecastHandlerHandler6(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/forecast", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Co-Ordinates", "38.2527 N, 85.7585° W")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetForecastHandler)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusBadRequest)
	}
}

// TestGetForecastHandlerHandler7 checks status 200 for Negative values for south & west
func TestGetForecastHandlerHandler7(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/forecast", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Co-Ordinates", "38.2527° S, 85.7585° W")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetForecastHandler)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusNotFound)
	}
}

/*
// TestGetForecastHandlerHandler8 checks for status 400 if offer discount is higher than total price
func TestGetForecastHandlerHandler8(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/forecast", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Co-Ordinates", "38.2527° N, 85.7585° W")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetForecastHandler)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusBadRequest)
	}
}

// TestGetForecastHandlerHandler9 checks checks for status 400 for bad order item name
func TestGetForecastHandlerHandler9(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/forecast", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Co-Ordinates", "38.2527° N, 85.7585° W")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetForecastHandler)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusBadRequest)
	}
}

// TestGetForecastHandlerHandler11 checks status 400 for invalid product price
func TestGetForecastHandlerHandler11(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/forecast", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Co-Ordinates", "38.2527° N, 85.7585° W")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetForecastHandler)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusBadRequest)
	}
}
*/

// TestGetForecastHandlerHandler12 checks status 400 for for missing request
func TestGetForecastHandlerHandler12(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/forecast", errReader(0))
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetForecastHandler)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusBadRequest)
	}
}
