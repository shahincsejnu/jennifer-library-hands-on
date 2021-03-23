* Several methods output parenthesis, summarized below:
 
|   Name    | Prefix    | Separator|               Example               |
|-----------|-----------|----------|-------------------------------------|
| Call      |           |  `,`     | `fmt.Println(b, c)`                 |
| Params    |           |  `,`     | `func (a *A) Foo(i int) { ... } `   |
| Defs      |           |  `\n`    | `const { ... }`                     |
| Parens    |           |          | `[]byte(s)` or `a / (b + c)`        |
| Assert    |   `.`     |          | `s, ok := i.(string)`               |