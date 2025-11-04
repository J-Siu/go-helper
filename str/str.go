/*
The MIT License

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

// collection of string related helper functions
package str

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/charlievieth/strcase"
)

// Check if string array contain a string.
//
// Return false if either arrIn or strIn is nil
func ArrayContains(arrIn *[]string, strIn *string, caseSensitive bool) bool {
	if arrIn != nil && strIn != nil {
		for _, s := range *arrIn {
			if caseSensitive && s == *strIn || strcase.Compare(s, *strIn) == 0 {
				return true
			}
		}
	}
	return false
}

// Check if string array contains a substring
//
// Return false if arrIn is nil
func ArrayContainsSubString(arrIn *[]string, strIn string, caseSensitive bool) bool {
	if arrIn != nil {
		for _, s := range *arrIn {
			if caseSensitive && strings.Contains(strIn, s) || strcase.Contains(s, strIn) {
				return true
			}
		}
	}
	return false
}

// Return a new *[]string with empty lines removed from *[]string.
//   - Original []string not modified.
func ArrayPtrRemoveEmpty(arrIn *[]string) *[]string {
	var arrOut []string
	if arrIn != nil {
		for _, s := range *arrIn {
			if s != "" {
				arrOut = append(arrOut, s)
			}
		}
	}
	return &arrOut
}

// *[]string to *string, each element followed by "\n"
func ArraySPrintln(arrIn *[]string) *string {
	var strOut string
	if arrIn != nil {
		for _, s := range *arrIn {
			strOut += s + "\n"
		}
	}
	return &strOut
}

// Check if string contains any substring of an array
//   - result (bool)
//   - if result is true, `resultVal` == matching substring
//   - if result is false, `resultVal` == ""
func ContainsAnySubStrings(strIn *string, subStrings *[]string, caseSensitive bool) (result bool, resultVal string) {
	// prefix := "matchList"
	if strIn != nil && subStrings != nil {
		for _, subStr := range *subStrings {
			if caseSensitive && strcase.Contains(*strIn, subStr) || strcase.Contains(*strIn, subStr) {
				result = true
				resultVal = subStr
				break
			}
		}
	}
	return result, resultVal
}

// Return only bool from [ContainsAnySubStrings]
func ContainsAnySubStringsBool(strIn *string, subStrings *[]string, caseSensitive bool) (result bool) {
	result, _ = ContainsAnySubStrings(strIn, subStrings, caseSensitive)
	return result
}

// *string to *[]string, split by "\n"
func LnSplit(strIn *string) *[]string {
	var strOut []string
	if strIn != nil {
		strOut = strings.Split(*strIn, "\n")
	}
	return &strOut
}

// Return *string of "" if failed
func JsonIndent(strIn *string) *string {
	var strOut string
	if strIn != nil {
		var byteA = []byte(*strIn)
		p := ByteJsonIndent(&byteA)
		if *p != "" {
			strOut = string(*p)
		} else {
			return strIn
		}
	}
	return &strOut
}

// Return *string of "" if json.Marshal failed
func JsonMarshal(strIn *string) *string {
	var strOut string
	if strIn != nil {
		p, e := json.Marshal(strIn)
		if e == nil {
			strOut = string(p)
		}
	}
	return &strOut
}

// Return *string of "" if json.Indent failed
func ByteJsonIndent(baP *[]byte) *string {
	var strOut string
	if baP != nil {
		var dst bytes.Buffer
		err := json.Indent(&dst, *baP, "", "  ")
		if err == nil {
			strOut = strings.Trim(dst.String(), "\n")
		}
	}
	return &strOut
}

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
