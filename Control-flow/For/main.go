// For render the keyword followed by a semicolon separated list

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := For(Id("i").Op(":=").Lit(0), Id("i").Op("<").Lit(10), Id("i").Op("++")).Block(
		Qual("fmt", "Println").Call(Id("i")),
	)

	fmt.Printf("%#v\n", c)
}