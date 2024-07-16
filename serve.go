package main

import (
	"flag"
	"net/http"
	"os"
)

type input struct {
	addr string
	root string
}

func parseInput(args []string) (input, error) {
	var in input

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	flags.StringVar(&in.addr, "addr", "localhost:9000", "Define what TCP address (host:port) to bind to")
	flags.StringVar(&in.root, "root", ".", "Define the root filesystem path")

	return in, flags.Parse(args[1:])
}

func main() {
	if err := run(os.Args); err != nil {
		print(err.Error() + "\n")
	}
}

func run(args []string) error {
	in, err := parseInput(args)
	if err != nil {
		return err
	}

	print("Serving '", in.root, "' on http://", in.addr, "\n")

	fs := http.FileServer(http.Dir(in.root))

	return http.ListenAndServe(in.addr, fs)
}
