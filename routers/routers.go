package routers

import (
	"net/http"

	"github.com/biFebriansyah/gorest/controllers"
	"github.com/biFebriansyah/gorest/repos"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	mainRutes := mux.NewRouter()

	// inisialisasi endpoint
	mainRutes.HandleFunc("/", simpleHandlers).Methods("GET")

	// inisalisasi repos
	userRep := repos.New()
	users := controllers.Users{Rp: *userRep}

	UserRoute(mainRutes, &users)
	return mainRutes

}

func simpleHandlers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellow from api"))
}
