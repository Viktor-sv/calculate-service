package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testBodyData = "{\"a\": 4, \"b\": 5}"

func TestValidateWithoutBody(t *testing.T) {
	router := httprouter.New()
	router.POST("/", Validate(nil))

	client := http.Client{}
	ts := httptest.NewServer(router)
	defer ts.Close()
	req, _ := http.NewRequest("POST", ts.URL, nil)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res != nil {
		if res.StatusCode == http.StatusOK {
			t.Errorf("Bad status code: %d, want %d", res.StatusCode, http.StatusOK)
		}
	}

}
