/*
Copyright © 2022 John, Sing Dao, Siu <john.sd.siu@gmail.com>

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

// string.go - a collection of string related helper functions
package helper

import (
	"bytes"
	"encoding/json"
	"strings"
)

// bool to "true"/"false"
func BoolString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// bool to "OK"/"Fail"
func BoolStatus(b bool) string {
	if b {
		return "ok"
	}
	return "fail"
}

// bool to "Yes/No"
func BoolYesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

// Json indent *[]byte to *string.
//  - If <endLn> is true, add new line at end of string if not exist.
//  - Return string pointer.
func JsonIndentSp(baP *[]byte, endLn bool) *string {
	var output string
	if len(*baP) > 0 {
		var dst bytes.Buffer
		err := json.Indent(&dst, *baP, "", "  ")
		if err == nil {
			output = dst.String()
		} else {
			output = string(*baP)
		}
	}
	if len(output) > 0 && output[len(output)-1] != '\n' && endLn {
		output += "\n"
	}
	return &output
}

// Check if string array contain a string.
func StrArrayPtrContain(aP *[]string, sP *string) bool {
	for _, s := range *aP {
		if s == *sP {
			return true
		}
	}
	return false
}

// Return a new *[]string with empty lines removed from *[]string.
//  - Original []string not modified.
func StrArrayPtrRemoveEmpty(saP *[]string) *[]string {
	// // This is faster but modify original array *saP
	// for i := len(*saP) - 1; i >= 0; i-- {
	// 	if (*saP)[i] == "" {
	// 		(*saP) = append((*saP)[:i], (*saP)[i+1:]...)
	// 	}
	// }
	var sa []string
	var l int = len(*saP)
	for i := 0; i < l; i++ {
		if (*saP)[i] != "" {
			sa = append(sa, (*saP)[i])
		}
	}
	return &sa
}

// *[]string output
func StrArrayPtrPrintln(saP *[]string) {
	for _, s := range *saP {
		println(s)
	}
}

// *string output
func StrArrayPtrPrintlnSp(saP *[]string) *string {
	var output string
	for _, s := range *saP {
		output += s + "\n"
	}
	return &output
}

// Split *string to *[]string by new line("\n")
func StrPtrToArrayPtr(sP *string) *[]string {
	var output []string = strings.Split(*sP, "\n")
	return &output
}

// Json indent *string to *string.
//  - If <endLn> is true, add new line at end of string if not exist.
//  - Return string pointer.
func StrPtrToJsonIndentSp(strP *string, endLn bool) *string {
	var byteA = []byte(*strP)
	return JsonIndentSp(&byteA, endLn)
}

// Json indent *string to *string.
//  - If <endLn> is true, add new line at end of string if not exist.
//  - Return string pointer.
func StrToJsonIndentSp(str string, endLn bool) *string {
	var byteA = []byte(str)
	return JsonIndentSp(&byteA, endLn)
}

// Json indent any to *string.
//  - If <endLn> is true, add new line at end of string if not exist.
//  - Return string pointer.
func AnyToJsonIndentSp(data any, endLn bool) *string {
	var str string
	j, e := json.MarshalIndent(data, "", "  ")
	if e != nil {
		Errs = append(Errs, e)
	} else if len(j) > 0 {
		str = string(j)
		if str[len(str)-1] != '\n' {
			str += "\n"
		}
	}
	return &str
}
