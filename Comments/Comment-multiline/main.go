// If the provided comment string contains a newline, the comment is formatted in multiline style

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Comment("a\nb")

	fmt.Printf("%#v\n", c)
}