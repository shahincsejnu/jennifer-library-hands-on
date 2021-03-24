// Empty adds an empty item
// Empty items render nothing but are followed by a separator in lists

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Id("a").Op(":=").Id("b").Index(Lit(1), Empty())

	fmt.Printf("%#v\n", c)
}