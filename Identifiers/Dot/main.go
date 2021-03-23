// Dot renders a period (.) followed by an identifier
// Use for fields and selectors

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Qual("a.b/c", "Foo").Call().Dot("Bar").Index(Lit(0)).Dot("Baz")

	fmt.Printf("%#v", c)
}