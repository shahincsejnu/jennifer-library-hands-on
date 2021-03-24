// Lit renders a literal.
// Lit supports only built-in types (see in README.md)
// Passing any other type will panic

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Id("a").Op(":=").Lit("a")

	fmt.Printf("%#v\n", c)

	c = Id("a").Op(":=").Lit(1.5)

	fmt.Printf("%#v\n", c)
}