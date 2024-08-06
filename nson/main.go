package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

func reader() io.ReadCloser {
	if len(os.Args) == 1 {
		return os.Stdin
	}
	f, err := os.Open(strings.TrimSpace(os.Args[1]))
	if err != nil {
		log.Fatalf("open input file %q: %v", os.Args[1], err)
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
		log.Fatalf("unmarshal JSON: %v\n", err)
	}

	// Normalize the JSON
	normalizeJSON(jsonData)

	// Marshal the normalized JSON
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(jsonData); err != nil {
		log.Fatalf("marshal normalized JSON: %v\n", err)
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
