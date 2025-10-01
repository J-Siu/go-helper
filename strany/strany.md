# StrAny

Convert anything to string.

## Installation

```sh
go get github.com/J-Siu/go-helper/v2/strany
```

## Usage

```go
import "github.com/J-Siu/go-helper/v2/strany"
```

## Example

Full example in top level example folder.

```go
func main() {
  var (
    n           = new(NUM).New()
    f32 float32 = 100.000001
    f64 float64 = 100.000001
  )
  fmt.Println(*strany.Any(n))
  fmt.Println(*strany.Any(f32))
  fmt.Println(*strany.Any(f64))
  fmt.Println(*strany.Any(&f32))
  fmt.Println(*strany.Any(&f64))
}
```

## License

The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
