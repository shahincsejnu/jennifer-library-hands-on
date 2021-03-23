// Op renders the provided operator / token

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Id("a").Op(":=").Id("b").Call()

	fmt.Printf("%#v\n", c)

	c = Id("a").Op("=").Op("*").Id("b")

	fmt.Printf("%#v\n", c)

	c = Id("a").Call(Id("b").Op("..."))

	fmt.Printf("%#v\n", c)

	c = If(Parens(Id("a").Op("||").Id("b")).Op("&&").Id("c")).Block(

	)

	fmt.Printf("%#v", c)
}