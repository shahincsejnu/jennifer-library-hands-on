// Struct render the keyword followed by a statement list enclosed by curly braces

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Id("c").Op(":=").Make(Chan().Struct())

	fmt.Printf("%#v\n", c)

	c = Type().Id("foo").Struct(
		List(Id("x"), Id("y")).Int(),
		Id("u").Float32(),
	)

	fmt.Printf("%#v\n", c)
}