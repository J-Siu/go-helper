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
func StrArrayPtrContain(aP *[]string, sP *string) bool {
	for _, s := range *aP {
		if s == *sP {
			return true
		}
	}
	return false
}

// Return a new *[]string with empty lines removed from *[]string.
//   - Original []string not modified.
func StrArrayPtrRemoveEmpty(saP *[]string) *[]string {
	var sa []string
	for _, s := range *saP {
		if s != "" {
			sa = append(sa, s)
		}
	}
	return &sa
}

// *[]string output, each element followed by "\n"
func StrArrayPtrPrintln(saP *[]string) {
	for _, s := range *saP {
		println(s)
	}
}

// *[]string to *string, each element followed by "\n"
func StrArrayPtrPrintlnSp(saP *[]string) *string {
	var output string
	for _, s := range *saP {
		output += s + "\n"
	}
	return &output
}

// *string to *[]string, split by "\n"
func StrPtrToArrayPtr(sP *string) *[]string {
	var output []string = strings.Split(*sP, "\n")
	return &output
}
