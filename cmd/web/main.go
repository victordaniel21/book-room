package main

import (
	"book-room/pkg/handlers"
	"fmt"
	"net/http"
)

const portNum = ":8010"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Start app on port %s", portNum)
	_ = http.ListenAndServe(portNum, nil)
}
