package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TARUNGORKA09/ToDo/handlers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "Project API", log.LstdFlags)

	proj := handlers.NewProject(l)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{title:[a-zA-Z]}", proj.GetProjects)

	http.ListenAndServe(":8080", sm)

}
