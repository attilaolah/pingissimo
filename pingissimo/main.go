// Package pingissimo implements a simple ping service.
package pingissimo

import (
	"net/http"
	"strings"

	"appengine"

	"pingissimo/ping"
)

func init() {
	http.HandleFunc("/ping/get", pingHandler)
	http.HandleFunc("/ping/head", pingHandler)
	http.HandleFunc("/ping/post", pingHandler)
	http.HandleFunc("/ping/put", pingHandler)
	http.HandleFunc("/ping/patch", pingHandler)
	http.HandleFunc("/ping/delete", pingHandler)
	http.HandleFunc("/ping/options", pingHandler)
	http.HandleFunc("/ping/connect", pingHandler)
	http.HandleFunc("/ping", pingHandler)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	parts := strings.Split(r.URL.Path, "/")
	method := strings.ToUpper(parts[len(parts)-1])
	if method == "PING" {
		method = "HEAD"
	}
	ok, err := ping.Ping(c, method, r.URL.Query()["url"]...)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.Errorf("%s", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
