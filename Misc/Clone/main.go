package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	// be careful when passing *Statement
	// consider the following...

	a := Id("a")
	c := Block(
		a.Call(),
		a.Call(),
	)

	fmt.Printf("%#v\n", c)

	// Id("a") returns a *Statement, which the Call() method appends to twice.
	// To avoid this, use Clone
	// Clone makes a copy of the Statement, so further tokens can be appended without affecting the original

	a = Id("a")
	c = Block(
		a.Clone().Call(),
		a.Clone().Call(),
		//a.Call(),
	)

	fmt.Printf("%#v\n", c)
}