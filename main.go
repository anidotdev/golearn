package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	_, _ = w.Write([]byte("Go net/http server"))
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("try going to 8080 port")

	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
