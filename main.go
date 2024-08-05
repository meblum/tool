package main

import (
	"fmt"
	"io"
	"os"

	xj "github.com/basgys/goxml2json"
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

	if len(os.Args) == 1 {
		if err := convert(os.Stdin, os.Stdout); err != nil {
			fmt.Println(err)
		}
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if err := convert(f, os.Stdout); err != nil {
		fmt.Println(err)
	}
}
