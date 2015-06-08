package main

import (
	"io/ioutil"
	"log"

	"github.com/gonum/matrix/mat64"
	"github.com/owulveryck/toscalib"
	"gopkg.in/yaml.v2"
)

func main() {

	var mystruct toscalib.TopologyTemplateStruct
	file, err := ioutil.ReadFile("examples/tosca_single_instance_wordpress.yaml")
	if err != nil {
		log.Panic("error: ", err)
	}
	err = yaml.Unmarshal(file, &mystruct)
	if err != nil {
		log.Panic("error: ", err)
	}
	// Here we have the structure
	// Count the nodes

	//log.Printf("Found %v nodes\n", len(mystruct.TopologyTemplate.NodeTemplates))
	nodeNum := len(mystruct.TopologyTemplate.NodeTemplates)
	i := 0
	nodeReferences := make(map[string]int, nodeNum)
	for nodeName, nodeDetail := range mystruct.TopologyTemplate.NodeTemplates {
		//log.Printf(" Node %v = %v", nodeName, nodeDetail.Type)
		nodeReferences[nodeName] = i
		i = i + 1
		log.Println(nodeDetail.Requirements)
	}
	// Create an adjacency Matrix
	adjacencyMatrix := mat64.NewDense(nodeNum, nodeNum, nil)
	for nodeName, nodeDetail := range mystruct.TopologyTemplate.NodeTemplates {
		// If requirements are found
		if nodeDetail.Requirements != nil {
			for _, reqDetail := range nodeDetail.Requirements {
				// Set 1 to the adjacencymatrix
				adjacencyMatrix.Set(nodeReferences[nodeName], nodeReferences[reqDetail["host"]], 1)
			}
		}
	}

}
