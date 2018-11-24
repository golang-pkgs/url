package url

import "net/http"

func GetFullQuestURI(r *http.Request) string {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}

	println(scheme + r.Host + r.RequestURI)

	return scheme + r.Host + r.RequestURI
}
