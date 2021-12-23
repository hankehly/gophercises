package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if destUrl, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(rw, r, destUrl, http.StatusFound)
		} else {
			fallback.ServeHTTP(rw, r)
		}
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathsToUrls, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	handler := MapHandler(pathsToUrls, fallback)
	return handler, nil
}

func parseYAML(yml []byte) (map[string]string, error) {
	var items []docItem
	err := yaml.Unmarshal(yml, &items)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]string, len(items))
	for _, item := range items {
		ret[item.Path] = item.Url
	}
	return ret, nil
}

func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	if pathsToUrls, err := parseJSON(jsonBytes); err == nil {
		return MapHandler(pathsToUrls, fallback), nil
	} else {
		return nil, err
	}
}

func parseJSON(jsonBytes []byte) (map[string]string, error) {
	var items []docItem
	err := json.Unmarshal(jsonBytes, &items)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]string, len(items))
	for _, item := range items {
		ret[item.Path] = item.Url
	}
	return ret, nil
}

// Struct fields are only unmarshalled if they are exported (have an upper case first letter)
// https://pkg.go.dev/gopkg.in/yaml.v2#Unmarshal
type docItem struct {
	Path string
	Url  string
}
