package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Welcome to the home page, try to /hello?name=animesh"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}
	_, _ = w.Write([]byte(name))
}

// json encoder
func successHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := map[string]any{
		"ok":       true,
		"message":  "JSON Encode Sucessful",
		"datetime": time.Now().UTC(),
	}

	_ = json.NewEncoder(w).Encode(res)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

type TestRqst struct {
	Name string `json:"name"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "only POST is allowed",
		})
		return
	}
	defer r.Body.Close()

	var req TestRqst

	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Invalid JSON format",
		})
		return
	}

	req.Name = strings.TrimSpace(req.Name)

	if req.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Name must not be empty",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":        true,
		"data":      req,
		"timestamp": time.Now().UTC(),
	})
}

func main() {
	// 	http.HandleFunc("/hello", helloHandler)
	//
	// 	fmt.Println("try going to 8080 port")
	//
	// 	err := http.ListenAndServe(":8080", nil)

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/ok", successHandler)
	http.HandleFunc("/test", testHandler)
	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)
}
