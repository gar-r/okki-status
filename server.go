package main

import (
	"bitbucket.org/dargzero/okki-status/core"
	"errors"
	"log"
	"net/http"
	"strings"
)

func startServer() {
	http.HandleFunc("/invalidate/", invalidateHandler)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func invalidateHandler(w http.ResponseWriter, r *http.Request) {
	module, err := findModule(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	invalidate(module)
	w.WriteHeader(http.StatusOK)
}

func findModule(r *http.Request) (core.Module, error) {
	idx := strings.LastIndex(r.URL.Path, "/")
	name := r.URL.Path[idx+1 : len(r.URL.Path)]
	for _, module := range config {
		if module.Name == name {
			return module, nil
		}
	}
	return core.Module{}, errors.New("module not found: " + name)
}
