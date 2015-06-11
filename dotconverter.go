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
	dotCode := fmt.Sprintf("digraph G {\n")
	dotCode = fmt.Sprintf("%v\tgraph [ rankdir = \"LR\" ];\n", dotCode)
	for nodeName, nodeDetail := range topology.TopologyTemplate.NodeTemplates {
		// For each node, create a record
		dotCode = fmt.Sprintf("%v\t%v [label=\"<nodeName> %v|<nodeType> %v", dotCode, nodeName, nodeName, nodeDetail.Type)
		//Display the properties
		if nodeDetail.Properties != nil {
			dotCode = fmt.Sprintf("%v|{Properties|{", dotCode)
			for property, _ := range nodeDetail.Properties {
				dotCode = fmt.Sprintf("%v%v|", dotCode, property)
			}
			dotCode = fmt.Sprintf("%v}}", dotCode)
		}
		// Display the requirements
		if nodeDetail.Requirements != nil {
			dotCode = fmt.Sprintf("%v|{Requirements|{", dotCode)
			i := 0
			pipe := "|"
			for _, requirementAssignements := range nodeDetail.Requirements {
				for requirement, _ := range requirementAssignements {
					if i == len(requirementAssignements) {
						pipe = ""
					}
					i = i + 1
					dotCode = fmt.Sprintf("%v<%v>%v%v", dotCode, requirement, requirement, pipe)
				}
			}
			dotCode = fmt.Sprintf("%v}}", dotCode)
		}
		// Display the capabilities
		dotCode = fmt.Sprintf("%v|{<capabilities>Capabilities|{", dotCode)
		if nodeDetail.Capabilities != nil {
			i := 1
			pipe := "|"
			for capability, _ := range nodeDetail.Capabilities {
				if i == len(nodeDetail.Capabilities) {
					pipe = ""
				}
				i = i + 1
				dotCode = fmt.Sprintf("%v<%v>%v%v", dotCode, capability, capability, pipe)
			}
		}
		dotCode = fmt.Sprintf("%v}}", dotCode)

		dotCode = fmt.Sprintf("%v\" shape=record style=rounded color=blue]\n", dotCode)
		// If requirements are found
		//		dotCode = fmt.Sprintf( "\t\"%v\" [ label = \"%v\" shape = circle color=blue]\n", nodeName, nodeName)
		//		dotCode = fmt.Sprintf( "\t\"%v\" [ label = \"%v\" shape = record color=red]\n", nodeDetail.Type, nodeDetail.Type)
		//dotCode = fmt.Sprintf( "\t\"%v\" -> \"%v\" [ color = red ]\n", nodeDetail.Type, nodeName)
		if nodeDetail.Requirements != nil {
			for _, requirementAssignements := range nodeDetail.Requirements {
				for requirementName, requirementAssignement := range requirementAssignements {

					dotCode = fmt.Sprintf("%v\t%v:%v -> %v:capabilities [label = %v color=red];\n", dotCode, nodeName, requirementName, requirementAssignement.Node, requirementName)
				}
			}
		}
	}
	dotCode = fmt.Sprintf("%v}\n", dotCode)
	fmt.Printf("DEBUG: \n%v\n", dotCode)
	fmt.Fprintf(w, "%v", dotCode)
}
