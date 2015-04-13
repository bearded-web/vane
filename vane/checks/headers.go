package checks

import (
	"net/http"
)

var (
	nonInterestingHeaders = []string{
		"Location",
		"Date",
		"Content-Type",
		"Content-Length",
		"Connection",
		"Etag",
		"Expires",
		"Last-Modified",
		"Pragma",
		"Vary",
		"Cache-Control",
		"X-Pingback",
		"Accept-Ranges",
	}
)

func FilterInterestingHeaders(headers http.Header) http.Header {
	for _, key := range nonInterestingHeaders {
		headers.Del(key)
	}

	return headers
}
