package main

import (
	"log"
	"net/http"

	"flag"
	"fmt"
	"github.com/owulveryck/toscalib"
	"github.com/owulveryck/toscaviewer"
	"os"
	"path/filepath"
)

func main() {

	// Get the rooted path name of the current directory
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	example := fmt.Sprintf("%v/../examples/tosca_single_instance_wordpress.yaml", pwd)
	example = filepath.Clean(example)
	var testFile = flag.String("testfile", example, "a tosca yaml file to process")
	flag.Parse()

	var topologyTemplate toscalib.ToscaDefinition
	file, err := os.Open(*testFile)

	if err != nil {
		log.Panic("error: ", err)
	}
	err = topologyTemplate.Parse(file)
	//err = yaml.Unmarshal(file, &topologyTemplate)
	if err != nil {
		log.Panic("error: ", err)
	}
	// Here we have the structure
	// Count the nodes

	//log.Printf("Found %v nodes\n", len(topologyTemplate.TopologyTemplate.NodeTemplates))
	nodeNum := len(topologyTemplate.TopologyTemplate.NodeTemplates)
	i := 0
	nodeReferences := make(map[string]int, nodeNum)
	for nodeName, _ := range topologyTemplate.TopologyTemplate.NodeTemplates {
		nodeReferences[nodeName] = i
		i = i + 1
	}
	// Create an adjacency Matrix
	//adjacencyMatrix := mat64.NewDense(nodeNum, nodeNum, nil)

	// This is the web displa
	router := toscaviewer.NewRouter(topologyTemplate)

	go log.Fatal(http.ListenAndServe(":8080", router))
	log.Println("connect here: http://localhost:8080/svg")

}
