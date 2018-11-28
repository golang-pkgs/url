package url

import "net/http"

func GetFullQuestURI(r *http.Request) string {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}

	return scheme + r.Host + r.RequestURI
}

func GetFullQuestPath(r *http.Request) string {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}

	return scheme + r.Host
}
