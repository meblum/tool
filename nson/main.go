package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/meblum/tool"
)

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
