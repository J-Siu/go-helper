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

type StrAny struct {
	*basestruct.Base
	indent       string // `indent` of json.MarshalIndent(v any, prefix, indent string)
	indentEnable bool   // If `true`, true, use `json.MarshalIndent` for struct, else `json.Marshal`
	indentPrefix string // `prefix` of json.MarshalIndent(v any, prefix, indent string)
	unquote      bool
	debug        bool
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
func (t *StrAny) Any(data any) *string {
	prefix := t.MyType + ".Any"
	out := ""
	switch v := data.(type) {
	case string:
		if t.debug {
			fmt.Println(prefix, "string")
		}
		out = *t.processStrIndent(&v)
	case *string:
		if t.debug {
			fmt.Println(prefix, "*string")
		}
		out = *t.processStrIndent(v)
	case []string:
		if t.debug {
			fmt.Println(prefix, "[]string")
		}
		out = *t.processStrArray(&v)
	case *[]string:
		if t.debug {
			fmt.Println(prefix, "*[]string")
		}
		out = *t.processStrArray(v)
	case []byte:
		if t.debug {
			fmt.Println(prefix, "[]byte")
		}
		out = *t.processByteArray(&v)
	case *[]byte:
		if t.debug {
			fmt.Println(prefix, "*[]byte")
		}
		out = *t.processByteArray(v)
	case bytes.Buffer:
		if t.debug {
			fmt.Println(prefix, "bytes.Buffer")
		}
		var b = v.Bytes()
		out = *t.processByteArray(&b)
	case *bytes.Buffer:
		if t.debug {
			fmt.Println(prefix, "*bytes.Buffer")
		}
		if v != nil {
			var b = v.Bytes()
			out = *t.processByteArray(&b)
		}
	case error:
		if t.debug {
			fmt.Println(prefix, "error")
		}
		out = v.Error()
	case *error:
		if t.debug {
			fmt.Println(prefix, "*error")
		}
		out = (*v).Error()
	case []error:
		if t.debug {
			fmt.Println(prefix, "[]error")
		}
		out = *t.processErrArray(&v)
	case *[]error:
		if t.debug {
			fmt.Println(prefix, "*[]error")
		}
		out = *t.processErrArray(v)
	case array.Array[error]:
		if t.debug {
			fmt.Println(prefix, "array.Array[error]")
		}
		out = *t.processGenericError(&v)
	case *array.Array[error]:
		if t.debug {
			fmt.Println(prefix, "*array.Array[error]")
		}
		out = *t.processGenericError(v)
	case array.Array[any]:
		if t.debug {
			fmt.Println(prefix, "array.Array[any]")
		}
		out = *t.processGenericArray(&v)
	case *array.Array[any]:
		if t.debug {
			fmt.Println(prefix, "*array.Array[any]")
		}
		out = *t.processGenericArray(v)
	case int:
		if t.debug {
			fmt.Println(prefix, "int")
		}
		out = fmt.Sprint(v)
	case int8:
		if t.debug {
			fmt.Println(prefix, "int8")
		}
		out = fmt.Sprint(v)
	case int16:
		if t.debug {
			fmt.Println(prefix, "int16")
		}
		out = fmt.Sprint(v)
	case int32:
		if t.debug {
			fmt.Println(prefix, "int32")
		}
		out = fmt.Sprint(v)
	case int64:
		if t.debug {
			fmt.Println(prefix, "int64")
		}
		out = fmt.Sprint(v)
	case uint:
		if t.debug {
			fmt.Println(prefix, "uint")
		}
		out = fmt.Sprint(v)
	case uint8:
		if t.debug {
			fmt.Println(prefix, "uint8")
		}
		out = fmt.Sprint(v)
	case uint16:
		if t.debug {
			fmt.Println(prefix, "uint16")
		}
		out = fmt.Sprint(v)
	case uint32:
		if t.debug {
			fmt.Println(prefix, "uint32")
		}
		out = fmt.Sprint(v)
	case uint64:
		if t.debug {
			fmt.Println(prefix, "uint64")
		}
		out = fmt.Sprint(v)
	case float32:
		if t.debug {
			fmt.Println(prefix, "float32")
		}
		out = fmt.Sprint(v)
	case float64:
		if t.debug {
			fmt.Println(prefix, "float64")
		}
		out = fmt.Sprint(v)
	case *int:
		if t.debug {
			fmt.Println(prefix, "*int")
		}
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *int8:
		if t.debug {
			fmt.Println(prefix, "*int8")
		}
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *int16:
		if t.debug {
			fmt.Println(prefix, "*int16")
		}
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *int32:
		if t.debug {
			fmt.Println(prefix, "*int32")
		}
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *int64:
		if t.debug {
			fmt.Println(prefix, "*int64")
		}
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *uint:
		if t.debug {
			fmt.Println(prefix, "*uint")
		}
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *uint8:
		if t.debug {
			fmt.Println(prefix, "*uint8")
		}
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *uint16:
		if t.debug {
			fmt.Println(prefix, "*uint16")
		}
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *uint32:
		if t.debug {
			fmt.Println(prefix, "*uint32")
		}
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *uint64:
		if t.debug {
			fmt.Println(prefix, "*uint64")
		}
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *float32:
		if t.debug {
			fmt.Println(prefix, "*float32")
		}
		if v != nil {
			out = fmt.Sprint(*v)
		}
	case *float64:
		if t.debug {
			fmt.Println(prefix, "*float64")
		}
		if v != nil {
			out = fmt.Sprint(*v)
		}
	default:
		if t.debug {
			fmt.Println(prefix, "default")
		}
		var b []byte
		if t.indentEnable {
			b, t.Err = json.MarshalIndent(v, t.indentPrefix, t.indent)
		} else {
			b, t.Err = json.Marshal(v)
		}
		if t.Err == nil {
			out = string(b)
		}
	}
	if t.unquote {
		if t.debug {
			fmt.Println(prefix, "unquote")
		}
		out = *t.processUnquote(&out)
	}
	return &out
}

// Convert any to *string. Alias of Any()
func (t *StrAny) String(data any) *string { return t.Any(data) }

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

// Unescape utf8 string
func (t *StrAny) processUnquote(sP *string) *string {
	out := ""
	if sP != nil {
		var e error
		// From Coconut: https://stackoverflow.com/a/51579784/1810391
		out, e = strconv.Unquote(strings.Replace(strconv.Quote(string(*sP)), `\\u`, `\u`, -1))
		if e != nil {
			// if unquote failed, return original string
			return sP
		}
	}
	return &out
}

// Indent processing
func (t *StrAny) processStrIndent(sP *string) *string {
	var out *string
	if sP != nil {
		if t.indentEnable {
			out = str.JsonIndent(sP)
		} else {
			out = sP
		}
	}
	return out
}

func (t *StrAny) processStrArray(saP *[]string) *string {
	out := ""
	if saP != nil {
		last := len(*saP) - 1
		for index, item := range *saP {
			if t.indentEnable {
				out += *str.JsonIndent(&item)
			} else {
				out += item
			}
			if index < last {
				out += "\n"
			}
		}
	}
	return &out
}

func (t *StrAny) processByteArray(baP *[]byte) *string {
	out := ""
	if baP != nil && len(*baP) > 0 {
		if t.indentEnable {
			out = *str.ByteJsonIndent(baP)
		} else {
			out = string(*baP)
		}
	}
	return &out
}

func (t *StrAny) processErrArray(eaP *[]error) *string {
	out := ""
	if eaP != nil {
		last := len(*eaP) - 1
		for index, item := range *eaP {
			out += item.Error()
			if index < last {
				out += "\n"
			}
		}
	}
	return &out
}

func (t *StrAny) processGenericArray(eaP *array.Array[any]) *string {
	out := ""
	if eaP != nil {
		last := len(*eaP) - 1
		for index, item := range *eaP {
			out += *t.String(item)
			if index < last {
				out += "\n"
			}
		}
	}
	return &out
}

func (t *StrAny) processGenericError(eaP *array.Array[error]) *string {
	out := ""
	if eaP != nil {
		last := len(*eaP) - 1
		for index, item := range *eaP {
			out += *t.String(item)
			if index < last {
				out += "\n"
			}
		}
	}
	return &out
}
