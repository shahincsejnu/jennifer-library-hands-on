// All constructs that accept a variadic list of items are paired with GroupFunc functions that accept a func(*Group)
// Use for embedding logic

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Id("numbers").Op(":=").Index().Id("int").ValuesFunc(
		func(g *Group) {
			for i := 0; i <= 5; i++ {
				g.Lit(i)
			}
		},
	)

	fmt.Printf("%#v\n", c)


	increment := true
	name := "a"

	c = Func().Id(name).Params().BlockFunc(
		func(g *Group) {
			g.Id("a").Op("=").Lit(1)
			if increment {
				g.Id(name).Op("++")
			} else {
				g.Id(name).Op("--")
			}
		},
	)

	fmt.Printf("%#v\n", c)
}