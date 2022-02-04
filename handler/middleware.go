package handler

import (
	"log"
	"net/http"
)

func (h *Handler) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(){
			if err := recover(); err != nil {
				log.Printf("Problem in %s.\n Please fix it \n", r.URL.Path)
			}
		}()
		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		next(w, r)
	}
}
