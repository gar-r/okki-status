package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"hu.okki.okki-status/core"
)

func Test_InvalidateHandler(t *testing.T) {

	modules = []core.Module{
		{
			Name:       "test",
			Gap:        core.DefaultGap,
			BlockOrder: core.IconFirst,
			Icon:       &core.StaticIcon{Icon: ""},
			Status:     &DummyStatusProvider{},
			Refresh:    1 * time.Second,
		},
	}

	bar = core.NewBar(modules)

	outputFn = stdout

	t.Run("existing module", func(t *testing.T) {
		req := makeRequest(t, "test")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(invalidateHandler)
		handler.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Error("expected status OK")
		}
	})

	t.Run("invalid module", func(t *testing.T) {
		req := makeRequest(t, "dummy")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(invalidateHandler)
		handler.ServeHTTP(rr, req)
		if rr.Code != http.StatusNotFound {
			t.Error("expected status: not found")
		}
	})
}

func makeRequest(t *testing.T, module string) *http.Request {
	t.Helper()
	req, err := http.NewRequest("GET", fmt.Sprintf("/invalidate/%s", module), nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

type DummyStatusProvider struct {
}

func (*DummyStatusProvider) GetStatus() string {
	return ""
}
