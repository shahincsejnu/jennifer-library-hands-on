// List renders a comma separated list
// Use for multiple return functions

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := List(Id("a"), Id("b"), Err()).Op(":=").Id("b").Call()

	fmt.Printf("%#v", c)
}

