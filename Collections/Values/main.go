// Values renders a comma separated list enclosed by curly braces
// Use for slice or composite literals

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Index().String().Values(Lit("a"), Lit("b"))

	fmt.Printf("%#v\n", c)
}