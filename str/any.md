# go-StrAny

Convert anything to string.

## Installation

```sh
go get github.com/J-Siu/go-helper/v2/str
```

## Example

```go
package main

import (
  "fmt"

  "github.com/J-Siu/go-helper/v2/str"
)

type NUM struct {
  I     int
  I8    int8
  I16   int16
  I32   int32
  I64   int64
  UI    uint
  UI8   uint8
  UI16  uint16
  UI32  uint32
  UI64  uint64
  F32   float32
  F64   float64
  PI    *int
  PI8   *int8
  PI16  *int16
  PI32  *int32
  PI64  *int64
  PUI   *uint
  PUI8  *uint8
  PUI16 *uint16
  PUI32 *uint32
  PUI64 *uint64
  PF32  *float32
  PF64  *float64
}

func (N *NUM) New() *NUM {
  N.I = 255
  N.I8 = 127
  N.I16 = 255
  N.I32 = 255
  N.I64 = 255
  N.UI = 255
  N.UI8 = 255
  N.UI16 = 255
  N.UI32 = 255
  N.UI64 = 255
  N.F32 = 100.00000002
  N.F64 = 100.00000001
  N.PI = &N.I
  N.PI8 = &N.I8
  N.PI16 = &N.I16
  N.PI32 = &N.I32
  N.PI64 = &N.I64
  N.PUI = &N.UI
  N.PUI8 = &N.UI8
  N.PUI16 = &N.UI16
  N.PUI32 = &N.UI32
  N.PUI64 = &N.UI64
  N.PF32 = &N.F32
  N.PF64 = &N.F64
  return N
}

func main() {

  var (
    n           = new(NUM).New()
    f32 float32 = 100.000001
    f64 float64 = 100.000001
    any         = new(str.Any).New()
  )

  fmt.Println(any.Str(n))
  fmt.Println(any.Str(f32))
  fmt.Println(any.Str(f64))
  fmt.Println(any.Str(&f32))
  fmt.Println(any.Str(&f64))
}
```

## License

The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
