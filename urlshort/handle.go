package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathstoUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathstoUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yamlUrls []byte, fallback http.Handler) (http.HandlerFunc, error) {

	type pathUrl struct {
		Path string `yaml:"path"`
		URL  string `yaml:"url"`
	}
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlUrls, &pathUrls)
	if err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		for _, item := range pathUrls {
			if item.Path == r.URL.Path {
				http.Redirect(w, r, item.URL, http.StatusFound)
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}

func JSONHandler(jsonUrls []byte, fallback http.Handler) (http.HandlerFunc, error) {

	type pathUrl struct {
		Path string `json:"path"`
		URL  string `json:"url"`
	}
	var pathUrls []pathUrl
	err := json.Unmarshal(jsonUrls, &pathUrls)
	if err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		for _, item := range pathUrls {
			if item.Path == r.URL.Path {
				http.Redirect(w, r, item.URL, http.StatusFound)
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}
