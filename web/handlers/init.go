package handlers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

func StartRouting(r *chi.Mux) (err error)  {
	var tmpl *template.Template
	tmpl, err = template.ParseFiles("./web/tmpl/layout.html")
	if err!= nil {
		return
	}


     r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		 tmpl.Execute(writer, nil)
	 })
     r.Route("/sub", func(r chi.Router) {
		 r.Get("/user", func(writer http.ResponseWriter, request *http.Request) {
			 writer.Write([]byte("sub router"))
		 })
		 r.Get("/hello", func(writer http.ResponseWriter, request *http.Request) {
			 writer.Write([]byte("hello from sub router"))
		 })
	 })
	return
}