// Anon adds an anonymous import

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("main")
	f.Anon("a")

	f.Func().Id("init").Params().Block()
	//f.Anon("a")

	fmt.Printf("%#v\n", f)
}