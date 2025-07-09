package server

import (
	"net/http"
)

func Start() {
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.Handle("/api/getMediaDetails", http.HandlerFunc(getMediaDetailsHandler))
	http.ListenAndServe(":8080", nil)
}

func getMediaDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	url := r.FormValue("url")
	println("Received URL:", url)
	w.WriteHeader(http.StatusOK)
}
