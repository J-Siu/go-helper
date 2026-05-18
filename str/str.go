/*
The MIT License

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

// collection of string related helper functions
package str

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"strings"
	"unsafe"

	"github.com/charlievieth/strcase"
)

// Check if string array contain a string.
//
// Return false if either arrIn is nil
func ArrayContains(arrP *[]string, str string, caseSensitive bool) bool {
	if arrP != nil {
		for _, s := range *arrP {
			if caseSensitive && s == str || strcase.Compare(s, str) == 0 {
				return true
			}
		}
	}
	return false
}

// Check if string array contains a substring
//
// Return false if arrIn is nil
func ArrayContainsSubString(arrP *[]string, subStr string, caseSensitive bool) bool {
	if arrP != nil {
		for _, s := range *arrP {
			if caseSensitive && strings.Contains(subStr, s) || strcase.Contains(s, subStr) {
				return true
			}
		}
	}
	return false
}

// Return a new []string with empty lines removed from *[]string.
//   - Original []string not modified.
func ArrayPtrRemoveEmpty(arrP *[]string) (out []string) {
	if arrP != nil {
		for _, s := range *arrP {
			if s != "" {
				out = append(out, s)
			}
		}
	}
	return out
}

// *[]string to string, each element followed by "\n"
func ArraySPrintln(arrP *[]string) (out string) {
	if arrP != nil {
		return strings.Join(*arrP, "\n")
	}
	return out
}

// Check if string contains any substring of an array
//   - result (bool)
//   - if result is true, `resultVal` == matching substring
//   - if result is false, `resultVal` == ""
func ContainsAnySubStrings(str string, subStrArrP *[]string, caseSensitive bool) (result bool, resultVal string) {
	// prefix := "matchList"
	if subStrArrP != nil {
		for _, subStr := range *subStrArrP {
			if caseSensitive && strcase.Contains(str, subStr) || strcase.Contains(str, subStr) {
				result = true
				resultVal = subStr
				break
			}
		}
	}
	return result, resultVal
}

// Return only bool from [ContainsAnySubStrings]
func ContainsAnySubStringsBool(str string, subStrArrP *[]string, caseSensitive bool) (result bool) {
	result, _ = ContainsAnySubStrings(str, subStrArrP, caseSensitive)
	return result
}

// *string to *[]string, split by "\n"
func LnSplit(strP *string) *[]string {
	var strOut []string
	if strP != nil {
		strOut = strings.Split(*strP, "\n")
	}
	return &strOut
}

// Return *string of "" if failed
func JsonIndent(strP *string) *string {
	out := ""
	if strP != nil {
		var byteA = []byte(*strP)
		return ByteJsonIndent(&byteA)
	}
	return &out
}

// Return *string of "" if json.Marshal failed
func JsonMarshal(strP *string) *string {
	out := ""
	if strP != nil {
		if p, e := json.Marshal(strP); e == nil {
			out = string(p)
		}
	}
	return &out
}

// Return *string of "" if json.Indent failed
func ByteJsonIndent(baP *[]byte) *string {
	out := ""
	if baP != nil {
		var dst bytes.Buffer
		if err := json.Indent(&dst, *baP, "", "  "); err == nil {
			out = strings.Trim(dst.String(), "\n")
		} else {
			return (*string)(unsafe.Pointer(baP))
		}
	}
	return &out
}

func ByteHex(b []byte) string { return hex.EncodeToString(b) }

// --- bool

// Return "OK"/"Fail"
func Ok(b bool) string {
	if b {
		return "OK"
	}
	return "Fail"
}

// Return "Success"/"Fail"
func Success(b bool) string {
	if b {
		return "Success"
	}
	return "Fail"
}

// Return "Yes"/"No"
func YesNo(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}
