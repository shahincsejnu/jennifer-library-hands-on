// NewFilePathName creates a new file with the specified package path and name.

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFilePathName("a.b/c", "main")
	f.Func().Id("main").Params().Block(
		//Qual("a.b/d", "Foo").Call(),
		Qual("a.b/c", "Foo").Call(),
	)

	fmt.Printf("%#v\n", f)
}