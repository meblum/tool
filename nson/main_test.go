package main

import (
	"encoding/json"
	"os"
	"testing"
)

func Test_normalizeJSON(t *testing.T) {
	in, out := inOut(t)
	c := convert(t, in)
	if string(c) != string(out) {
		t.Errorf("expected %s, got %s", out, c)
	}
}

func inOut(t *testing.T) ([]byte, []byte) {
	t.Helper()
	in, err := os.ReadFile("testdata/in.json")
	if err != nil {
		t.Errorf("open in.json: %v", err)
	}
	out, err := os.ReadFile("testdata/out.json")
	if err != nil {
		t.Errorf("open out.json: %v", err)
	}
	return in, out
}

func convert(t *testing.T, in []byte) []byte {
	t.Helper()
	var i interface{}
	if err := json.Unmarshal(in, &i); err != nil {
		t.Errorf("unmarshal input: %v", err)
	}
	normalizeJSON(i)
	b, err := json.Marshal(i)
	if err != nil {
		t.Errorf("marshal conversion: %v", err)
	}
	return b
}
