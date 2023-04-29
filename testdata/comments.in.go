package test

// docstring should wrap to 80 chars test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2
// test3 test4 test5 test6 test7 test3 test4 test5 test6 test7 test8 test9 test1 test2 test3 test4 test5
func docstring() string

// docstring with a word over 80 chars cannot be wrapped
// verylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongword
// test1 test2 test3 test4 test5 test6 test7 test8 test9 test1 test2 verylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongwordverylongword
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
