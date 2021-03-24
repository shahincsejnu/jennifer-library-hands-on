// Do calls the provided function with the statement as a parameter
// Use for embedding logic

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	check := true

	c := Id("a").Op(":=").Do(
		func(s *Statement) {
			if check {
				s.Map(String()).String().Values()
			}
		},
	)

	fmt.Printf("%#v\n", c)

	f := func(name string, isMap bool) *Statement {
		return Id(name).Op(":=").Do(
			func(s *Statement) {
				if isMap {
					s.Map(String()).String()
				} else {
					s.Index().String()
				}
			},
		).Values()
	}

	fmt.Printf("%#v\n", f("a", true))
	fmt.Printf("%#v\n", f("b", false))
}