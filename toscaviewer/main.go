package main

import (
	"io/ioutil"
	"log"
	"net/http"

	//"github.com/gonum/matrix/mat64"
	"github.com/owulveryck/toscalib"
	"github.com/owulveryck/toscaviewer"
	"gopkg.in/yaml.v2"
	"flag"
)

func main() {

	var testFile = flag.String("testfile","../examples/tosca_single_instance_wordpress.yaml", "a tosca yaml file to process")
	flag.Parse()

	var topologyTemplate toscalib.TopologyTemplateStruct
	file, err := ioutil.ReadFile(*testFile)
	if err != nil {
		log.Panic("error: ", err)
	}
	err = yaml.Unmarshal(file, &topologyTemplate)
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
