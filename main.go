package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Wrold!") //这个写入到w的是输出到客户端的
}
func main() {
	http.HandleFunc("/api", sayhelloName) //设置访问的路由
	var port = "8080"
	if key := os.Getenv("PORT_KEY"); key != "" {
		port = os.Getenv(key)
	}
	fmt.Println("server run port:", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
