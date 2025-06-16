package middleware

import(
	"net/http"
	
	"cors-proxy/logger"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		
		if r.Method == "OPTIONS" {
			logger.Infof("CORS preflight request from %s for %s", r.RemoteAddr, r.URL.Path)
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}