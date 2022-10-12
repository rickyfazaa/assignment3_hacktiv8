package main

import (
	"assignment3/static"
	"fmt"
	"net/http"
)

func main() {
	go static.AutoReload()
	http.HandleFunc("/", static.ReloadWeb)
	fmt.Println("server running on PORT:", ":4000")
	http.ListenAndServe(":4000", nil)
}
