package short

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

// Maps shortened urls to full actual urls in the form of a map
// Keys are short urls and values are the full url to use
// If a short url cannot be found in the map it will return an http.Handler as a fallback
func MapHandler(pathToUrl map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if val, ok := pathToUrl[path]; ok {
			http.Redirect(w, r, val, http.StatusFound)
			return
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// This is the same as the map handler but uses a yaml file to provide the mappings
// The format should be like
//
//	-path: /short-path
//	-url: https://www.a-real-website.com/whatever

type yamlUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var m []yamlUrl
	err := yaml.Unmarshal(yml, &m)
	if err != nil {
		return nil, err
	}

	paths := make(map[string]string)
	for _, p := range m {
		paths[p.URL] = p.URL
	}

	return MapHandler(paths, fallback), nil
}
