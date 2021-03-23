// Defs renders a statement list enclosed in parenthesis
// Use for definition list

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Const().Defs(
		Id("a").Op("=").Lit("a"),
		Id("b").Op("=").Lit("b"),
	)

	fmt.Printf("%#v\n", c)
}