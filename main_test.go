package main

import (
	"io/ioutil"
	"testing"
)

var inDir = "testdata/"

func TestCheckPath(t *testing.T) {
	files, err := ioutil.ReadDir(inDir)
	if err != nil {
		t.Error(err)
	}
	for _, file := range files {
		t.Run(file.Name(), func(t *testing.T) {
			inPath := inDir + file.Name()

			diffs, err := checkPath(inPath)
			if err != nil {
				t.Error(err)
			}
			if diffs != 0 {
				t.Errorf("expected 0 diffs but got %d", diffs)
			}
		})
	}
}
