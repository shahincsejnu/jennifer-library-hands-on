// simple keywords, predeclared identifiers and built-in functions are self explanatory

// Keywords: Break, Chan, Const, Continue, Default, Defer, Else, Fallthrough, Func, Go, Goto,
// Range, Select, Type, Var

// Functions: Append, Cap, Close, Complex, Copy, Delete, Imag, Len, Make, New, Panic, Print, Println
// Real, Recover

// Types: Bool, Byte, Complex64, Complex128, Error, Float32, Float64, Int, Int8, Int16, Int32, Int64, Rune
// String, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr

// Constants: True, False, Iota, Nil

// Helpers: Err

// Built-in functions take a list of parameters and render them appropriately

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Id("a").Op("=").Append(Id("a"), Id("b").Op("..."))

	fmt.Printf("%#v", c)
}