package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	xj "github.com/basgys/goxml2json"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	j, err := xj.Convert(bytes.NewReader(input))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if _, err := io.Copy(os.Stdout, j); err != nil {
		fmt.Println(err.Error())
		return
	}
}
