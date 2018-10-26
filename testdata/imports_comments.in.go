package test

// import doc comment to be preserved
import (
  "a" // a line comment
)

// interleaved comment to be preserved

// import doc block comment to be removed
import (
	"b" // b line comment
	"c"
	// d doc comment
	"d"
	// e doc comment
	"e" // e line comment
)

// f doc comment
import "f" // f line comment

import "g" // g line comment

// import doc comment to be deleted; there's nothing sensible we can do with it
import (
  "h"
)


var _ = a.Foo
var _ = b.Foo
var _ = c.Foo
var _ = d.Foo
var _ = e.Foo
var _ = f.Foo
var _ = g.Foo
var _ = h.Foo
