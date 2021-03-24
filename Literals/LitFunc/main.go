// LitFunc generates the value to render by executing the provided function

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Id("a").Op(":=").LitFunc(
		func() interface{} {
			return 1 + 1
		},
	)

	fmt.Printf("%#v\n", c)
}