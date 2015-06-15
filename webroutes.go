package toscaviewer

import (
	"github.com/gorilla/mux"
	"github.com/owulveryck/toscalib"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(topology toscalib.ToscaDefinition) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.Headers("Content-Type", "application/json", "X-Requested-With", "XMLHttpRequest")
	router.Methods("GET").Path("/tosca").Name("SVG Representation").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		displaySvg(w, r, topology, "functionnal")
	})
	router.Methods("GET").Path("/workflow").Name("SVG Representation").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		displaySvg(w, r, topology, "workflow")
	})
	//router.Methods("GET").Path("/tasks").Name("TaskIndex").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	showTasks(w, r, topology)
	//})
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../htdocs/")))

	return router
}
