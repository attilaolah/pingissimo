// Package pingissimo implements a simple ping service.
package pingissimo

import (
	"net/http"

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
	ok, err := ping.Ping(c, r.URL)
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
