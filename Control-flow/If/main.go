// If render the keyword followed by a semicolon separated list

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := If(Err().Op(":=").Id("a").Call(), Err().Op("!=").Nil()).Block(
		Return(Err()),
	)

	fmt.Printf("%#v\n", c);
}