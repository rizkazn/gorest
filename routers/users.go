package routers

import (
	"github.com/biFebriansyah/gorest/controllers"
	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router, cr *controllers.Users) {
	route := r.PathPrefix("/users").Subrouter()
	route.HandleFunc("/", cr.GetAll).Methods("GET")
	route.HandleFunc("/", cr.Add).Methods("POST")
}
