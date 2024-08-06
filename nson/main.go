package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// Unmarshal JSON into a dynamic structure
	var jsonData interface{}
	if err := json.NewDecoder(os.Stdin).Decode(&jsonData); err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling JSON: %v\n", err)
		return
	}

	// Normalize the JSON
	normalizeJSON(jsonData)

	// Marshal the normalized JSON
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(jsonData); err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling normalized JSON: %v\n", err)
		return
	}
}

func normalizeJSON(data interface{}) {
	switch v := data.(type) {
	case []interface{}:
		normalizeArray(v)
	case map[string]interface{}:
		normalizeObject(v)
	}
}

func normalizeArray(arr []interface{}) {
	allKeys := make(map[string]bool)

	// Collect all keys
	for _, item := range arr {
		if obj, ok := item.(map[string]interface{}); ok {
			for key := range obj {
				allKeys[key] = true
			}
		}
	}

	// Normalize each object in the array
	for _, item := range arr {
		if obj, ok := item.(map[string]interface{}); ok {
			for key := range allKeys {
				if _, exists := obj[key]; !exists {
					obj[key] = nil
				}
			}
			normalizeJSON(obj)
		}
	}
}

func normalizeObject(obj map[string]interface{}) {
	for _, value := range obj {
		normalizeJSON(value)
	}
}
