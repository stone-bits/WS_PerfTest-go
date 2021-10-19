package main

import (
	"testing"
)

func TestGetRespose(t *testing.T) {
	resp := getRespose()
	if resp == "" {
		t.Fatalf("response is empty")
	}
}
