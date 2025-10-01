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

package bigObj

import (
	"errors"
	"strconv"

	"github.com/J-Siu/go-helper/v2/basestruct"
)

type BigObject struct {
	basestruct.Base

	length int
	width  int
	height int

	name *string
}

func New() *BigObject {
	b := new(BigObject)

	// Init basestruct value
	b.Initialized = true
	b.MyType = "BigObject"

	return b
}

func (b *BigObject) Length() int { return b.length }
func (b *BigObject) Width() int  { return b.width }
func (b *BigObject) Height() int { return b.height }
func (b *BigObject) Name() string {
	prefix := b.MyType + ".Name"
	if b.CheckErrInit(prefix) {
		if b.name == nil {
			b.Err = errors.New(prefix + ": Name is not set.")
			return ""
		}
	}
	return *b.name
}

func (b *BigObject) SetLength(v int) *BigObject {
	prefix := b.MyType + ".SetLength"
	if b.CheckErrInit(prefix) {
		b.length = v
	}
	return b
}
func (b *BigObject) SetWidth(v int) *BigObject {
	prefix := b.MyType + ".SetWidth"
	if b.CheckErrInit(prefix) {
		b.width = v
	}
	return b
}
func (b *BigObject) SetHeight(v int) *BigObject {
	prefix := b.MyType + ".SetHeight"
	if b.CheckErrInit(prefix) {
		b.height = v
	}
	return b
}
func (b *BigObject) SetName(v string) *BigObject {
	prefix := b.MyType + ".SetName"
	if b.CheckErrInit(prefix) {
		b.name = &v
	}
	return b
}

func (b *BigObject) String() string {
	prefix := b.MyType + ".String"
	var str string
	if b.CheckErrInit(prefix) {
		if b.name == nil {
			b.Err = errors.New(prefix + ": Name is not set.")
		} else {
			str = *b.name + ": H: " + strconv.Itoa(b.height) + " L: " + strconv.Itoa(b.length) + " W: " + strconv.Itoa(b.width)
		}
	}
	return str
}
