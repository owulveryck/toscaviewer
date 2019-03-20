package toscaviewer

// This is a basic example
// Thanks http://thenewstack.io/make-a-restful-json-api-go/ for the tutorial
import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
)

// UploadHandler courtesy of http://noypi-linux.blogspot.fr/2014/07/golang-web-server-basic-operatons-using.html
func (toscaGraph *ToscaGraph) UploadHandler(res http.ResponseWriter, req *http.Request) {
	// For the redirect at the end
	//url, _ := mux.CurrentRoute(req).Subrouter().Get("/index.html").URL()
	log.Println("UploadHandler called")
	var (
		status int
		err    error
	)
	defer func() {
		if nil != err {
			http.Error(res, err.Error(), status)
		}
	}()
	// parse request
	const _24K = (1 << 20) * 24
	if err = req.ParseMultipartForm(_24K); nil != err {
		status = http.StatusInternalServerError
		return
	}
	for _, fheaders := range req.MultipartForm.File {
		for _, hdr := range fheaders {
			// open uploaded
			var infile multipart.File
			if infile, err = hdr.Open(); nil != err {
				status = http.StatusInternalServerError
				return
			}

			err = toscaGraph.ToscaDefinition.Parse(infile)
			//res.Write([]byte(toscaGraph.ToscaDefinition.Bytes()))
			if err != nil {
				log.Println(err)
			}
			toscaGraph.Initialize()
			//http.Redirect(res, req, url.String(), http.StatusFound)
			//res.Write([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))))
		}
	}
	http.Redirect(res, req, "/", http.StatusFound)
}

// GetState returns a json file of the current topology
func (toscaGraph *ToscaGraph) GetState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	d, err := json.Marshal(toscaGraph.ToscaDefinition.TopologyTemplate.NodeTemplates)
	if err == nil {
		fmt.Fprintf(w, string(d))
	}
}

// ViewToscaYaml is a http handler that output the yaml file
func (toscaGraph *ToscaGraph) ViewToscaYaml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text; charset=UTF-8")
	fmt.Fprintf(w, string(toscaGraph.Graph["ToscaYaml"]))
}

// ViewToscaDefinition is a http handler that output the SVG representation of the current tosca structure
func (toscaGraph *ToscaGraph) ViewToscaDefinition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=UTF-8")
	fmt.Fprintf(w, string(toscaGraph.Graph["ToscaDefinition"]))
}

// ViewToscaWorkflow is a http handler that output the SVG representation of the current tosca execution workflow
func (toscaGraph *ToscaGraph) ViewToscaWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=UTF-8")
	fmt.Fprintf(w, string(toscaGraph.Graph["ToscaWorkflow"]))
}
