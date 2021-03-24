// DictFunc executes a func(Dict) to generate the value

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Id("a").Op(":=").Map(String()).String().Values(
		DictFunc(
			func(d Dict) {
				d[Lit("a")] = Lit("b")
				d[Lit("c")] = Lit("d")
			},
		),
	)

	fmt.Printf("%#v\n", c)
}

// Note: the items are ordered by key when rendered to ensure repeatable code