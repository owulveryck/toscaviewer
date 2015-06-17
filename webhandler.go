package toscaviewer

// This is a basic example
// Thanks http://thenewstack.io/make-a-restful-json-api-go/ for the tutorial
import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

// UploadHandler courtesy of http://noypi-linux.blogspot.fr/2014/07/golang-web-server-basic-operatons-using.html
func (toscaGraph *ToscaGraph) UploadHandler(res http.ResponseWriter, req *http.Request) {
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
			log.Println("Got file:", hdr.Filename)

			outfile := os.Stdout
			/*
				// open destination
				var outfile *os.File
				if outfile, err = os.Create("./uploaded/" + hdr.Filename); nil != err {
					status = http.StatusInternalServerError
					return
				}
			*/
			// 32K buffer copy
			var written int64
			if written, err = io.Copy(outfile, infile); nil != err {
				status = http.StatusInternalServerError
				return
			}
			err = toscaGraph.ToscaDefinition.Parse(infile)
			res.Write([]byte(toscaGraph.ToscaDefinition.Bytes()))
			if err != nil {
				log.Println(err)
			}
			toscaGraph.Initialize()
			res.Write([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))))
		}
	}

}

func (toscaGraph ToscaGraph) ViewToscaYaml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text; charset=UTF-8")
	fmt.Fprintf(w, string(toscaGraph.Graph["ToscaYaml"]))
}
func (toscaGraph ToscaGraph) ViewToscaDefinition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=UTF-8")
	fmt.Fprintf(w, string(toscaGraph.Graph["ToscaDefinition"]))
}
func (toscaGraph ToscaGraph) ViewToscaWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=UTF-8")
	fmt.Fprintf(w, string(toscaGraph.Graph["ToscaWorkflow"]))
}
