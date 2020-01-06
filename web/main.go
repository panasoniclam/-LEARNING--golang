package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/panasoniclam/golang-basic/web/handlers"
	"net/http"
)

func main() {
   r:= chi.NewRouter()
   r.Use(middleware.Logger)
   handlers.StartRouting(r)


   http.ListenAndServe(":8000", r)

}