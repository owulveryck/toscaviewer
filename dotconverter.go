package toscaviewer

// This is a basic example
// Thanks http://thenewstack.io/make-a-restful-json-api-go/ for the tutorial
import (
	"fmt"
	"github.com/owulveryck/toscalib"
	"io"
)


// PrintDot convert the ToscaDefinitionure in dot format
// in order to generate a graph with graphviz
// This function is mostly used for debugging purpose and may change a lot in the future
func PrintDot(w io.Writer, topology toscalib.ToscaDefinition) {
	fmt.Fprintf(w, "digraph G {\n")
	fmt.Fprintf(w, "\tgraph [ rankdir = \"LR\" ];\n")
	fmt.Printf("digraph G {\n")
	fmt.Printf("\tgraph [ rankdir = \"LR\" ];\n")
	for nodeName, nodeDetail := range topology.TopologyTemplate.NodeTemplates {
	        // For each node, create a record
		fmt.Fprintf(w,"\t%v [label=\"<nodeName> %v|<nodeType> %v",nodeName, nodeName, nodeDetail.Type)
		fmt.Printf("\t%v [label=\"<nodeName> %v|<nodeType> %v",nodeName, nodeName, nodeDetail.Type)
		//Display the properties
		if nodeDetail.Properties != nil  {
		    fmt.Fprintf(w, "|{Properties|{")
		    fmt.Printf("|{Properties|{")
		    for property, _ := range nodeDetail.Properties {
			fmt.Fprintf(w,"%v|",property)
			fmt.Printf("%v|",property)
		    }
		    fmt.Fprintf(w, "}}")
		    fmt.Printf("}}")
		}
		fmt.Fprintf(w,"\" shape=record style=rounded color=blue]\n")
		fmt.Printf("\" shape=record style=rounded color=blue]\n")
		// If requirements are found
//		fmt.Fprintf(w, "\t\"%v\" [ label = \"%v\" shape = circle color=blue]\n", nodeName, nodeName)
		fmt.Fprintf(w, "\t\"%v\" [ label = \"%v\" shape = record color=red]\n", nodeDetail.Type, nodeDetail.Type)
		//fmt.Fprintf(w, "\t\"%v\" -> \"%v\" [ color = red ]\n", nodeDetail.Type, nodeName)
		if nodeDetail.Requirements != nil {
			for _, requirementAssignements := range nodeDetail.Requirements {
			    for requirementName , requirementAssignement := range requirementAssignements {
				fmt.Fprintf(w, "\t%v -> %v [label = %v];\n", nodeName, requirementAssignement.Node, requirementName)
			    }
			}
		}
	}
	fmt.Fprintf(w, "}\n")
}
