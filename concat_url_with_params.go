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

func stripHash(requestURL string) string {
	hashIndex := strings.Index(requestURL, "#")

	if hashIndex != -1 {
		return requestURL[:hashIndex]
	}

	return requestURL
}

// ConcatURLWithParams concatenates a basic requestURL with
func ConcatURLWithParams(requestURL string, param interface{}) string {
	requestURL = stripHash(requestURL)

	rawQuery := ""

	queryDelimiterIndex := strings.Index(requestURL, "?")

	if queryDelimiterIndex > -1 {
		rawQuery = requestURL[queryDelimiterIndex:]
		requestURL = requestURL[:queryDelimiterIndex]
	}

	vs, err := query.Values(param)

	if err == nil {
		q := vs.Encode()

		if rawQuery != "" {
			requestURL = requestURL + rawQuery + "&" + q
		} else {
			requestURL = requestURL + "?" + q
		}
	}

	return requestURL + rawQuery
}
