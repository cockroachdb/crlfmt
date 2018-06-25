// Copyright 2016 The Cockroach Authors.
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
//
// Author: Daniel Harrison (daniel.harrison@gmail.com)
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/cockroachdb/ttycolor"
	"golang.org/x/tools/imports"
)

var (
	wrap      = flag.Int("wrap", 100, "column to wrap at")
	tab       = flag.Int("tab", 8, "tab width for column calculations")
	overwrite = flag.Bool("w", false, "overwrite modified files")
	ignore    = flag.String("ignore", "", "regex matching files to skip")
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Println("missing argument: filepath")
		return
	}

	root, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		fmt.Printf("Error finding absolute path: %s", err)
		return
	}
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		fmt.Printf("Error following symlinks in input path: %s", err)
		return
	}

	var ignoreRE *regexp.Regexp
	if len(*ignore) > 0 {
		ignoreRE, err = regexp.Compile(*ignore)
		if err != nil {
			fmt.Printf("Error compiling ignore regexp: %s", err)
			return
		}
	}

	var diffs int
	err = filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error during filesystem walk: %v\n", err)
			return nil
		}
		if ignoreRE != nil && ignoreRE.MatchString(path) {
			return nil
		}
		if fi.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		pathDiffs, err := checkPath(path)
		if err != nil {
			return err
		}
		diffs += pathDiffs
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during walk:\n%s\n", err)
		os.Exit(1)
	}
	if diffs > 0 {
		fmt.Printf("Found %d diffs\n", diffs)
	}
}

func maybeWrite(output *bytes.Buffer, b []byte) {
	if *overwrite {
		output.Write(b)
	}
}

func getColors() (string, string, string) {
	return string(ttycolor.StdoutProfile[ttycolor.Red]),
		string(ttycolor.StdoutProfile[ttycolor.Green]),
		string(ttycolor.StdoutProfile[ttycolor.Reset])
}

func checkPath(path string) (int, error) {
	fset := token.NewFileSet()

	src, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}

	// Run goimports, which also runs gofmt.
	importOpts := imports.Options{
		AllErrors:  true,
		Comments:   true,
		TabIndent:  false,
		TabWidth:   *tab,
		FormatOnly: false,
	}
	fileBytes, err := imports.Process(path, src, &importOpts)
	if err != nil {
		return 0, err
	}
	var diffs int

	// If goimports made any change, count that as a diff so the file
	// can be overwritten at the end.
	if bytes.Compare(fileBytes, src) != 0 {
		fmt.Printf("%s: import list mismatch\n", path)
		diffs = 1
	}

	// Load the AST from the output of goimports.
	f, err := parser.ParseFile(fset, path, fileBytes, parser.AllErrors)
	if err != nil {
		return 0, err
	}

	fileSlice := func(beg token.Pos, end token.Pos) []byte {
		return fileBytes[fset.Position(beg).Offset:fset.Position(end).Offset]
	}

	var curFunc bytes.Buffer
	output := new(bytes.Buffer)

	red, green, reset := getColors()

	lastPos := token.NoPos
	for _, d := range f.Decls {
		if f, ok := d.(*ast.FuncDecl); ok {
			params := f.Type.Params
			results := f.Type.Results

			opening := params.Pos() + 1
			closing := f.Type.End()
			// f.Body is nil if the FuncDecl is a forward declaration.
			if f.Body != nil {
				closing = f.Body.Pos() + 1
			}

			maybeWrite(output, fileBytes[fset.Position(lastPos).Offset:fset.Position(opening).Offset])
			lastPos = closing

			var paramsBuf bytes.Buffer
			if params != nil {
				paramsPrefix := ""
				for _, f := range params.List {
					paramsBuf.WriteString(paramsPrefix)
					paramsBuf.Write(fileSlice(f.Pos(), f.End()))
					paramsPrefix = ", "
				}
			}
			paramsJoined := paramsBuf.Bytes()

			// Final comma needed if params are written out onto their own single line.
			const paramsLineEndComma = `,`

			var resultsBuf bytes.Buffer
			if results != nil {
				resultsPrefix := ""
				for _, f := range results.List {
					resultsBuf.WriteString(resultsPrefix)
					resultsBuf.Write(fileSlice(f.Pos(), f.End()))
					resultsPrefix = ", "
				}
			}
			resultsJoined := resultsBuf.Bytes()

			funcMid := `) (`
			funcEnd := `)`
			if results == nil || len(results.List) == 0 {
				funcMid = `)`
				funcEnd = ``
			} else if len(results.List) == 1 && len(results.List[0].Names) == 0 {
				funcMid = `) `
				funcEnd = ``
			}

			brace := 0
			if f.Body != nil {
				brace = len(` {`)
			}

			curFunc.Reset()
			// colOffset - 1 accounts for `func (r *foo) bar(`
			colOffset := fset.Position(opening).Column - 1
			if colOffset+len(paramsJoined)+len(funcMid)+len(resultsJoined)+len(funcEnd)+brace <= *wrap {
				curFunc.Write(paramsJoined)
				curFunc.WriteString(funcMid)
				curFunc.Write(resultsJoined)
				curFunc.WriteString(funcEnd)
			} else {
				if len(params.List) == 0 {
					// pass
				} else if *tab+len(paramsJoined)+len(paramsLineEndComma) <= *wrap {
					fmt.Fprintf(&curFunc, "\n\t%s,\n", paramsJoined)
				} else {
					for _, param := range params.List {
						fmt.Fprintf(&curFunc, "\n\t%s,", fileSlice(param.Pos(), param.End()))
					}
					curFunc.WriteByte('\n')
				}
				if *tab+len(funcMid)+len(resultsJoined)+len(funcEnd)+brace <= *wrap {
					curFunc.WriteString(funcMid)
					curFunc.Write(resultsJoined)
					curFunc.WriteString(funcEnd)
				} else {
					curFunc.WriteString(funcMid)
					for _, result := range results.List {
						fmt.Fprintf(&curFunc, "\n\t%s,", fileSlice(result.Pos(), result.End()))
					}
					curFunc.WriteByte('\n')
					curFunc.WriteString(funcEnd)
				}
			}
			curFunc.Write(fileBytes[fset.Position(f.Type.End()).Offset:fset.Position(closing).Offset])

			oldFunc := fileSlice(opening, closing)
			if !bytes.Equal(oldFunc, curFunc.Bytes()) {
				prefix := string(fileBytes[fset.Position(d.Pos()).Offset:fset.Position(opening).Offset])
				fmt.Printf("%s:%d\n", path, fset.Position(d.Pos()).Line)
				for _, line := range strings.Split(prefix+string(oldFunc), "\n") {
					fmt.Printf("%s-%s%s\n", red, line, reset)
				}
				for _, line := range strings.Split(prefix+curFunc.String(), "\n") {
					fmt.Printf("%s+%s%s\n", green, line, reset)
				}
				fmt.Print("\n")
				diffs++
				maybeWrite(output, curFunc.Bytes())
			} else {
				maybeWrite(output, oldFunc)
			}
		}
	}
	maybeWrite(output, fileBytes[fset.Position(lastPos).Offset:])

	if *overwrite && diffs > 0 {
		err = ioutil.WriteFile(path, output.Bytes(), 0)
		if err != nil {
			return 0, err
		}
	}

	return diffs, nil
}
