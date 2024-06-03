package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

/* Basic unit test for a custom mux */
func TestCustomMux(t *testing.T) {
	mux := http.NewServeMux()
	setupCustomMuxHandlers(mux)

	tcs := []struct {
		desc   string
		path   string
		want   string
		method string
	}{
		{
			desc:   "handler created using handler func",
			path:   "/custommuxhttp",
			want:   "Custom Mux: Handler Func\n",
			method: "GET",
		},
		{
			desc:   "handler created using a custom struct that implements serveHTTP",
			path:   "/custommuxhttp2",
			want:   "Custom Mux: Handler Struct\n",
			method: "GET",
		},
	}

	for _, tc := range tcs {
		var buf bytes.Buffer
		w := httptest.NewRecorder()
		req := httptest.NewRequest(tc.method, tc.path, nil)

		mux.ServeHTTP(w, req)
		_, err := io.Copy(&buf, w.Body)
		require.NoError(t, err, "%s: %v", tc.desc, err)
		require.Equal(t, tc.want, buf.String(), tc.desc)
	}

}

/* Basic integration test for a handler using a test server */
func TestCustomHTTPMuxServer(t *testing.T) {

	mux := http.NewServeMux()
	setupCustomMuxHandlers(mux)
	testServer := httptest.NewServer(mux)

	tcs := []struct {
		desc string
		path string
		want string
	}{
		{
			desc: "server path created using handler func",
			path: "/custommuxhttp",
			want: "Custom Mux: Handler Func\n",
		},
		{
			desc: "server path created using a custom struct that implements serveHTTP",
			path: "/custommuxhttp2",
			want: "Custom Mux: Handler Struct\n",
		},
	}

	for _, tc := range tcs {
		var buf bytes.Buffer

		resp, err := http.Get(testServer.URL + tc.path)
		require.NoError(t, err, "%s: %v", tc.desc, err)
		_, err = io.Copy(&buf, resp.Body)
		require.NoError(t, err, "%s: %v", tc.desc, err)
		require.Equal(t, tc.want, buf.String(), "%s: %v", tc.desc, err)
	}
}
