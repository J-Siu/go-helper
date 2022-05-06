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

package helper

import (
	"bytes"
	"encoding/json"
	"os"
	"path"
	"strings"
)

func CurrentDirBase() string {
	d, _ := os.Getwd()
	return path.Base(d)
}

func StrArrayPtrContain(aP *[]string, sP *string) bool {
	for _, s := range *aP {
		if s == *sP {
			return true
		}
	}
	return false
}

// Remove empty lines from *[]string
func StrArrayPtrRemoveEmpty(saP *[]string) *[]string {
	// var sa []string
	for i := len(*saP) - 1; i >= 0; i-- {
		if (*saP)[i] == "" {
			(*saP) = append((*saP)[:i], (*saP)[i+1:]...)
		}
	}
	return saP
}

// *[]string output
func StrArrayPtrPrintln(saP *[]string) {
	for _, s := range *saP {
		println(s)
	}
}

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

// bool -> "Yes/No"
func BoolYesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

// *string -> *[]string
func StrPToArrayP(sP *string) *[]string {
	r := strings.Split(*sP, "\n")
	return StrArrayPtrRemoveEmpty(&r)
}

// Json indent *[]byte -> *string
func JsonIndentSp(baP *[]byte, endl bool) *string {
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
	if len(output) > 0 && output[len(output)-1] != '\n' && endl {
		output += "\n"
	}
	return &output
}
