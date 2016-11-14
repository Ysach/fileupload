package main

import (
	"net/http"
	"log"
	"./upload"
	"./index"
	"./login"
)


func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index.Index)       //设置访问的路由
	http.HandleFunc("/login", login.Login)         //设置访问的路由
	http.HandleFunc("/upload", upload.Upload)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

