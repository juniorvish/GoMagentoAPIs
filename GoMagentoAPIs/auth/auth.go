package auth

import (
	"net/http"
)

const AuthTokenHeader = "X-Auth-Token"

func CheckAuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get(AuthTokenHeader)
		if authToken == "" {
			http.Error(w, "Unauthorized: Missing auth token", http.StatusUnauthorized)
			return
		}

		// Add your logic here to validate the authToken with Magento API

		next.ServeHTTP(w, r)
	})
}