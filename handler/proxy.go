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

	targetURL, err := url.Parse(target)
	if err != nil || targetURL.Scheme == "" || targetURL.Host == ""{
		http.Error(w, "Invalid 'url' query parameter", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest(r.Method, target, r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header = r.Header.Clone()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch target URL", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func copyHeader(dst, src http.Header) {
	for k, v := range src {
		dst[k] = v
	}
}