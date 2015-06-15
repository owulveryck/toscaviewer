package toscaviewer

import (
	"github.com/owulveryck/toscalib"
	"log"
	"reflect"
)

type ToscaNodesCompute toscalib.NodeTemplate

func (nodeTemplate *ToscaNodesCompute) Create() error {
	log.Printf("Running the ToscaNodesCompute's Create method")
	if nodeTemplate.Interfaces != nil {
		log.Printf("Found Interfaces %v", nodeTemplate.Interfaces)
	}
	return nil
}

func (nodeTemplate *ToscaNodesCompute) Configure() error {
	log.Printf("Running the ToscaNodesCompute's Configure method")
	return nil
}

func (nodeTemplate *ToscaNodesCompute) Start() error {
	log.Printf("Running the ToscaNodesCompute's Start method")
	return nil
}

func (nodeTemplate *ToscaNodesCompute) Stop() error {
	log.Printf("Running the ToscaNodesCompute's Stop method")
	return nil
}

func (nodeTemplate *ToscaNodesCompute) Delete() error {
	log.Printf("Running the ToscaNodesCompute's Delete method")
	return nil
}

//DefaultNodeType is set when no other implementation is found
// By default it trigger the artifcat for any method if any interface implementation is found
type DefaultNodeType toscalib.NodeTemplate

func (nodeTemplate *DefaultNodeType) Create() error {
	if nodeTemplate.Interfaces != nil {
		for _, methods := range nodeTemplate.Interfaces {
			for method, value := range methods {
				if method == "create" {
					v := reflect.ValueOf(value)
					if v.Kind() == reflect.Map {
						log.Printf("Running %v", value)
					} else {
						log.Printf("Running %v", value)
					}
				}
			}
		}
		//log.Printf("Found Interfaces %v", nodeTemplate.Interfaces)
	}
	return nil
}

func (nodeTemplate *DefaultNodeType) Configure() error {
	log.Printf("Running the DefaultNodeType's Configure method")
	return nil
}

func (nodeTemplate *DefaultNodeType) Start() error {
	log.Printf("Running the DefaultNodeType's Start method")
	return nil
}

func (nodeTemplate *DefaultNodeType) Stop() error {
	log.Printf("Running the DefaultNodeType's Stop method")
	return nil
}

func (nodeTemplate *DefaultNodeType) Delete() error {
	log.Printf("Running the DefaultNodeType's Delete method")
	return nil
}
