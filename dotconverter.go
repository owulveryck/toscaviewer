package toscaviewer

// This is a basic example
// Thanks http://thenewstack.io/make-a-restful-json-api-go/ for the tutorial
import (
	"fmt"
	"github.com/owulveryck/toscalib"
	"io"
)
// PrintDot convert the TopologyTemplateStructure in dot format
// in order to generate a graph with graphviz
// This function is mostly used for debugging purpose and may change a lot in the future
func PrintDot(w io.Writer, topology toscalib.TopologyTemplateStruct) {
	fmt.Fprintf(w, "digraph G {\n")
	for nodeName, nodeDetail := range topology.TopologyTemplate.NodeTemplates {
		// If requirements are found
		if nodeDetail.Requirements != nil {
			for _, requirementType := range nodeDetail.Requirements {
				for requirementTypeProp, value := range requirementType {
					fmt.Fprintf(w, "\t%v -> %v [label = %v];\n", value["node"], nodeName, requirementTypeProp)
				}
			}
		}
	}
	fmt.Fprintf(w, "}\n")
}
