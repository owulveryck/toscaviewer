package main

import (
	"io/ioutil"
	"log"

	//"github.com/gonum/matrix/mat64"
	"github.com/owulveryck/toscalib"
	"gopkg.in/yaml.v2"
	"fmt"
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
	for nodeName, _ := range mystruct.TopologyTemplate.NodeTemplates {
		//log.Printf(" Node %v = %v", nodeName, nodeDetail.Type)
		nodeReferences[nodeName] = i
		i = i + 1
	}
	// Create an adjacency Matrix
	//adjacencyMatrix := mat64.NewDense(nodeNum, nodeNum, nil)
	//2015/06/09 11:41:51 (mysql_database) reqDetail : [map[host:map[node:mysql_dbms]]]
	//2015/06/09 11:41:51 (mysql_dbms) reqDetail : [map[host:map[node:server]]]
	//2015/06/09 11:41:51 (webserver) reqDetail : [map[host:map[node:server]]]
	//2015/06/09 11:41:51 (wordpress) reqDetail : [map[host:map[node:webserver]] map[database_endpoint:map[node:mysql_database]]]

	fmt.Printf("digraph G {\n")
	for nodeName, nodeDetail := range mystruct.TopologyTemplate.NodeTemplates {
		// If requirements are found
		if nodeDetail.Requirements != nil {
			log.Printf("(%v) reqDetail : %v",nodeName, nodeDetail.Requirements)
			for _, requirementType := range nodeDetail.Requirements {
				for requirementTypeProp, value := range requirementType {
					fmt.Printf("\t%v -> %v [label = %v];\n",value["node"], nodeName, requirementTypeProp)
				}
				//adjacencyMatrix.Set(nodeReferences[nodeName],nodeReferences[target],1)
			}
		}
	}
	fmt.Printf("}\n")
}
