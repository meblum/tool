package main

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
	"github.com/twpayne/go-xmlstruct"
)

func convert(r io.Reader, w io.Writer) {
	g := xmlstruct.NewGenerator(
		xmlstruct.WithPackageName(""),
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
	convert(r, os.Stdout)
}
