// Block renders a statement list enclosed by curly braces
// User for code blocks

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Func().Id("foo").Call().Id("string").Block(
		Id("a").Op("=").Id("b"),
		Id("b").Op("++"),
		Return(Id("b")),
	)

	fmt.Printf("%#v\n", c)

	c = If(Id("a").Op(">").Lit(10)).Block(
		Id("a").Op("=").Id("a").Op("/").Lit(2),
	)

	fmt.Printf("%#v\n", c);
}