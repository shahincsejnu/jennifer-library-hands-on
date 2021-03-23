* Several methods render curly braces, summarized below:
 
|   Name    | Prefix      | Separator |               Example               |
|-----------|-------------|-----------|-------------------------------------|
| Block     |             |  `\n`     | `func a() { ... }` or `if a { ... }`|
| Interface | `interface` |  `\n`     | `interface { ... }`                 |
| Struct    | `struct`    |  `\n`     | `struct { ... }`                    |
| Values    |             |  `,`      | `[]int{1, 2}` or `A{B: "c"}`        |