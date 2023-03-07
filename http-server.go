package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Body struct {
	Name    string  `json:"name"`
	Request Request `json:"request"`
}

type Request struct {
	Path   string              `json:"path"`
	Method string              `json:"method"`
	Header map[string][]string `json:"headers"`
}

func name() string {
	return os.Getenv("NAME")
}

func info(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := Body{
		Name: name(),
		Request: Request{
			Path:   req.URL.Path,
			Method: req.Method,
			Header: req.Header,
		}}

	json.NewEncoder(w).Encode(body)
}

func main() {
	http.HandleFunc("/", info)

	fmt.Println("start", name())
	http.ListenAndServe(":8080", nil)
}
