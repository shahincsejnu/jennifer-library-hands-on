// Id renders an identifier

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := If(Id("i").Op("==").Id("j")).Block(
		Return(Id("i")),
	)

	fmt.Printf("%#v", c)

	// as `c` is a *jen.Statement not a file so can not save it by c.Save as it is only for File.Save
}