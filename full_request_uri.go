package url

import "net/http"

// GetFullRequestURI parses the full request URI
func GetFullRequestURI(r *http.Request) string {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}

	return scheme + r.Host + r.RequestURI
}

// GetFullRequestPath parses the full request path(without URI)
func GetFullRequestPath(r *http.Request) string {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}

	return scheme + r.Host
}
