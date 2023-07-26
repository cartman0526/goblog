package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>这是主页</h1>")
	case "/about":
		fmt.Fprint(w, "<h1>这是一个网站简介</h1>"+
			"<a href=\"mailto:tangwenfeng0301@outlook.com\">tangwenfeng0301@outlook.com</a>")
	default:
		fmt.Fprint(w, "<h1>请求页面未找到 :( </h1>")
	}
}
