// Note that the output here is bogus due to https://go.dev/issues/48688.

package test

// import doc comment to be preserved
import (
	"a"
	"b"
	"c"
	"d"
	"e"
	"f"
	"g"
	"h" // a line comment
)

// interleaved comment to be preserved

// import doc block comment to be removed

// b line comment

// d doc comment

// e doc comment
// e line comment

// f doc comment
// f line comment

// g line comment

// import doc comment to be deleted; there's nothing sensible we can do with it

var _ = a.Foo
var _ = b.Foo
var _ = c.Foo
var _ = d.Foo
var _ = e.Foo
var _ = f.Foo
var _ = g.Foo
var _ = h.Foo
