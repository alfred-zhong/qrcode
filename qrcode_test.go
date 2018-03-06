package main

import "testing"

func TestResolveFilePath(t *testing.T) {
	fp1 := ""
	p1 := resolveFilePath(fp1)
	if p1 == "" {
		t.Error("path expected not empty, got empty")
	}

	fp2 := "/tmp"
	p2 := resolveFilePath(fp2)
	if p2 != fp2 {
		t.Errorf("path expected %s, got %s", fp2, p2)
	}
}
