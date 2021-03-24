// Commentf adds a comment, using a format string and a list of parameters

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	name := "foo"
	val := "bar"

	c := Id(name).Op(":=").Lit(val).Commentf("%s is the string \"%s\"", name, val)

	fmt.Printf("%#v\n", c)
}