package test

// It's weird to import "C" before other things, but it's legal.
import "C"

import (
	_ "a"
	_ "b"
)
