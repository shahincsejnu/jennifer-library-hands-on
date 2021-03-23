// Parens renders a single item in parenthesis.
// Use for type conversion or to specify evaluation order

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Id("b").Op(":=").Index().Byte().Parens(Id("s"))

	fmt.Printf("%#v\n", c)

	c = Id("a").Op("/").Parens(Id("b").Op("+").Id("c"))

	fmt.Printf("%#v\n", c)
}