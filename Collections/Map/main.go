// Map renders the keyword followed by a single item enclosed by square brackets
// User for map definitions

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Id("a").Op(":=").Map(Id("string")).Id("string").Values()

	fmt.Printf("%#v\n", c)
}