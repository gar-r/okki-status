package main

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"hu.okki.okki-status/core"
)

func listenForExternal() {
	http.HandleFunc("/invalidate/", invalidateHandler)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func invalidateHandler(w http.ResponseWriter, r *http.Request) {
	module, err := findModule(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	bar.Invalidate(module)
	w.WriteHeader(http.StatusOK)
}

func findModule(r *http.Request) (core.Module, error) {
	name := getName(r)
	for _, module := range modules {
		if module.Name == name {
			return module, nil
		}
	}
	return core.Module{}, errors.New("module not found: " + name)
}

func getName(r *http.Request) string {
	idx := strings.LastIndex(r.URL.Path, "/")
	return r.URL.Path[idx+1 : len(r.URL.Path)]
}
