package server

import (
	"log"
	"mailsender/handler"
	"net/http"
	"os"
)

func Run(h *handler.Handler) error {
	port := os.Getenv("serverAddr")
	srv := http.Server{
		Addr:    ":" + port,
		Handler: http.DefaultServeMux,
	}

	http.Handle("/sendemail", h.Middleware(h.SendEmail))
	http.Handle("/receive",h.Middleware(h.Receive))
	http.Handle("/",h.Middleware(h.MainPage))
	log.Printf("Server started on %v address \n", port)
	return srv.ListenAndServe()
}
