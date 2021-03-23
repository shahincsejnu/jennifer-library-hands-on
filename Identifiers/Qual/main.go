// Qual renders a qualified identifier

package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func main() {
	c := Qual("encoding/gob", "NewEncoder").Call()

	fmt.Printf("%#v\n", c)

	// Imports are automatically added when used with a File.
	// If the path matches local path, the package name is omitted
	// If package names conflict they are automatically renamed

	f := NewFilePath("a.b/c")

	f.Func().Id("init").Call().Block(
		Qual("a.b/c", "Foo").Call().Comment("Local package - name is omitted."),
		Qual("d.e/f", "Bar").Call().Comment("Import is automatically added."),
		Qual("g.h/f", "Baz").Call().Comment("Colliding package name is renamed"),
	)

	fmt.Printf("%#v", f)

	//f.Save("generated-by-Qual.go")
}