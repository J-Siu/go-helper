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

## Types and Functions

### Structure

```go
type StrAny struct {
  *basestruct.Base
  indent       string // `indent` of json.MarshalIndent(v any, prefix, indent string)
  indentEnable bool   // If `true`, true, use `json.MarshalIndent` for struct, else `json.Marshal`
  indentPrefix string // `prefix` of json.MarshalIndent(v any, prefix, indent string)
  unquote      bool
  debug        bool
}

func (t *StrAny) New() *StrAny
func (t *StrAny) Any(data any) *string
func (t *StrAny) String(data any) *string
func (t *StrAny) DebugEnable(enable bool) *StrAny
func (t *StrAny) Indent(indent string) *StrAny
func (t *StrAny) IndentEnable(enable bool) *StrAny
func (t *StrAny) IndentPrefix(prefix string) *StrAny
func (t *StrAny) UnquoteEnable(enable bool) *StrAny
```

### Package Functions

```go
func New() *StrAny
func Any(data any) *string
func String(data any) *string
func DebugEnable(enable bool) *StrAny
func ToPtr[T any](v T) *T
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
