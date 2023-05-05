package test

// docstring should wrap to 80 chars test1 test2 test3 test4 test5 test6 test7
// test8 test9 test1 test2
// test3 test4 test5 test6 test7 test3 test4 test5 test6 test7 test8 test9 test1
// test2 test3 test4 test5
//
// test3 test4 test5 test6 test7 test3 test4 test5 test6 test7 test8 test9 test1
// test2 test3 test4 test5
func docstring() string

// docstring with a word over 80 chars cannot be wrapped
// verylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongword
// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2
// verylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongword
// test1 test2
func docstringWithLongWords() string

// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2 test3 aaaaa
// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2 test3 bbbbb
// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2 test3 ccccc
func docstringWith80CharComments() string

// This is an explanation
// - foo
// - bar
func docstringWithBullets() string

/*
multiline comments are unchanged  multiline comments are unchanged multiline comments are unchanged
multiline comments are unchanged  multiline comments are unchanged  multiline comments are unchanged
multiline comments are unchanged  multiline comments are unchanged  multiline comments are unchanged
*/
func multilineComment() string

// docstring should wrap to 80 chars test1 test2 test3 test4 test5 test6 test7
// test8 test9 test1 test2
// test3 test4 test5 test6 test7 test3 test4 test5 test6 test7 test8 test9 test1
// test2 test3 test4 test5
type docstring string

// docstring with a word over 80 chars cannot be wrapped
// verylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongword
// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2
// verylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongword
// test1 test2
type longWords struct {
	s string
}

// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2 test3 aaaaa
// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2 test3 bbbbb
// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2 test3 ccccc
type eightyCharLines struct {
	s string
}

// This is an explanation
// - foo
// - bar
type bullets struct {
	s string
}

/*
multiline comments are unchanged  multiline comments are unchanged multiline comments are unchanged
multiline comments are unchanged  multiline comments are unchanged  multiline comments are unchanged
multiline comments are unchanged  multiline comments are unchanged  multiline comments are unchanged
*/
type multiline struct {
	s string
}

// docstring should wrap to 80 chars test1 test2 test3 test4 test5 test6 test7
// test8 test9 test1 test2
// test3 test4 test5 test6 test7 test3 test4 test5 test6 test7 test8 test9 test1
// test2 test3 test4 test5
const docstring = ""

// docstring with a word over 80 chars cannot be wrapped
// verylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongword
// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2
// verylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongword
// test1 test2
const longWords = "asdf"

// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2 test3 aaaaa
// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2 test3 bbbbb
// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2 test3 ccccc
const eightyCharLines = 80

// This is an explanation
// - foo
// - bar
const bullets = `foo
  bar
baz
`

/*
multiline comments are unchanged  multiline comments are unchanged multiline comments are unchanged
multiline comments are unchanged  multiline comments are unchanged  multiline comments are unchanged
multiline comments are unchanged  multiline comments are unchanged  multiline comments are unchanged
*/
const multiline = `foo
  bar
baz
`

// docstring should wrap to 80 chars test1 test2 test3 test4 test5 test6 test7
// test8 test9 test1 test2
// test3 test4 test5 test6 test7 test3 test4 test5 test6 test7 test8 test9 test1
// test2 test3 test4 test5
var docstring = struct{ string }{}

// docstring with a word over 80 chars cannot be wrapped
// verylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongword
// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2
// verylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongword
// test1 test2
var longWords = [10]int{}

// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2 test3 aaaaa
// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2 test3 bbbbb
// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2 test3 ccccc
var eightyCharLines = 80

// This is an explanation
// - foo
// - bar
var bullets = `foo
  bar
baz
`

/*
multiline comments are unchanged  multiline comments are unchanged multiline comments are unchanged
multiline comments are unchanged  multiline comments are unchanged  multiline comments are unchanged
multiline comments are unchanged  multiline comments are unchanged  multiline comments are unchanged
*/
var multiline = make(map[string]int)
