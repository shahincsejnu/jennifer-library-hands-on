// NewFile creates a new file, with the specified package name

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("main")

	f.Func().Id("main").Params().Block(
		Qual("fmt", "Println").Call(Lit("Hello, World")),
	)

	fmt.Printf("%#v\n", f)
}