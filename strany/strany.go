/*
Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Convert anything to string
package strany

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type strAny struct {
	err          error
	indentEnable bool   // If `true`, true, use `json.MarshalIndent` for struct, else `json.Marshal`
	indent       string // `indent` of json.MarshalIndent(v any, prefix, indent string)
	prefix       string // `prefix` of json.MarshalIndent(v any, prefix, indent string)
}

// Initialize
func (e *strAny) New() *strAny {
	e.err = nil
	e.indent = "  "
	e.prefix = ""
	return e
}

// `enable` = `true“, use `json.MarshalIndent` for struct, else `json.Marshal`
func (e *strAny) IndentEnable(enable bool) *strAny {
	e.indentEnable = enable
	return e
}

// `indent` of json.MarshalIndent(v any, prefix, indent string)
func (e *strAny) Indent(indent string) *strAny {
	e.indent = indent
	return e
}

// `prefix` of json.MarshalIndent(v any, prefix, indent string)
func (e *strAny) Prefix(prefix string) *strAny {
	e.prefix = prefix
	return e
}

// Return json.Marshal* error
func (e *strAny) Err() error { return e.err }

// Output `data` as string
//
// If `IndentEnable` is true, struct will be converted with `json.MarshalIndent`, else `json.Marshal`
func (e *strAny) Str(data any) (str string) {
	switch v := data.(type) {
	case string:
		str = v
	case *string:
		str = *v
	case error:
		str = v.Error()
	case *error:
		str = (*v).Error()
	case bytes.Buffer:
		str = v.String()
	case *bytes.Buffer:
		if v != nil {
			str = v.String()
		}
	case int:
		str = fmt.Sprint(v)
	case int8:
		str = fmt.Sprint(v)
	case int16:
		str = fmt.Sprint(v)
	case int32:
		str = fmt.Sprint(v)
	case int64:
		str = fmt.Sprint(v)
	case uint:
		str = fmt.Sprint(v)
	case uint8:
		str = fmt.Sprint(v)
	case uint16:
		str = fmt.Sprint(v)
	case uint32:
		str = fmt.Sprint(v)
	case uint64:
		str = fmt.Sprint(v)
	case float32:
		str = fmt.Sprint(v)
	case float64:
		str = fmt.Sprint(v)
	case *int:
		if v != nil {
			str = fmt.Sprint(*v)
		}
	case *int8:
		if v != nil {
			str = fmt.Sprint(*v)
		}
	case *int16:
		if v != nil {
			str = fmt.Sprint(*v)
		}
	case *int32:
		if v != nil {
			str = fmt.Sprint(*v)
		}
	case *int64:
		if v != nil {
			str = fmt.Sprint(*v)
		}
	case *uint:
		if v != nil {
			str = fmt.Sprint(*v)
		}
	case *uint8:
		if v != nil {
			str = fmt.Sprint(*v)
		}
	case *uint16:
		if v != nil {
			str = fmt.Sprint(*v)
		}
	case *uint32:
		if v != nil {
			str = fmt.Sprint(*v)
		}
	case *uint64:
		if v != nil {
			str = fmt.Sprint(*v)
		}
	case *float32:
		if v != nil {
			str = fmt.Sprint(*v)
		}
	case *float64:
		if v != nil {
			str = fmt.Sprint(*v)
		}
	default:
		var b []byte
		if e.indentEnable {
			b, e.err = json.MarshalIndent(v, e.prefix, e.indent)
		} else {
			b, e.err = json.Marshal(v)
		}
		if e.err == nil {
			str = string(b)
		}
	}
	return str
}

// ---

var fromAny = New()

// Initialize
func New() *strAny { return new(strAny).New() }

// `enable` = `true“, use `json.MarshalIndent` for struct, else `json.Marshal`
func IndentEnable(enable bool) *strAny { return fromAny.IndentEnable(enable) }

// `indent` of json.MarshalIndent(v any, prefix, indent string)
func Indent(indent string) *strAny { return fromAny.Indent(indent) }

// `prefix` of json.MarshalIndent(v any, prefix, indent string)
func Prefix(prefix string) *strAny { return fromAny.Prefix(prefix) }

// Return json.Marshal* error
func Err() error { return fromAny.Err() }

// Output `data` as string
//
// If `IndentEnable` is true, struct will be converted with `json.MarshalIndent`, else `json.Marshal`
func Str(data any) string { return fromAny.Str(data) }
