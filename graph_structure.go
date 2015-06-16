package toscaviewer

import (
	"bytes"
	"fmt"
	"github.com/owulveryck/toscalib"
	"net/http"
	"os"
	"os/exec"
)

// ToscaGraph is a type holding the SVG representations of the graph
// the structure is "in memory" by now for debugging purpose
type ToscaGraph map[string][]byte

func (toscaGraph ToscaGraph) ViewToscaDefinition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=UTF-8")
	fmt.Fprintf(w, string(toscaGraph["ToscaDefinition"]))
}
func (toscaGraph ToscaGraph) ViewToscaWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=UTF-8")
	fmt.Fprintf(w, string(toscaGraph["ToscaWorkflow"]))
}

// Initialize the ToscaGraph
func (toscaGraph *ToscaGraph) Initialize(toscaDefinition toscalib.ToscaDefinition) error {
	var tempGraph ToscaGraph
	tempGraph = make(map[string][]byte, 2)
	for i, value := range []string{"ToscaDefinition", "ToscaWorkflow"} {
		dotProcess := exec.Command("dot", "-Tsvg")

		// Set the stdin stdout and stderr of the dot subprocess
		stdinOfDotProcess, err := dotProcess.StdinPipe()
		if err != nil {
			fmt.Println(err) //replace with logger, or anything you want
		}
		defer stdinOfDotProcess.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line
		readCloser, err := dotProcess.StdoutPipe()
		if err != nil {
			fmt.Println(err) //replace with logger, or anything you want
		}
		dotProcess.Stderr = os.Stderr

		// Actually run the dot subprocess
		if err = dotProcess.Start(); err != nil { //Use start, not run
			fmt.Println("An error occured: ", err) //replace with logger, or anything you want
		}
		switch i {
		case 0:
			toscaDefinition.PrintDot(stdinOfDotProcess)
			stdinOfDotProcess.Close()
		case 1:
			toscaDefinition.DotExecutionWorkflow(stdinOfDotProcess)
			stdinOfDotProcess.Close()
		}
		// Read from stdout and store it in the correct structure
		var buf bytes.Buffer
		buf.ReadFrom(readCloser)

		tempGraph[value] = buf.Bytes()

		//toscaDefinition.PrintDot(stdinOfDotProcess)
		//toscaDefinition.DotExecutionWorkflow(stdinOfDotProcess)
		dotProcess.Wait()
	}
	//*toscaGraph = ToscaGraph(make(map[string][]byte, 2))
	*toscaGraph = tempGraph
	return nil
}
