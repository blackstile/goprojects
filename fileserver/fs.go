package main

import (
	"net/http"
)

func main() {
	// dir, _ := os.Getwd()
	dir := http.Dir(".")
	http.ListenAndServe(":3000", http.FileServer(dir))
}
