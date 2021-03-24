// Add appends the provided items to the statement

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	ptr := Op("*")

	//c := Id("a").Op("=").Add(Op("*")).Id("b")
	c := Id("a").Op("=").Add(ptr).Id("b")

	fmt.Printf("%#v\n", c)

	a := Id("a")
	i := Int()
	c = Var().Add(a, i)

	fmt.Printf("%#v\n", c)
}