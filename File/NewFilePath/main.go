// NewFilePath creates a new file while specifying the package path
// - the package name is inferred from the path.

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFilePath("a.b/c")

	f.Func().Id("main").Params().Block(
		Qual("a.b/d", "Foo").Call(),
	)

	fmt.Printf("%#v\n", f)
}