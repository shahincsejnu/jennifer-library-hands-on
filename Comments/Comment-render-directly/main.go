// If the comment string starts with "//" or "/*", the automatic formatting is disabled and the string is rendered directly.
// so, in this type of comment do not use "\n" in the comment string cause it will go another line but not as comment

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Id("foo").Call(Comment("/* inline */")).Comment("//no-space")
	//c := Id("foo").Call(Comment("/* inline */")).Comment("//no-space\noka")

	fmt.Printf("%#v\n", c)
}