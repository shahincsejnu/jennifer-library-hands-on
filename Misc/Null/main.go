// Null adds a null item
// Null items render nothing and are not followed by a separator in lists
// In lists, nil will produce the same effect

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Func().Id("foo").Params(nil, Id("s").String(), Null(), Id("i").Int()).Values()

	fmt.Printf("%#v\n", c)
}