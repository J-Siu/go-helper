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
	"strconv"
	"strings"

	"github.com/J-Siu/go-helper/v2/str"
)

type StrAny struct {
	err          error
	indent       string // `indent` of json.MarshalIndent(v any, prefix, indent string)
	indentEnable bool   // If `true`, true, use `json.MarshalIndent` for struct, else `json.Marshal`
	indentPrefix string // `prefix` of json.MarshalIndent(v any, prefix, indent string)
	unquote      bool
}

// Initialize
func (t *StrAny) New() *StrAny {
	t.err = nil
	t.indent = "  "
	t.indentEnable = true
	t.indentPrefix = ""
	t.unquote = true
	return t
}

// Return json.Marshal* error
func (t *StrAny) Err() error { return t.err }

// `indent` of json.MarshalIndent(v any, prefix, indent string)
func (t *StrAny) Indent(indent string) *StrAny {
	t.indent = indent
	return t
}

// `enable` = `true“, use `json.MarshalIndent` for struct, else `json.Marshal`
func (t *StrAny) IndentEnable(enable bool) *StrAny {
	t.indentEnable = enable
	return t
}

// `prefix` of json.MarshalIndent(v any, prefix, indent string)
func (t *StrAny) IndentPrefix(prefix string) *StrAny {
	t.indentPrefix = prefix
	return t
}

func (t *StrAny) UnquoteEnable(enable bool) *StrAny {
	t.unquote = enable
	return t
}

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
func (t *StrAny) processStr(sP *string) *string {
	out := ""
	if sP != nil {
		if t.indentEnable {
			out = *str.JsonIndent(sP)
		}
	}
	return &out
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
	var out string
	if eaP != nil {
		last := len(*eaP) - 1
		for i, e := range *eaP {
			out += e.Error()
			if i < last {
				out += "\n"
			}
		}
	}
	return &out
}

// Output `data` as string
//
// If `IndentEnable` is true, struct will be converted with `json.MarshalIndent`, else `json.Marshal`
func (t *StrAny) String(data any) *string {
	var strOut string
	switch v := data.(type) {
	case string:
		strOut = *t.processStr(&v)
	case *string:
		strOut = *t.processStr(v)
	case []string:
		strOut = *t.processStrArray(&v)
	case *[]string:
		strOut = *t.processStrArray(v)
	case []byte:
		strOut = *t.processByteArray(&v)
	case *[]byte:
		strOut = *t.processByteArray(v)
	case bytes.Buffer:
		var b = v.Bytes()
		strOut = *t.processByteArray(&b)
	case *bytes.Buffer:
		if v != nil {
			var b = v.Bytes()
			strOut = *t.processByteArray(&b)
		}
	case error:
		strOut = v.Error()
	case *error:
		strOut = (*v).Error()
	case []error:
		strOut = *t.processErrArray(&v)
	case *[]error:
		strOut = *t.processErrArray(v)
	case int:
		strOut = fmt.Sprint(v)
	case int8:
		strOut = fmt.Sprint(v)
	case int16:
		strOut = fmt.Sprint(v)
	case int32:
		strOut = fmt.Sprint(v)
	case int64:
		strOut = fmt.Sprint(v)
	case uint:
		strOut = fmt.Sprint(v)
	case uint8:
		strOut = fmt.Sprint(v)
	case uint16:
		strOut = fmt.Sprint(v)
	case uint32:
		strOut = fmt.Sprint(v)
	case uint64:
		strOut = fmt.Sprint(v)
	case float32:
		strOut = fmt.Sprint(v)
	case float64:
		strOut = fmt.Sprint(v)
	case *int:
		if v != nil {
			strOut = fmt.Sprint(*v)
		}
	case *int8:
		if v != nil {
			strOut = fmt.Sprint(*v)
		}
	case *int16:
		if v != nil {
			strOut = fmt.Sprint(*v)
		}
	case *int32:
		if v != nil {
			strOut = fmt.Sprint(*v)
		}
	case *int64:
		if v != nil {
			strOut = fmt.Sprint(*v)
		}
	case *uint:
		if v != nil {
			strOut = fmt.Sprint(*v)
		}
	case *uint8:
		if v != nil {
			strOut = fmt.Sprint(*v)
		}
	case *uint16:
		if v != nil {
			strOut = fmt.Sprint(*v)
		}
	case *uint32:
		if v != nil {
			strOut = fmt.Sprint(*v)
		}
	case *uint64:
		if v != nil {
			strOut = fmt.Sprint(*v)
		}
	case *float32:
		if v != nil {
			strOut = fmt.Sprint(*v)
		}
	case *float64:
		if v != nil {
			strOut = fmt.Sprint(*v)
		}
	default:
		var b []byte
		if t.indentEnable {
			b, t.err = json.MarshalIndent(v, t.indentPrefix, t.indent)
		} else {
			b, t.err = json.Marshal(v)
		}
		if t.err == nil {
			strOut = string(b)
		}
	}
	if t.unquote {
		strOut = *t.processUnquote(&strOut)
	}
	return &strOut
}

// Matching package level Any()
func (t *StrAny) Any(data any) *string { return t.String(data) }

var strAny = New()

func New() *StrAny { return new(StrAny).New() }

// Convert any to *string
func Any(data any) *string { return strAny.String(data) }
