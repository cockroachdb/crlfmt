package test

// cgo 1
import "C"

// "C" imports without a docstring are meaningless; they can be removed

import "C" // bai

// cgo 2
import "C"

// this comment doesn't matter
import (
	"C" // bai
	// cgo 3
	"C"
	// cgo 4
	"C"
)

import (
	// cgo 5
	"C"
)

var _ = _
