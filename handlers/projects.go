package handlers

import (
	"log"
	"net/http"

	"github.com/TARUNGORKA09/ToDo/tree/master/data"
	"github.com/gorilla/mux"
)

type Project struct {
	l *log.Logger
}

func NewProject(l *log.Logger) *Project {
	return &Project{l}
}

func (p *Project) GetProjects(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	lp := data.GetProject(vars)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to encode", http.StatusBadRequest)
	}
}
