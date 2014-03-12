// Package ping implements a simple HTTP Ping method.
package ping

import (
	"net/http"
	"net/url"
	"strings"

	"appengine"
	"appengine/urlfetch"
)

func Ping(c appengine.Context, u *url.URL) (ok bool, err error) {
	parts := strings.Split(u.Path, "/")
	method := strings.ToUpper(parts[len(parts)-1])
	if method == "PING" {
		method = "HEAD"
	}
	req, err := http.NewRequest(method, u.Query().Get("url"), nil)
	if err != nil {
		return
	}
	ok = true
	c.Infof("%s %s %s", req.Proto, req.Method, req.URL.String())
	if resp, err := urlfetch.Client(c).Do(req); err == nil {
		c.Infof("%s %s", resp.Proto, resp.Status)
	}
	return
}
