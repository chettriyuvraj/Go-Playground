package handlers

import (
	"http-server-clean/app"
	"io"
	"net/http/httptest"
	"testing"
)

func TestHandleOkRoute(t *testing.T) {
	/* Setup initial structs */
	config := app.AppConfig{}
	req := httptest.NewRequest("GET", "/ok", nil)
	recorder := httptest.NewRecorder()

	/* Execute and test */
	HandleOkRoute(recorder, req, &config)
	response := recorder.Result()
	respData, err := io.ReadAll(response.Body)
	if err != nil {
		if err != io.EOF {
			t.Errorf("error opening response body: %v", err)
		}
	}
	defer response.Body.Close()
	wantResponse := "This is an ok response!"
	if string(respData) != wantResponse {
		t.Errorf("ok route response not matching: got %s, want %s", respData, wantResponse)
	}

}
