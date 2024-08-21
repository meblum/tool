package main

import (
	"io"
	"log"
	"os"

	"github.com/antonmedv/clipboard"
	"github.com/meblum/tool"
)

func main() {
	var s string
	r := tool.Reader()
	if r == nil {
		cs, err := clipboard.ReadAll()
		if err != nil {
			log.Fatalf("read from clipboard: %v", err)
		}
		s = cs
	} else {
		defer func() {
			if err := r.Close(); err != nil {
				log.Fatalf("close input stream: %v", err)
			}
		}()
		b, err := io.ReadAll(r)
		if err != nil {
			log.Fatalf("read input stream: %v", err)
		}
		s = string(b)
		// input provided, override clipboard
		if err := clipboard.WriteAll(s); err != nil {
			log.Fatalf("write to clipboard: %v", err)
		}
	}

	if _, err := os.Stdout.WriteString(s); err != nil {
		log.Fatalf("write to stdout: %v", err)
	}
}
