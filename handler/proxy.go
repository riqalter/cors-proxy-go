package handler

import (
	"io"
	"net/http"
	"net/url"
)

func Proxy(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")

	if target == "" {
		http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
		return
	}
}