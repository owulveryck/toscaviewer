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

	// Fet the rooted path name of the current directory
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	example := fmt.Sprintf("%v/../examples/tosca_single_instance_wordpress.yaml", pwd)
	example = filepath.Clean(example)
	var testFile = flag.String("testfile", example, "a tosca yaml file to process")
	flag.Parse()

	var toscaTemplate toscalib.ToscaDefinition
	file, err := os.Open(*testFile)

	if err != nil {
		log.Panic("error: ", err)
	}
	//err = yaml.Unmarshal(file, &toscaTemplate)
	err = toscaTemplate.Parse(file)
	if err != nil {
		log.Panic("error: ", err)
	}
	router := toscaviewer.NewRouter(&toscaTemplate)

	log.Println("connect here: http://localhost:8080/svg")
	log.Fatal(http.ListenAndServe(":8080", router))

}
