package middleware

import "net/http"

type CORSMiddleware struct {
	allowedOrigins []string
	allowedMethods []string
	allowedHeaders []string
}

func NewCORSMiddleware(origins, methods, headers []string) *CORSMiddleware {
	return &CORSMiddleware{
		allowedOrigins: origins,
		allowedMethods: methods,
		allowedHeaders: headers,
	}
}

func (c *CORSMiddleware) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Check if the origin is allowed
		allowed := false
		for _, allowedOrigin := range c.allowedOrigins {
			if origin == allowedOrigin {
				allowed = true
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		if !allowed {
			http.Error(w, "Not allowed", http.StatusForbidden)
			return
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "86400") // 24 hours

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
