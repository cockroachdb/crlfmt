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
	"strconv"
	"strings"

	"github.com/cockroachdb/gostdlib/go/format"
	"github.com/cockroachdb/gostdlib/x/tools/imports"
	"github.com/cockroachdb/ttycolor"
)

var (
	wrap         = flag.Int("wrap", 100, "column to wrap at")
	tab          = flag.Int("tab", 8, "tab width for column calculations")
	overwrite    = flag.Bool("w", false, "overwrite modified files")
	fast         = flag.Bool("fast", false, "skip running goimports")
	groupImports = flag.Bool("groupimports", false, "group imports by type")
	printDiff    = flag.Bool("diff", true, "print diffs")
	ignore       = flag.String("ignore", "", "regex matching files to skip")
)

var (
	red   = string(ttycolor.StdoutProfile[ttycolor.Red])
	green = string(ttycolor.StdoutProfile[ttycolor.Green])
	reset = string(ttycolor.StdoutProfile[ttycolor.Reset])
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		content, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		*overwrite = true
		*printDiff = false
		_, out, err := checkBuf("<standard input>", content)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if _, err := out.WriteTo(os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	if flag.NArg() > 1 {
		fmt.Println("must specify exactly one path argument (or zero for stdin)")
		os.Exit(1)
	}

	root, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		fmt.Printf("Error finding absolute path: %s", err)
		os.Exit(1)
	}
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		fmt.Printf("Error following symlinks in input path: %s", err)
		os.Exit(1)
	}

	var ignoreRE *regexp.Regexp
	if len(*ignore) > 0 {
		ignoreRE, err = regexp.Compile(*ignore)
		if err != nil {
			fmt.Printf("Error compiling ignore regexp: %s", err)
			os.Exit(1)
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

func maybePrintDiff(where token.Position, old, new []byte) {
	if *printDiff {
		fmt.Printf("%s:%d\n", where.Filename, where.Line)
		if old != nil {
			for _, line := range bytes.Split(old, []byte{'\n'}) {
				fmt.Printf("%s-%s%s\n", red, line, reset)
			}
		}
		if new != nil {
			for _, line := range bytes.Split(new, []byte{'\n'}) {
				fmt.Printf("%s+%s%s\n", green, line, reset)
			}
		}
		fmt.Println()
	}
}

func importPath(spec ast.Spec) string {
	if t, err := strconv.Unquote(spec.(*ast.ImportSpec).Path.Value); err == nil {
		return t
	}
	return ""
}

func concat(bs ...[]byte) (out []byte) {
	for _, b := range bs {
		out = append(out, b...)
	}
	return
}

func checkPath(path string) (int, error) {
	var diffs int

	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}

	diffs, output, err := checkBuf(path, fileBytes)
	if err != nil {
		return 0, err
	}

	if *overwrite && diffs > 0 {
		err = ioutil.WriteFile(path, output.Bytes(), 0)
		if err != nil {
			return 0, err
		}
	}

	return diffs, nil
}

func checkBuf(path string, fileBytes []byte) (int, *bytes.Buffer, error) {
	output := new(bytes.Buffer)
	var diffs int
	if !*fast {
		// Run goimports, which also runs gofmt.
		importOpts := imports.Options{
			AllErrors:  true,
			Comments:   true,
			TabIndent:  false,
			TabWidth:   *tab,
			FormatOnly: false,
		}
		newFileBytes, err := imports.Process(path, fileBytes, &importOpts)
		if err != nil {
			return 0, output, err
		}
		// If goimports made any change, count that as a diff so the file
		// can be overwritten at the end.
		if *printDiff && bytes.Compare(fileBytes, newFileBytes) != 0 {
			fmt.Printf("%s: import list mismatch\n", path)
			diffs = 1
		}
		fileBytes = newFileBytes
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, fileBytes, parser.AllErrors|parser.ParseComments)
	if err != nil {
		return 0, output, err
	}

	fileSlice := func(beg token.Pos, end token.Pos) []byte {
		return fileBytes[fset.Position(beg).Offset:fset.Position(end).Offset]
	}

	var (
		stdlibImports     []*ast.ImportSpec
		thirdPartyImports []*ast.ImportSpec
		firstPartyImports []*ast.ImportSpec
	)

	for _, imp := range f.Imports {
		switch impPath := importPath(imp); {
		case impPath == "C":
			continue
		case !strings.Contains(impPath, "."):
			stdlibImports = append(stdlibImports, imp)
		case strings.HasPrefix(impPath, "github.com/cockroachdb"):
			firstPartyImports = append(firstPartyImports, imp)
		default:
			thirdPartyImports = append(thirdPartyImports, imp)
		}
	}

	renderImports := func(b *bytes.Buffer, imps []*ast.ImportSpec) {
		for _, imp := range imps {
			if imp.Doc != nil {
				b.WriteByte('\t')
				b.Write(fileSlice(imp.Doc.Pos(), imp.Doc.End()))
				b.WriteByte('\n')
			}
			b.WriteByte('\t')
			if imp.Name != nil {
				b.WriteString(imp.Name.String())
				b.WriteByte(' ')
			}
			b.WriteString(imp.Path.Value)
			if imp.Comment != nil {
				b.WriteByte(' ')
				b.Write(fileSlice(imp.Comment.Pos(), imp.Comment.End()))
			}
			b.WriteByte('\n')
		}
		if len(imps) > 0 {
			b.WriteByte('\n')
		}
	}

	var curFunc bytes.Buffer
	var seenImportDecl bool
	lastPos := token.NoPos
outer:
	for _, d := range f.Decls {
		if g, ok := d.(*ast.GenDecl); ok && g.Tok == token.IMPORT && *groupImports {
			for _, spec := range g.Specs {
				if importPath(spec) == "C" {
					// "C" is a very special import as it causes the comment
					// on the gen decl to be interpreted as C code that is
					// compiled and linked into the binary. Best to just leave
					// this import block alone.
					continue outer
				}
			}

			numImports := len(stdlibImports) + len(thirdPartyImports) + len(firstPartyImports)
			if seenImportDecl || numImports == 0 {
				// We've already output all of the imports in the first import
				// block. All other import blocks should be removed.
				//
				// If the import block is surrounded by blank lines, remove
				// the blank lines too.
				startPos, endPos := g.Pos(), g.End()
				if off := fset.Position(startPos).Offset; off-2 >= 0 && fileBytes[off-1] == '\n' && fileBytes[off-2] == '\n' {
					startPos = fset.File(startPos).Pos(off - 1)
				}
				if off := fset.Position(g.End()).Offset; off+1 < len(fileBytes) && fileBytes[off] == '\n' && fileBytes[off+1] == '\n' {
					endPos = fset.File(endPos).Pos(off + 1)
				}
				maybeWrite(output, fileSlice(lastPos, startPos))
				lastPos = endPos
				maybePrintDiff(fset.Position(startPos), fileSlice(startPos, lastPos), nil)
				diffs++
				continue
			}

			// This is the first import block. Render all of the imports grouped
			// by type, sorted alphabetically within each group, and with blank
			// lines between each group, unless there aren't any imports to
			// render.
			maybeWrite(output, fileSlice(lastPos, g.Pos()))
			lastPos = g.End()
			var importBuf bytes.Buffer
			importBuf.WriteString("import ")
			if numImports > 1 {
				importBuf.WriteString("(\n")
			}
			renderImports(&importBuf, stdlibImports)
			renderImports(&importBuf, thirdPartyImports)
			renderImports(&importBuf, firstPartyImports)
			importBuf.Truncate(importBuf.Len() - 1) // trim trailing blank line
			if numImports > 1 {
				importBuf.WriteByte(')')
			}
			newImports, err := format.Source(importBuf.Bytes())
			if err != nil {
				return 0, nil, fmt.Errorf("grouping imports for %s: %s", path, err)
			}
			if numImports == 1 {
				newImports = newImports[:len(newImports)-1] // trim trailing newline
			}
			if !bytes.Equal(fileSlice(g.Pos(), g.End()), newImports) {
				maybePrintDiff(fset.Position(g.Pos()), fileSlice(g.Pos(), g.End()), newImports)
				diffs++
			}
			maybeWrite(output, newImports)
			seenImportDecl = true
		}
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
			singleLineLen := colOffset + len(paramsJoined) + len(funcMid) + len(resultsJoined) + len(funcEnd) + brace
			if singleLineLen <= *wrap {
				curFunc.Write(paramsJoined)
				curFunc.WriteString(funcMid)
				curFunc.Write(resultsJoined)
				curFunc.WriteString(funcEnd)
			} else {
				// we're into wrapping, so the return types block usually starts on own
				// line intended by `tab`.
				resTypeStartingCol := *tab
				if len(params.List) == 0 {
					// special case: if we have no params, the res type starts on the same
					// line rather than on its own.
					resTypeStartingCol = colOffset
				} else if *tab+len(paramsJoined)+len(paramsLineEndComma) <= *wrap {
					fmt.Fprintf(&curFunc, "\n\t%s,\n", paramsJoined)
				} else {
					for _, param := range params.List {
						fmt.Fprintf(&curFunc, "\n\t%s,", fileSlice(param.Pos(), param.End()))
					}
					curFunc.WriteByte('\n')
				}
				curFunc.WriteString(funcMid)
				singleLineRetunsLen := resTypeStartingCol + len(funcMid) + len(resultsJoined) + len(funcEnd) + brace
				if singleLineRetunsLen <= *wrap {
					curFunc.Write(resultsJoined)
					curFunc.WriteString(funcEnd)
				} else {
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
				prefix := fileSlice(d.Pos(), opening)
				maybePrintDiff(fset.Position(d.Pos()), concat(prefix, oldFunc), concat(prefix, curFunc.Bytes()))
				diffs++
				maybeWrite(output, curFunc.Bytes())
			} else {
				maybeWrite(output, oldFunc)
			}
		}
	}
	maybeWrite(output, fileBytes[fset.Position(lastPos).Offset:])
	return diffs, output, nil
}
