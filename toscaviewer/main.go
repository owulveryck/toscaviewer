package main

import (
	"log"
	"net/http"
	"os"

	//"github.com/gonum/matrix/mat64"
	"flag"
	"github.com/owulveryck/toscalib"
	"github.com/owulveryck/toscaviewer"
)

func main() {

	var testFile = flag.String("testfile", "../examples/tosca_single_instance_wordpress.yaml", "a tosca yaml file to process")
	flag.Parse()

	var toscaTemplate toscalib.ToscaDefinition
	file, err := os.Open(*testFile)
	if err != nil {
		log.Panic("error: ", err)
	}
	err = toscaTemplate.Parse(file)
	if err != nil {
		log.Panic("error: ", err)
	}
	// Here we have the structure
	// Count the nodes

	//log.Printf("Found %v nodes\n", len(toscaTemplate.TopologyTemplate.NodeTemplates))
	nodeNum := len(toscaTemplate.TopologyTemplate.NodeTemplates)
	i := 0
	nodeReferences := make(map[string]int, nodeNum)
	for nodeName, _ := range toscaTemplate.TopologyTemplate.NodeTemplates {
		nodeReferences[nodeName] = i
		i = i + 1
	}
	// Create an adjacency Matrix
	//adjacencyMatrix := mat64.NewDense(nodeNum, nodeNum, nil)

	// TEST the implementation of the node type instanciation
	for nodeName, nodeImplementation := range toscaTemplate.TopologyTemplate.NodeTemplates {
		log.Printf("Playing %v", nodeName)
		switch nodeImplementation.Type {
		case "tosca.nodes.Compute":
			//var test toscaviewer.ToscaNodesCompute
			test := toscaviewer.ToscaNodesCompute(nodeImplementation)
			var testInterface toscalib.ToscaInterfacesNodeLifecycleStandarder
			testInterface = &test
			testInterface.Create()
		default:
			test := toscaviewer.DefaultNodeType(nodeImplementation)
			var testInterface toscalib.ToscaInterfacesNodeLifecycleStandarder
			testInterface = &test
			testInterface.Create()
		}
	}

	// This is the web display
	router := toscaviewer.NewRouter(toscaTemplate)

	go log.Fatal(http.ListenAndServe(":8080", router))
	//log.Println("connect here: http://localhost:8080/svg")

}
