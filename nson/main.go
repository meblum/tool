package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

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

	// Unmarshal JSON into a dynamic structure
	var jsonData interface{}
	if err := json.NewDecoder(r).Decode(&jsonData); err != nil {
		log.Fatalf("unmarshal JSON: %v", err)
	}

	// Normalize the JSON
	normalizeJSON(jsonData)

	// Marshal the normalized JSON
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(jsonData); err != nil {
		log.Fatalf("marshal normalized JSON: %v", err)
		return
	}
}

func normalizeJSON(data interface{}) {
	switch v := data.(type) {
	case []interface{}:
		allKeys := make(map[string]bool)

		// Collect all keys
		for _, item := range v {
			if obj, ok := item.(map[string]interface{}); ok {
				for key := range obj {
					allKeys[key] = true
				}
			}
		}

		// Normalize each object in the array
		for _, item := range v {
			if obj, ok := item.(map[string]interface{}); ok {
				for key := range allKeys {
					if _, exists := obj[key]; !exists {
						obj[key] = nil
					}
				}
				normalizeJSON(obj)
			}
		}
	case map[string]interface{}:
		for _, value := range v {
			normalizeJSON(value)
		}
	}
}
