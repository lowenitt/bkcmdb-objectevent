package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", ObjectEvent) //	设置访问路由
	log.Fatal(http.ListenAndServe(":8686", nil))
}
