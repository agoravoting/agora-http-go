package testing

import (
	"bytes"
	"encoding/json"
	s "github.com/agoravoting/agora-http-go/server"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestServer struct {
	t *testing.T
}

// initializes the test server
func New(t *testing.T, config string) (ts *TestServer) {
	var (
		name string
	)

	// generate config file. needs to be done this way, because go test could be
	// being executed in any path and we can't assume it's anywhere
	if !s.Server.Initialized {
		f, _ := ioutil.TempFile("", "testfile")
		name = f.Name()
		f.Write([]byte(config))
		f.Close()
	}

	ts = &TestServer{t: t}

	if err := s.Server.Init(name); err != nil {
		panic(err)
	}
	return
}

// tears down the test server
func (ts *TestServer) TearDown() {
}

func (ts *TestServer) Request(method, path string, expectedStatus int, headers map[string]string, requesTBody string) string {
	r, _ := http.NewRequest(method, path, bytes.NewBufferString(requesTBody))
	w := httptest.NewRecorder()
	u := r.URL
	r.RequestURI = u.RequestURI()
	for key, value := range headers {
		r.Header.Set(key, value)
	}
	s.Server.Http.ServeHTTP(w, r)
	body := w.Body.String()
	if w.Code != expectedStatus {
		ts.t.Errorf("Expected %d for route %s %s found: Code=%d, req-Headers=%v ret-body=%s\n", expectedStatus, method, u, w.Code, headers, body)
	}

	return body
}

func (ts *TestServer) RequestJson(method, path string, expectedStatus int, headers map[string]string, requestBody string) interface{} {
	body := ts.Request(method, path, expectedStatus, headers, requestBody)
	var f interface{}
	err := json.Unmarshal([]byte(body), &f)
	if err != nil {
		ts.t.Error(err)
	}
	return f
}
