package main

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gopherdojo/dojo3/kadai4/nguyengiabk/omikuji"
)

type testCase struct {
	os         OmikujiServer
	statusCode int
	response   string
}

var TestHandlerFixtures = map[string]testCase{
	"Test json encode error": {
		OmikujiServer{
			JSONEncodeFunc(func(data interface{}, w io.Writer) error {
				return errors.New("Encode error")
			}),
			omikuji.Omikuji{},
		},
		http.StatusInternalServerError,
		"Server error\n",
	},
	"Test normal case": {
		OmikujiServer{
			// want to fixed response
			omikuji: omikuji.Omikuji{Randomize: omikuji.RandomizeFunc(func(max int) int {
				return 0
			})},
		},
		http.StatusOK,
		"{\"result\":\"大吉\"}\n",
	},
}

func TestHandler(t *testing.T) {
	for name, tc := range TestHandlerFixtures {
		tc := tc
		t.Run(name, func(t *testing.T) {
			RunTestCase(t, tc)
		})
	}
}

// split to use defer
func RunTestCase(t *testing.T, tc testCase) {
	t.Helper()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	tc.os.handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != tc.statusCode {
		t.Errorf("unexpected status code, actual = %v, expected = %v", rw.StatusCode, tc.statusCode)
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Errorf("unexpected error when read response body")
	}
	if s := string(b); s != tc.response {
		t.Fatalf("unexpected response: actual = %s, expected = %s", s, tc.response)
	}
}
