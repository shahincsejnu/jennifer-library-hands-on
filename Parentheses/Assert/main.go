// Assert renders a period followed by a single item enclosed by parenthesis
// Use for type assertions

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := List(Id("b"), Id("ok").Op(":=").Id("a").Assert(Bool()))

	fmt.Printf("%#v\n", c)
}