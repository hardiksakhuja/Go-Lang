package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"urlshort"
)

func main() {
	var jsonFile, yamlFile string
	flag.StringVar(&jsonFile, "json", "", "Path to the Json file")
	flag.StringVar(&yamlFile, "yaml", "", "Path to the Yaml file")

	flag.Parse()
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	if jsonFile != "" {
		jsonData, err := ioutil.ReadFile(jsonFile)
		if err != nil {
			panic(err)
		}

		jsonHandler, err := urlshort.JSONHandler([]byte(jsonData), mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", jsonHandler)
	} else if yamlFile != "" {
		yamlData, err := ioutil.ReadFile(yamlFile)
		if err != nil {
			panic(err)
		}
		yamlHandler, err := urlshort.YAMLHandler([]byte(yamlData), mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", yamlHandler)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
