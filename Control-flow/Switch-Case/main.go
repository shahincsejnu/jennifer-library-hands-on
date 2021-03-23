// Switch, Select, Case, and Block are used to build switch or select statements
// Switch renders the keyword followed by a semicolon separated list

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Switch(Id("value").Dot("Kind").Call()).Block(
		Case(Qual("reflect", "Float32"), Qual("reflect", "Float64")).Block(
			Return(Lit("float")),
		),
		Case(Qual("reflect", "Bool")).Block(
			Return(Lit("bool")),
		),
		Case(Qual("reflect", "Uintptr")).Block(
			Fallthrough(),
		),
		Default().Block(
			Return(Lit("none")),
		),
	)

	fmt.Printf("%#v\n", c)
}