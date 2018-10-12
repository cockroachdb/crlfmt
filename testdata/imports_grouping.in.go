package test

import "a"

import (
	"b"

	"github.com/cockroachdb/fake/c"
)

import (
	"elsewhere.com/fake/e"
	"github.com/fake/d"
)

var _ = a.Foo
var _ = b.Foo
var _ = c.Foo
var _ = d.Foo
var _ = e.Foo
