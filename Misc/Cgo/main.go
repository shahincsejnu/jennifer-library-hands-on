// The cgo "C" pseudo-package is a special case, and always renders without a package alias.
// The import can be added with `Qual`, `Anon` or by supplying a preamble.
// The preamble is added with `File.CgoPreamble` which has the same semantics as Comment.
// If a preamble is provided, the import is separated, and preceded by the preamble.

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("a")
	f.CgoPreamble(`
		#include <stdio.h>
		#include <stdlib.h>
		
		void myprint(char* s) {
			printf("%s\n", s);
		}
	`)

	fmt.Printf("%#v\n", f)

	f.Func().Id("init").Params().Block(
		Id("cs").Op(":=").Qual("C", "CString").Call(Lit("Hello from stdio\n")),
		Qual("C", "myprint").Call(Id("cs")),
		Qual("C", "free").Call(Qual("unsafe", "Pointer").Parens(Id("cs"))),
	)

	fmt.Printf("%#v\n", f)
}