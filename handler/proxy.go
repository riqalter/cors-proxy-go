package handler

import (
	"io"
	"net/http"
	"net/url"
	"time"

	"cors-proxy/logger"
)

func Proxy(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	target := r.URL.Query().Get("url")

	if target == "" {
		logger.Warningf("Proxy request missing 'url' parameter from %s", r.RemoteAddr)
		http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
		return
	}

	targetURL, err := url.Parse(target)
	if err != nil || targetURL.Scheme == "" || targetURL.Host == ""{
		logger.Warningf("Invalid URL '%s' from %s: %v", target, r.RemoteAddr, err)
		http.Error(w, "Invalid 'url' query parameter", http.StatusBadRequest)
		return
	}

	logger.Infof("Proxying %s request to %s from %s", r.Method, target, r.RemoteAddr)

	req, err := http.NewRequest(r.Method, target, r.Body)
	if err != nil {
		logger.Errorf("Failed to create request for %s: %v", target, err)
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header = r.Header.Clone()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("Failed to fetch %s: %v", target, err)
		http.Error(w, "Failed to fetch target URL", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Log the proxy response details
	duration := time.Since(start)
	logger.LogProxyRequest(target, r.Method, r.RemoteAddr, resp.StatusCode, duration)

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func copyHeader(dst, src http.Header) {
	for k, v := range src {
		dst[k] = v
	}
}