/*
Copyright © 2026 John, Sing Dao, Siu <john.sd.siu@gmail.com>

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

package strany

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/J-Siu/go-helper/v2/array"
	"github.com/J-Siu/go-helper/v2/basestruct"
	"github.com/J-Siu/go-helper/v2/str"
)

// var strEmpty = ""

type StrAny struct {
	*basestruct.Base
	debug        bool
	indent       string // `indent` of json.MarshalIndent(v any, prefix, indent string)
	indentEnable bool   // If `true`, true, use `json.MarshalIndent` for struct, else `json.Marshal`
	indentPrefix string // `prefix` of json.MarshalIndent(v any, prefix, indent string)
	unquote      bool
}

// Initialize
func (t *StrAny) New() *StrAny {
	t.Base = new(basestruct.Base)
	t.MyType = "StrAny"

	t.indent = "  "
	t.indentEnable = true
	t.indentPrefix = ""
	t.unquote = true

	t.Initialized = true
	return t
}

// Output `data` as string
//
// If `IndentEnable` is true, struct will be converted with `json.MarshalIndent`, else `json.Marshal`
func (t *StrAny) Any(data any) (out string) {
	prefix := t.MyType + ".Any"
	var (
		bBuffer  bytes.Buffer
		sBuilder strings.Builder
	)
	if t.debug {
		fmt.Printf("%s %T\n", prefix, data)
		fmt.Printf("%s unquote: %t\n", prefix, t.unquote)
	}
	switch v := data.(type) {
	case string:
		str.JsonIndent(&v, &bBuffer)
		out = bBuffer.String()
	case *string:
		str.JsonIndent(v, &bBuffer)
		out = bBuffer.String()
	case []string:
		out = t.processStrArray(&v, &sBuilder).String()
	case *[]string:
		out = t.processStrArray(v, &sBuilder).String()
	case []byte:
		str.ByteJsonIndent(&v, &bBuffer)
		out = bBuffer.String()
	case *[]byte:
		str.ByteJsonIndent(v, &bBuffer)
		out = bBuffer.String()
	case bytes.Buffer:
		var b = v.Bytes()
		str.ByteJsonIndent(&b, &bBuffer)
		out = bBuffer.String()
	case *bytes.Buffer:
		if v != nil {
			var b = v.Bytes()
			str.ByteJsonIndent(&b, &bBuffer)
			out = bBuffer.String()
		}
	case error:
		out = v.Error()
	case *error:
		out = (*v).Error()
	case []error:
		out = t.processErrArray(&v, &sBuilder).String()
	case *[]error:
		out = t.processErrArray(v, &sBuilder).String()
	case array.Array[error]:
		out = t.processGenericError(&v, &sBuilder).String()
	case *array.Array[error]:
		out = t.processGenericError(v, &sBuilder).String()
	case array.Array[any]:
		out = t.processGenericArray(&v, &sBuilder).String()
	case *array.Array[any]:
		out = t.processGenericArray(v, &sBuilder).String()
	case int:
		out = fmt.Sprint(v)
	case int8:
		out = fmt.Sprint(v)
	case int16:
		out = fmt.Sprint(v)
	case int32:
		out = fmt.Sprint(v)
	case int64:
		out = fmt.Sprint(v)
	case uint:
		out = fmt.Sprint(v)
	case uint8:
		out = fmt.Sprint(v)
	case uint16:
		out = fmt.Sprint(v)
	case uint32:
		out = fmt.Sprint(v)
	case uint64:
		out = fmt.Sprint(v)
	case float32:
		out = fmt.Sprint(v)
	case float64:
		out = fmt.Sprint(v)
	case *int:
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *int8:
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *int16:
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *int32:
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *int64:
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *uint:
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *uint8:
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *uint16:
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *uint32:
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *uint64:
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *float32:
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *float64:
		if v != nil {
			out = fmt.Sprint(*v)
		}
	default:
		if t.debug {
			fmt.Println(prefix, "default")
			fmt.Println(prefix, "indentEnable:", t.indentEnable)
		}
		var b []byte
		if b, t.Err = json.Marshal(v); t.Err == nil {
			if t.indentEnable {
				str.ByteJsonIndent(&b, &bBuffer)
				out = bBuffer.String()
			} else {
				out = string(b)
			}
		}
	}
	if t.unquote {
		// From Coconut: https://stackoverflow.com/a/51579784/1810391
		if dst, e := strconv.Unquote(strings.ReplaceAll(strconv.Quote(out), `\\u`, `\u`)); e == nil {
			out = dst
		}
	}
	return out
}

func (t *StrAny) DebugEnable(enable bool) *StrAny {
	t.debug = enable
	return t
}

// `indent` of json.MarshalIndent(v any, prefix, indent string)
func (t *StrAny) Indent(indent string) *StrAny {
	t.indent = indent
	return t
}
func (t *StrAny) GetIndent() string { return t.indent }

// `enable` = `true“, use `json.MarshalIndent` for struct, else `json.Marshal`
func (t *StrAny) IndentEnable(enable bool) *StrAny {
	t.indentEnable = enable
	return t
}
func (t *StrAny) GetIndentEnable() bool { return t.indentEnable }

// `prefix` of json.MarshalIndent(v any, prefix, indent string)
func (t *StrAny) IndentPrefix(prefix string) *StrAny {
	t.indentPrefix = prefix
	return t
}
func (t *StrAny) GetIndentPrefix() string { return t.indentPrefix }

func (t *StrAny) UnquoteEnable(enable bool) *StrAny {
	t.unquote = enable
	return t
}
func (t *StrAny) GetUnquoteEnable() bool { return t.unquote }

func (t *StrAny) processStrArray(src *[]string, dst *strings.Builder) *strings.Builder {
	var (
		bBuffer bytes.Buffer
	)
	if src != nil {
		last := len(*src) - 1
		for index, item := range *src {
			if t.indentEnable {
				bBuffer.Reset()
				dst.WriteString(str.JsonIndent(&item, &bBuffer).String())
			} else {
				dst.WriteString(item)
			}
			if index < last {
				dst.WriteString("\n")
			}
		}
	}
	return dst
}

func (t *StrAny) processErrArray(src *[]error, dst *strings.Builder) *strings.Builder {
	if src != nil {
		last := len(*src) - 1
		for index, item := range *src {
			dst.WriteString(item.Error())
			if index < last {
				dst.WriteString("\n")
			}
		}
	}
	return dst
}

func (t *StrAny) processGenericArray(src *array.Array[any], dst *strings.Builder) *strings.Builder {
	if src != nil {
		last := len(*src) - 1
		for index, item := range *src {
			dst.WriteString(t.Any(item))
			if index < last {
				dst.WriteString("\n")
			}
		}
	}
	return dst
}

func (t *StrAny) processGenericError(src *array.Array[error], dst *strings.Builder) *strings.Builder {
	if src != nil {
		last := len(*src) - 1
		for index, item := range *src {
			dst.WriteString(item.Error())
			if index < last {
				dst.WriteString("\n")
			}
		}
	}
	return dst
}
