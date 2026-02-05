package handlers

import(
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T){
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil{
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(HealthHandler)
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK{
		t.Errorf("expected status 200, got %d", recorder.Code)
	}

	expected := "ok"
	if recorder.Body.String() != expected {
		t.Errorf("expected body %s, got %s",
		 expected,
		  recorder.Body.String(),
		)
	}
}