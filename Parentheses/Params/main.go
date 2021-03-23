// Params renders a comma separated list enclosed by parenthesis
// Use for function parameters and method receivers

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Func().Params(Id("a").Id("A")).Id("foo").Params(Id("b"), Id("c").String()).String().Block(
		Return(Id("b").Op("+").Id("c")),
	)

	fmt.Printf("%#v\n", c)
}