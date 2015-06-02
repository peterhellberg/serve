package main

import (
	"flag"
	"net/http"
)

var (
	addr = flag.String("addr", "0.0.0.0:9000", "Define what TCP address (host:port) to bind to")
	root = flag.String("root", ".", "Define the root filesystem path")
)

func main() {
	flag.Parse()

	go func() {
		print("Serving '", *root, "' on http://", *addr, "\n")
	}()

	if err := http.ListenAndServe(*addr, http.FileServer(http.Dir(*root))); err != nil {
		print(err.Error() + "\n")
	}
}
