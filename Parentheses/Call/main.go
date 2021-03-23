// Call renders a comma separated list enclosed by parenthesis
// Use for function calls

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Qual("fmt", "Printf").Call(Lit("%#v: %T\n"), Id("a"), Id("b"))

	fmt.Printf("%#v", c)
}