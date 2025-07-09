package server

import (
	"encoding/json"
	"net/http"
	"youtube-downloader/internal/ytdlp"
)

func Start() {
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.Handle("/api/getMediaDetails", http.HandlerFunc(getMediaDetailsHandler))
	http.Handle("/api/download", http.HandlerFunc(downloadHandler))
	http.ListenAndServe(":8080", nil)
}

func getMediaDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	url := r.FormValue("url")
	formats, err := ytdlp.GetFormats(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("yt-dlp error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(formats)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	url := r.FormValue("url")
	formatID := r.FormValue("format_id")
	filename := r.FormValue("filename")
	if url == "" || formatID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing url or format_id"))
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"")
	err := ytdlp.StreamDownload(w, url, formatID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("yt-dlp download error"))
	}
}
