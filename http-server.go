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
	Url    string              `json:"url"`
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
			Url:    req.URL.String(),
			Method: req.Method,
			Header: req.Header,
		}}

	json.NewEncoder(w).Encode(body)
}

func setCookie(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Set cookies:", req.URL.Query())

	for k, v := range req.URL.Query() {
		http.SetCookie(w, &http.Cookie{Name: k, Value: v[0]})
	}
}

func main() {
	http.HandleFunc("/", info)
	http.HandleFunc("/setcookie", setCookie)

	fmt.Println("start", name())
	http.ListenAndServe(":8080", nil)
}
