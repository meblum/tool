package main

import (
	"io"
	"log"
	"os"

	"github.com/meblum/tool"
	"github.com/twpayne/go-xmlstruct"
)

func convert(r io.Reader, w io.Writer) {
	g := xmlstruct.NewGenerator(
		xmlstruct.WithPackageName(""),
		xmlstruct.WithHeader(""),
		xmlstruct.WithPackageName(""),
		xmlstruct.WithImports(false),
		xmlstruct.WithNamedRoot(true),
		xmlstruct.WithNamedTypes(true),
		xmlstruct.WithEmptyElements(false),
		xmlstruct.WithTopLevelAttributes(true),
	)
	if err := g.ObserveReader(r); err != nil {
		log.Fatalf("read input: %v", err)
	}
	out, err := g.Generate()
	if err != nil {
		log.Fatalf("generate structs: %v", err)
	}
	if _, err := w.Write(out); err != nil {
		log.Fatalf("write output: %v", err)
	}
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
	convert(r, os.Stdout)
}
