package main

import (
	"flag"
	"net/http"
)

var (
	port = flag.String("port", "127.0.0.1:9000", "Define what TCP port to bind to")
	root = flag.String("root", ".", "Define the root filesystem path")
)

func main() {
	flag.Parse()

	print("Serving '", *root, "' on http://", *port, "\n")
	panic(http.ListenAndServe(*port, http.FileServer(http.Dir(*root))))
}
