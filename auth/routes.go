package auth

import "net/http"

func HandleAuthRoutes(r *http.ServeMux) {
	r.Handle("/auth/singup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "post" {
			w.WriteHeader(201)
			w.Write([]byte("{status}"))
		}
	}))
}
