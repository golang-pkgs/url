package url

import (
	"strings"

	"github.com/google/go-querystring/query"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func stripHash(requestUrl string) string {
	hashIndex := strings.Index(requestUrl, "#")

	if hashIndex != -1 {
		return requestUrl[:hashIndex]
	}

	return requestUrl
}

func ComposeUrl(requestUrl string, params ...interface{}) string {
	requestUrl = stripHash(requestUrl)

	if len(params) == 0 {
		return requestUrl
	}

	param := params[0]

	rawQuery := ""

	queryDelimiterIndex := strings.Index(requestUrl, "?")

	if queryDelimiterIndex > -1 {
		rawQuery = requestUrl[queryDelimiterIndex:]
		requestUrl = requestUrl[:queryDelimiterIndex]
	}

	vs, err := query.Values(param)

	if err == nil {
		q := vs.Encode()

		if rawQuery != "" {
			requestUrl = requestUrl + rawQuery + "&" + q
		} else {
			requestUrl = requestUrl + "?" + q
		}
	}

	return requestUrl + rawQuery
}
