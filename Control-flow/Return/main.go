// Return renders the keyword followed by a comma separated list

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Return(Id("a"), Id("b"))

	fmt.Printf("%#v\n", c)
}