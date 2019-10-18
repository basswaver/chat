package main

import (
	"chat/server"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var port string = ":5000"

	if len(os.Args) >= 2 {
		port = fmt.Sprintf(":%s", os.Args[1])
	}

	http.HandleFunc("/user", server.HandleUser)

	http.ListenAndServe(port, nil)
}
