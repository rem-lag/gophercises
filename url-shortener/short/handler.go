package short

import (
	"net/http"
)

// Maps shortened urls to full actual urls in the form of a map
// Keys are short urls and values are the full url to use
// If a short url cannot be found in the map it will return an http.Handler as a fallback
func MapHandler(pathToUrl map[string]string, fallback http.Handler) http.HandlerFunc {

	return nil
}

// This is the same as the map handler but uses a yaml file to provide the mappings
// The format should be like
//
//	-path: /short-path
//	-url: https://www.a-real-website.com/whatever
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	return nil, nil
}
