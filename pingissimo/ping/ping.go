// Package ping implements a simple HTTP Ping method.
package ping

import (
	"net/http"
	"sync"

	"appengine"
	"appengine/urlfetch"
)

func Ping(c appengine.Context, method string, urls ...string) (ok bool, err error) {
	reqs := make([]*http.Request, len(urls))
	for i, u := range urls {
		if reqs[i], err = http.NewRequest(method, u, nil); err != nil {
			ok = false
			return
		}
	}
	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(len(reqs))
	for i, r := range reqs {
		go func(r *http.Request, i int) {
			defer wg.Done()
			c.Infof("[%d] %s %s %s", i, r.Proto, r.Method, r.URL.String())
			if resp, rerr := urlfetch.Client(c).Do(r); err == nil {
				c.Infof("[%d] %s %s", i, resp.Proto, resp.Status)
			} else {
				c.Errorf("[%d] %s", i, err)
				err = rerr
			}
		}(r, i+1)
	}
	ok = true
	return
}
