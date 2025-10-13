/*
The MIT License (MIT)

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

package file

import (
	"errors"
	"os"

	"github.com/J-Siu/go-helper/v2/basestruct"
	"github.com/J-Siu/go-helper/v2/errs"
	"github.com/J-Siu/go-helper/v2/ezlog"
)

// A very simple text file struct supporting read, write
type TxtFile struct {
	*basestruct.Base

	Content  string `json:"Content"`
	FilePath string `json:"FilePath"` //README path
}

func (t *TxtFile) New(filePath string) *TxtFile {
	t.Base = new(basestruct.Base)
	t.MyType = "TxtFile"
	prefix := t.MyType + ".New"

	t.FilePath = filePath

	ezlog.Debug().Nn(prefix).M(t).Out()
	if !IsRegularFile(t.FilePath) {
		t.Err = errors.New("TypeReadme.Init: " + t.FilePath + " not found")
		errs.Queue(prefix, t.Err)
	}
	t.Initialized = true
	return t
}

// Read `filePath` into `Content`
func (t *TxtFile) Read() *TxtFile {
	prefix := t.MyType + ".Read"
	if !t.CheckErrInit(prefix) {
		return t
	}
	var buf []byte
	if t.Err == nil {
		buf, t.Err = os.ReadFile(t.FilePath)
	}
	if t.Err == nil {
		t.Content = string(buf)
	}
	errs.Queue(prefix, t.Err)
	return t
}

// Write `Content` into `FilePath`
//
// `permission` is only used on new file
func (t *TxtFile) Write(permission os.FileMode) *TxtFile {
	prefix := t.MyType + ".Write"
	if !t.CheckErrInit(prefix) {
		return t
	}
	perm := permission
	if t.Err == nil {
		fileStats, err := os.Stat(t.FilePath)
		if err == nil {
			perm = fileStats.Mode()
		}
		t.Err = WriteStr(t.FilePath, &t.Content, perm)
	}
	errs.Queue(prefix, t.Err)
	return t
}
