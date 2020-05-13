// Copyright 2020 Lauris BH. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logger

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func testRequest(t *testing.T, req *http.Request, middleware ...func(h http.Handler) http.Handler) chi.Router {
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Use(middleware...)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Response status code")

	return r
}

func TestLogger(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Host = "localhost:8080"

	logger, hook := test.NewNullLogger()

	testRequest(t, req, Logger("router", logger))

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Nil(t, hook.LastEntry().Data["request_id"])
	assert.Equal(t, "GET", hook.LastEntry().Data["method"])
	assert.Equal(t, 11, hook.LastEntry().Data["bytes"])
	assert.Greater(t, hook.LastEntry().Data["duration"], int64(0))
	assert.Equal(t, "http://localhost:8080/", hook.LastEntry().Message)
}

func TestLoggerRequestID(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Host = "localhost:8080"

	logger, hook := test.NewNullLogger()

	testRequest(t, req, middleware.RequestID, Logger("router", logger))

	assert.Equal(t, 1, len(hook.Entries))
	assert.NotNil(t, hook.LastEntry().Data["request_id"])
}
