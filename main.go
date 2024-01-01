package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	// "hello" を表示するgoroutineを開始
	go func() {
		fmt.Println("hello")
	}()

	// HTTPサーバーを開始
	log.Fatal(http.ListenAndServe(":8080", nil))
}
