package main

import (
	"fmt"
	"net/http"
	"os"
)

func name() string {
	return os.Getenv("NAME")
}

func info(w http.ResponseWriter, req *http.Request) {
	serverName := name()

	if len(serverName) > 0 {
		fmt.Fprintf(w, "Server: %v\n\n", serverName)
	}

	fmt.Fprintf(w, "Headers:\n")

	for name, headers := range req.Header {
		for _, v := range headers {
			fmt.Fprintf(w, "   %v: %v\n", name, v)
		}
	}
}

func main() {
	http.HandleFunc("/", info)

	fmt.Println("start", name())
	http.ListenAndServe(":8080", nil)
}
