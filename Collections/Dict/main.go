// Dict renders as key/value pairs
// Use with Values for map or composite literals

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Map(String()).String().Values(
		Dict{
			Lit("a"): Lit("b"),
			Lit("c"): Lit("d"),
		},
	)

	fmt.Printf("%#v\n", c)


	c = Op("&").Id("Person").Values(
		Dict{
			Id("Age"): Lit(1),
			Id("Name"): Lit("a"),
		},
	)

	fmt.Printf("%#v\n", c)
}