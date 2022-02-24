package main

import (
	"fmt"
	"net/http"

	handler "github.com/anikkatiyar99/forecast/pkg/http/handler"
	middleware "github.com/anikkatiyar99/forecast/pkg/http/middleware"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/forecast", handler.GetForecastHandler).Methods("GET")
	r.Use(middleware.Middleware)
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)
}
