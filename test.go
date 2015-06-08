package main

import (
	"io/ioutil"
	"log"

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

	log.Printf("Found %v nodes\n", len(mystruct.TopologyTemplate.NodeTemplates))
	for nodeName, nodeDetail := range mystruct.TopologyTemplate.NodeTemplates {
		log.Printf(" Node %v = %v", nodeName, nodeDetail.Type)
	}
}
