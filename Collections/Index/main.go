// Index renders a colon separated list enclosed by square brackets
// Use for array / slice indexes and definitions

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Var().Id("a").Index().String()

	fmt.Printf("%#v\n", c)

	c = Id("a").Op(":=").Id("b").Index(Lit(0), Lit(1))

	fmt.Printf("%#v\n", c)

	c = Id("a").Op(":=").Id("b").Index(Lit(1), Empty())

	fmt.Printf("%#v\n", c)
}