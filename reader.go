package tool

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

func Reader() io.ReadCloser {
	istty := isatty.IsTerminal(os.Stdin.Fd()) || isatty.IsCygwinTerminal(os.Stdin.Fd())
	if !istty {
		return os.Stdin
	}
	if len(os.Args) > 1 {
		f, err := os.Open(strings.TrimSpace(os.Args[1]))
		if err != nil {
			log.Fatalf("open file: %v", err)
		}
		return f
	}
	return nil
}
