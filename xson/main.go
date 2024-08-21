package main

import (
	"io"
	"log"
	"os"

	xj "github.com/basgys/goxml2json"
	"github.com/meblum/tool"
)

// convert reads xml from r and encodes it as json to w
func convert(r io.Reader, w io.Writer) error {
	var root xj.Node
	if err := xj.NewDecoder(r).Decode(&root); err != nil {
		return err
	}
	return xj.NewEncoder(w).Encode(&root)
}

func main() {

	r := tool.Reader()
	if r == nil {
		log.Fatal("input not provided")
	}
	defer func() {
		if err := r.Close(); err != nil {
			log.Fatalf("close input stream: %v", err)
		}
	}()

	if err := convert(r, os.Stdout); err != nil {
		log.Fatalf("convert xml: %v", err)
	}
}
