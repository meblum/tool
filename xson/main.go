package main

import (
	"io"
	"log"
	"os"
	"strings"

	xj "github.com/basgys/goxml2json"
	"github.com/mattn/go-isatty"
)

// convert reads xml from r and encodes it as json to w
func convert(r io.Reader, w io.Writer) error {
	var root xj.Node
	if err := xj.NewDecoder(r).Decode(&root); err != nil {
		return err
	}
	return xj.NewEncoder(w).Encode(&root)
}

func reader() io.ReadCloser {
	istty := isatty.IsTerminal(os.Stdin.Fd()) || isatty.IsCygwinTerminal(os.Stdin.Fd())
	if !istty {
		return os.Stdin
	}
	if len(os.Args) < 2 {
		log.Fatal("input not provided")
	}
	f, err := os.Open(strings.TrimSpace(os.Args[1]))
	if err != nil {
		log.Fatalf("open file: %v", err)
	}
	return f
}

func main() {

	r := reader()
	defer func() {
		if err := r.Close(); err != nil {
			log.Fatalf("close input stream: %v", err)
		}
	}()

	if err := convert(r, os.Stdout); err != nil {
		log.Fatalf("convert xml: %v", err)
	}
}
