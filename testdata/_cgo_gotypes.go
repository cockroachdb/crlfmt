//go:cgo_ldflag "-static-libgcc"
//go:cgo_ldflag "-static-libstdc++"
//go:cgo_ldflag "-Wl,-unresolved-symbols=ignore-all"
// Created by cgo - DO NOT EDIT

package rocksdb

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno

func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool

//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, ...interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})

func someSignatureThatIs100Chars____________________________________(someArg, someOtherArg string) {
}

func someSignatureThatIs101Chars_____________________________________(
	someArg, someOtherArg string,
) {
}

func someSignatureWithResults(someArg, someOtherArg string) (string, string, string, string, bool) {
}

func someSignatureWithLongResults(
	someArg, someOtherArg string,
) (string, string, string, string, string) {
}

func someSigWithLongArgs(
	someArg string,
	someOtherArg string,
	someLoooooooooooooooooooooooooooooooooooooooooooooooooooooooog int,
) {
}

func someSigWithLongArgsAndElidedTypeShorthand(
	someArg, someOtherArg string,
	someLoooooooooooooooooooooooooooooooooooooooooooooooooooooooog int,
) {
}
