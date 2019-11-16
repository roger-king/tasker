package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/roger-king/tasker/utils"
)

// ServeWebAdmin -
func ServeWebAdmin(w http.ResponseWriter, r *http.Request) {

	if len(utils.TaskerEnv) == 0 || utils.TaskerEnv == "local" {
		url, _ := url.Parse("http://localhost:3000")
		// create the reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(url)

		// Update the headers to allow for SSL redirection
		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = url.Host

		// Note that ServeHttp is non blocking and uses a go routine under the hood
		proxy.ServeHTTP(w, r)
	}
	return
}
