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
	"strings"
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
func ArrayContainsSubString(arrIn []string, strIn string) bool {
	for _, s := range arrIn {
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

// *string to *[]string, split by "\n"
func LnSplit(strIn *string) *[]string {
	var strOut []string = strings.Split(*strIn, "\n")
	return &strOut
}
