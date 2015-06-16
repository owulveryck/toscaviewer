# toscaviewer
A graphical representation of a topology described in TOSCA

# What's working
If you have [graphviz](http://www.graphviz.org) installed, you can launch the main.go in toscaviewer and connect to [localhost](http://localhost:8080/svg)

```
 export GOPATH=somepath
 go get github.com/owulveryck/toscaviewer
 cd $GOPATH/src/github.com/owulveryck/toscaviewer/toscaviewer
 go run main.go
```

# Depedencies
* The [graphviz](http://www.graphviz.org) program to generate the svg
* bootstrap (included)
* [http://www.gorillatoolkit.org/pkg/mux](the gorilla web toolkit) (go gettable)
