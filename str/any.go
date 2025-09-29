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
package str

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Any struct {
	err          error
	indentEnable bool   // If `true`, true, use `json.MarshalIndent` for struct, else `json.Marshal`
	indent       string // `indent` of json.MarshalIndent(v any, prefix, indent string)
	prefix       string // `prefix` of json.MarshalIndent(v any, prefix, indent string)
}

// Initialize
func (s *Any) New() *Any {
	s.err = nil
	s.indent = "  "
	s.indentEnable = true
	s.prefix = ""
	return s
}

// `enable` = `true“, use `json.MarshalIndent` for struct, else `json.Marshal`
func (s *Any) IndentEnable(enable bool) *Any {
	s.indentEnable = enable
	return s
}

// `indent` of json.MarshalIndent(v any, prefix, indent string)
func (s *Any) Indent(indent string) *Any {
	s.indent = indent
	return s
}

// `prefix` of json.MarshalIndent(v any, prefix, indent string)
func (s *Any) Prefix(prefix string) *Any {
	s.prefix = prefix
	return s
}

// Return json.Marshal* error
func (s *Any) Err() error { return s.err }

func (s *Any) processStr(sP *string) *string {
	if s.indentEnable {
		return JsonIndent(sP)
	}
	return sP
}

func (s *Any) processStrArray(saP *[]string) *string {
	var out string
	if saP != nil {
		last := len(*saP) - 1
		for index, item := range *saP {
			if s.indentEnable {
				out += *JsonIndent(&item)
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

func (s *Any) processByteArray(baP *[]byte) *string {
	var out string
	if baP != nil && len(*baP) > 0 {
		if s.indentEnable {
			return ByteJsonIndent(baP)
		} else {
			out = string(*baP)
		}
	}
	return &out
}

func (s *Any) processErrArray(eaP *[]error) *string {
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
func (s *Any) Str(data any) *string {
	var out string
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
		out = v.Error()
	case *error:
		out = (*v).Error()
	case []error:
		return s.processErrArray(&v)
	case *[]error:
		return s.processErrArray(v)
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
		var b []byte
		if s.indentEnable {
			b, s.err = json.MarshalIndent(v, s.prefix, s.indent)
		} else {
			b, s.err = json.Marshal(v)
		}
		if s.err == nil {
			out = string(b)
		}
	}
	return &out
}
