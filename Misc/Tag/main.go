// Tag renders a struct tag

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Type().Id("foo").Struct(
		Id("A").String().Tag(map[string]string{"json": "a"}),
		Id("B").Int().Tag(map[string]string{"json": "b", "bar": "baz"}),
	)

	fmt.Printf("%#v\n", c)
}

// Note: the items are ordered by key when rendered to ensure repeatable code