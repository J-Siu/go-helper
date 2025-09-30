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

	"github.com/J-Siu/go-helper/v2/str"
)

type StrAny struct {
	err          error
	indent       string // `indent` of json.MarshalIndent(v any, prefix, indent string)
	indentEnable bool   // If `true`, true, use `json.MarshalIndent` for struct, else `json.Marshal`
	prefix       string // `prefix` of json.MarshalIndent(v any, prefix, indent string)
}

// Initialize
func (s *StrAny) New() *StrAny {
	s.err = nil
	s.indent = "  "
	s.indentEnable = true
	s.prefix = ""
	return s
}

// `enable` = `true“, use `json.MarshalIndent` for struct, else `json.Marshal`
func (s *StrAny) IndentEnable(enable bool) *StrAny {
	s.indentEnable = enable
	return s
}

// `indent` of json.MarshalIndent(v any, prefix, indent string)
func (s *StrAny) Indent(indent string) *StrAny {
	s.indent = indent
	return s
}

// `prefix` of json.MarshalIndent(v any, prefix, indent string)
func (s *StrAny) Prefix(prefix string) *StrAny {
	s.prefix = prefix
	return s
}

// Return json.Marshal* error
func (s *StrAny) Err() error { return s.err }

func (s *StrAny) processStr(sP *string) *string {
	if s.indentEnable {
		return str.JsonIndent(sP)
	}
	return sP
}

func (s *StrAny) processStrArray(saP *[]string) *string {
	var out string
	if saP != nil {
		last := len(*saP) - 1
		for index, item := range *saP {
			if s.indentEnable {
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

func (s *StrAny) processByteArray(baP *[]byte) *string {
	var out string
	if baP != nil && len(*baP) > 0 {
		if s.indentEnable {
			return str.ByteJsonIndent(baP)
		} else {
			out = string(*baP)
		}
	}
	return &out
}

func (s *StrAny) processErrArray(eaP *[]error) *string {
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
func (s *StrAny) String(data any) *string {
	var strOut string
	switch v := data.(type) {
	case string:
		return s.processStr(&v)
	case *string:
		return s.processStr(v)
	case []string:
		return s.processStrArray(&v)
	case *[]string:
		return s.processStrArray(v)
	case []byte:
		return s.processByteArray(&v)
	case *[]byte:
		return s.processByteArray(v)
	case bytes.Buffer:
		var b = v.Bytes()
		return s.processByteArray(&b)
	case *bytes.Buffer:
		if v != nil {
			var b = v.Bytes()
			return s.processByteArray(&b)
		}
	case error:
		strOut = v.Error()
	case *error:
		strOut = (*v).Error()
	case []error:
		return s.processErrArray(&v)
	case *[]error:
		return s.processErrArray(v)
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
		if s.indentEnable {
			b, s.err = json.MarshalIndent(v, s.prefix, s.indent)
		} else {
			b, s.err = json.Marshal(v)
		}
		if s.err == nil {
			strOut = string(b)
		}
	}
	return &strOut
}

func (s *StrAny) Str(data any) *string { return s.String(data) }

var strAny = New()

func New() *StrAny { return new(StrAny).New() }

func Any(data any) *string { return strAny.Str(data) }
