package toscaviewer

import (
	"github.com/gorilla/mux"
	"github.com/owulveryck/toscalib"
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

// NewRouter is a new router instance.
func NewRouter(toscaTemplate *toscalib.ToscaDefinition) *mux.Router {
	// Initializing the ToscaGraph structure
	type ToscaDefinition toscalib.ToscaDefinition
	var toscaGraph ToscaGraph
	toscaGraph.ToscaDefinition = toscaTemplate
	toscaGraph.Initialize()

	// This is the web display

	// Definition des routes
	var routes = routes{
		route{
			"Tosca diagram",
			"GET",
			"/tosca.svg",
			(&toscaGraph).ViewToscaDefinition,
		},
		route{
			"Execution workflow",
			"GET",
			"/workflow.svg",
			(&toscaGraph).ViewToscaWorkflow,
		},
		route{
			"Tosca",
			"GET",
			"/tosca.yaml",
			(&toscaGraph).ViewToscaYaml,
		},
		route{
			"Upload Tosca",
			"POST",
			"/upload",
			(&toscaGraph).UploadHandler,
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
	router.Headers("Cache-Control", "no-cache, no-store, must-revalidate")

	return router
}
