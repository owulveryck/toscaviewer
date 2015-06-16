package toscaviewer

import (
	"github.com/gorilla/mux"
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

func NewRouter(toscaGraph ToscaGraph) *mux.Router {

	// Definition des routes
	var routes = routes{
		route{
			"Tosca diagram",
			"GET",
			"/tosca.svg",
			toscaGraph.ViewToscaDefinition,
		},
		route{
			"Execution workflow",
			"GET",
			"/workflow.svg",
			toscaGraph.ViewToscaWorkflow,
		},
		route{
			"Tosca file",
			"GET",
			"/tosca.yaml",
			toscaGraph.ViewToscaYaml,
		},
	}
	router := mux.NewRouter().StrictSlash(true)
	//router.Headers("Content-Type", "application/json", "X-Requested-With", "XMLHttpRequest")
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	// Define the access to the root of the web
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../htdocs/")))

	return router
}
