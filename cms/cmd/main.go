package main

import (
	"goprojects/cms"
	"net/http"
)

func main() {
	// p := &cms.Page{
	// 	Title:   "Hello, world",
	// 	Content: "This is the body of our webpage",
	// }
	// cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p)

	http.HandleFunc("/", cms.ServeIndex)
	http.HandleFunc("/new", cms.HandleNew)
	http.HandleFunc("/page/", cms.ServePage)
	http.HandleFunc("/post/", cms.ServePost)

	http.ListenAndServe(":3000", nil)

}
