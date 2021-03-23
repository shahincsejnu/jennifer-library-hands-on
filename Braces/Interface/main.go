// Interface render the keyword followed by a statement list enclosed by curly braces

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Var().Id("a").Interface()

	fmt.Printf("%#v\n", c)

	c = Type().Id("a").Interface(
		Id("b").Call().Id("string"),
	)

	fmt.Printf("%#v\n", c)
}