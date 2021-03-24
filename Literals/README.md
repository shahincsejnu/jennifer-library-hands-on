## Intro
* Lit renders a literal. 
* Lit supports only built-in types (bool, string, int, complex128, float64, float32, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr and complex64). 
* Passing any other type will panic.
* whatever will (valid and except `byte` and `rune`) put inside the `Lit()` it will render as it is



## Concepts

* For the default constant types (bool, int, float64, string, complex128), Lit will render the untyped constant.

    | Code          | Output     |
    |---------------|------------|
    | `Lit(true)`   | `true`     |
    | `Lit(1)`      | `1`        |
    | `Lit(1.0)`    | `1.0`      |
    | `Lit("foo")`  | `"foo"`    |
    | `Lit(0 + 1i)` | `(0 + 1i)` |


* For all other built-in types (float32, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, complex64), Lit will also render the type.

    |       Code                |       Output         |
    |---------------------------|----------------------|
    | `Lit(float32(1))`         | `float32(1)`         |
    | `Lit(int16(1))`           | `int16(1)`           |
    | `Lit(uint8(0x1))`         | `uint8(0x1)`         |
    | `Lit(complex64(0 + 1i))`  | `complex64(0 + 1i)`  |


* The built-in alias types `byte` and `rune` need a special case. `LitRune` and `LitByte` render rune and byte literals.

    |       Code            |    Output     |
    |-----------------------|---------------|
    | `LitRune('x')`        | `x`           |
    | `LitByte(byte(0x1))`  | `byte(0x1)`   |


