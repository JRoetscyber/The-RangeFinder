package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/download-transcript/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/download-transcript/")
		id = strings.TrimSuffix(id, ".txt")

		switch id {
		case "1":
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "TRANSCRIPT 1: [ADMIN CHAT] Password is: P@ssw0rd_Admin_9981")
		case "2":
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "TRANSCRIPT 2: User: Hi, I need help with my account invoice.")
		case "3":
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "TRANSCRIPT 3: Agent: Sure, what is your account number?")
		case "42":
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "TRANSCRIPT 42: [INTERNAL] System architecture diagram v2.0")
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "Transcript not found.")
		}
	})

	fmt.Println("🚀 Mock IDOR Target Server listening on http://127.0.0.1:8080")
	fmt.Println("👉 Try testing: http://127.0.0.1:8080/download-transcript/{ID}.txt")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
