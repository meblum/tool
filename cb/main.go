package main

import (
	"io"
	"log"
	"os"

	"github.com/antonmedv/clipboard"
	"github.com/mattn/go-isatty"
)

func getInput() (string, bool) {
	in := os.Stdin.Fd()
	isInTTY := isatty.IsTerminal(in) || isatty.IsCygwinTerminal(in)
	if !isInTTY {
		s, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("read stdin: %v", err)
		}
		return string(s), true
	}

	if len(os.Args) > 1 {
		s, err := os.ReadFile(os.Args[1])
		if err != nil {
			log.Fatalf("read file: %v", err)
		}
		return string(s), true
	}
	return "", false
}

func main() {
	s, ok := getInput()
	if ok {
		// input provided, override clipboard
		if err := clipboard.WriteAll(string(s)); err != nil {
			log.Fatalf("write to clipboard: %v", err)
		}
	} else {
		cs, err := clipboard.ReadAll()
		if err != nil {
			log.Fatalf("read from clipboard: %v", err)
		}
		s = cs
	}

	if _, err := os.Stdout.WriteString(s); err != nil {
		log.Fatalf("write to stdout: %v", err)
	}
}
