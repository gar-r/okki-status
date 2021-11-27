package refresh

import (
	"log"
	"net/http"
	"strings"
)

var addr = ":12650"

func Listen() {
	http.HandleFunc("/invalidate/", invalidateHandler)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func invalidateHandler(w http.ResponseWriter, r *http.Request) {
	name := getName(r)
	module := find(name)
	if module == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	refresh(module)
	w.WriteHeader(http.StatusOK)
}

func getName(r *http.Request) string {
	idx := strings.LastIndex(r.URL.Path, "/")
	return r.URL.Path[idx+1 : len(r.URL.Path)]
}
