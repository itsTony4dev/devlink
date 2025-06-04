package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DevLink is running 🚀!"))
	}).Methods("GET")

	return r
}
