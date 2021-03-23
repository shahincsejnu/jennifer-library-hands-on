// Select renders the select keyword

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Select().Block(
		Case(Id("c").Op("<-").Id("x")).Block(
			List(Id("x"), Id("y").Op("=").Id("y"), Id("x").Op("+").Id("y")),
			Return(),
		),
	)

	fmt.Printf("%#v\n", c)
}