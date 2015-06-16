package toscaviewer

// This is a basic example
// Thanks http://thenewstack.io/make-a-restful-json-api-go/ for the tutorial
import (
	"fmt"
	"net/http"
)

func (toscaGraph ToscaGraph) ViewToscaYaml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text; charset=UTF-8")
	fmt.Fprintf(w, string(toscaGraph["ToscaYaml"]))
}
func (toscaGraph ToscaGraph) ViewToscaDefinition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=UTF-8")
	fmt.Fprintf(w, string(toscaGraph["ToscaDefinition"]))
}
func (toscaGraph ToscaGraph) ViewToscaWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=UTF-8")
	fmt.Fprintf(w, string(toscaGraph["ToscaWorkflow"]))
}
