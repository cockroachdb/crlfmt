// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckPath(t *testing.T) {
	defer func(old bool) { *printDiff = old }(*printDiff)
	*printDiff = false

	files, err := filepath.Glob("testdata/*.in.go")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			inBytes, err := ioutil.ReadFile(file)
			if err != nil {
				t.Fatal(err)
			}
			expBytes, err := ioutil.ReadFile(strings.Replace(file, ".in.go", ".out.go", -1))
			if err != nil {
				t.Fatal(err)
			}
			output, err := checkBuf(file, inBytes)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(t, string(expBytes), string(output))
		})
	}
}
