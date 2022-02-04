package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func (c *Handler) SendEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if err := c.confirm.DeliverEmail(email); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400), 400)
		return
	}
}
func (c *Handler) Receive(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("confrim")
	val, ok := c.confirm.Get(code)
	if !ok {
		http.Error(w, "Sorry you arent logged in", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "hello "+val+" rady vas videt")
}

func (c *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("main.html")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if err = t.Execute(w, nil); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}
