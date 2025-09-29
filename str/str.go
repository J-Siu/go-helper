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
func ArrayContains(arrIn *[]string, strIn *string) bool {
	for _, s := range *arrIn {
		if s == *strIn {
			return true
		}
	}
	return false
}

// Check if string array contains a substring
func ArrayContainsSubString(arrIn *[]string, strIn string) bool {
	for _, s := range *arrIn {
		if strings.Contains(strIn, s) {
			return true
		}
	}
	return false
}

// Return a new *[]string with empty lines removed from *[]string.
//   - Original []string not modified.
func ArrayPtrRemoveEmpty(arrIn *[]string) *[]string {
	var arrOut []string
	for _, s := range *arrIn {
		if s != "" {
			arrOut = append(arrOut, s)
		}
	}
	return &arrOut
}

// *[]string output, each element followed by "\n"
func ArrayPrintln(arrIn *[]string) {
	for _, s := range *arrIn {
		println(s)
	}
}

// *[]string to *string, each element followed by "\n"
func ArraySPrintln(arrIn *[]string) *string {
	var strOut string
	for _, s := range *arrIn {
		strOut += s + "\n"
	}
	return &strOut
}

// Check if string contains any substring of an array
//   - result (bool)
//   - if result is true, `resultVal` == matching substring
//   - if result is false, `resultVal` == ""
func ContainsAnySubStrings(strIn *string, subStrings *[]string) (result bool, resultVal string) {
	// prefix := "matchList"
	for _, subStr := range *subStrings {
		if strcase.Contains(*strIn, subStr) {
			result = true
			resultVal = subStr
			break
		}
	}
	return result, resultVal
}

// Return only bool from [ContainsAnySubStrings]
func ContainsAnySubStringsBool(strIn *string, subStrings *[]string) (result bool) {
	result, _ = ContainsAnySubStrings(strIn, subStrings)
	return result
}

// *string to *[]string, split by "\n"
func LnSplit(strIn *string) *[]string {
	var strOut []string = strings.Split(*strIn, "\n")
	return &strOut
}

// Return original *string if strP is nil failed
func JsonIndent(strIn *string) *string {
	var output string
	if strIn != nil {
		var byteA = []byte(*strIn)
		p := ByteJsonIndent(&byteA)
		if *p != "" {
			output = string(*p)
		} else {
			return strIn
		}
	}
	return &output
}

// Return "" if json.Indent failed
func ByteJsonIndent(baP *[]byte) *string {
	var output string
	if baP != nil {
		var dst bytes.Buffer
		err := json.Indent(&dst, *baP, "", "  ")
		if err == nil {
			output = strings.Trim(dst.String(), "\n")
		}
	}
	return &output
}
