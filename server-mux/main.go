package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/message", ShowMessage)
	http.ListenAndServe(":8081", mux)
}

func ShowMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello Human"))
}
