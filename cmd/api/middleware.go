package main

import "net/http"

func (app *application) enableCORS(h http.Handler) http.Handler {
	// Middleware to allow Cross Origin Scripting which is disallowed by default.
	// This allows a front-end invoked via one port to access 
	// a back-end listening on another port
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, X-CSRF-Token, Authorization")
			return
		} else {
			h.ServeHTTP(w, r)
		}
	})
}