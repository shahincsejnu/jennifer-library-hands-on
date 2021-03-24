// Line inserts a blank line

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Id("oka").Line()

	fmt.Printf("%#v\n", c)

	c = Id("okaoka")

	fmt.Printf("%#v\n", c)
}