package toscaviewer

// This is a basic example
// Thanks http://thenewstack.io/make-a-restful-json-api-go/ for the tutorial
import (
	"fmt"
	"github.com/owulveryck/toscalib"
	"net/http"
	"os"
	"os/exec"
)

func displaySvg(w http.ResponseWriter, r *http.Request, topology toscalib.TopologyTemplateStruct) {
	subProcess := exec.Command("dot", "-Tsvg")

	stdin, err := subProcess.StdinPipe()
	if err != nil {
		fmt.Println(err) //replace with logger, or anything you want
	}
	defer stdin.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line

	subProcess.Stdout = w
	subProcess.Stderr = os.Stderr
	if err = subProcess.Start(); err != nil { //Use start, not run
		fmt.Println("An error occured: ", err) //replace with logger, or anything you want

	}
	topology.PrintDot(stdin)
	// Command was successful
	stdin.Close()
	subProcess.Wait()

}
