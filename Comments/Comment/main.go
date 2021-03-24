// Comment adds a comment
// If the provided string contains a newline, the comment is formatted in multiple style

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("main")
	f.Comment("demo comment")
	f.Func().Id("Foo").Params().String().Block(
		Return(Lit("foo")).Comment("return the string foo"),
	)

	fmt.Printf("%#v\n", f)

	f.Save("generated-by-Comment.go")
}