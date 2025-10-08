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

// A simple array to stash errors
package errs

import (
	"errors"

	"github.com/J-Siu/go-helper/v2/array"
)

// error list
var Errs array.Array[error]

func Clear()         { Errs = nil }
func IsEmpty() bool  { return Errs.IsEmpty() }
func NotEmpty() bool { return Errs.NotEmpty() }
func Len() int       { return Errs.Len() }

// If `err` != nil, add `err` to helper Errs array
func Queue(prefix string, e error) {
	if e != nil {
		Errs = append(Errs, New(prefix, e.Error()))
	}
}

// shorthand for creating error with prefix
//
// if `prefix“ is not empty, returning error string will be prefixed by `<prefix>: `
//
// This does not add to the Errs queue
func New(prefix string, strErr string) error {
	if len(prefix) == 0 {
		return errors.New(strErr)
	} else {
		return errors.New(prefix + ": " + strErr)
	}
}
