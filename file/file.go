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

// collection of file helper functions
package file

import (
	"os"
	"path"
	"regexp"
	"strings"
)

// Get full path of current directory.
//   - Return string pointer.
func CurrentPath() *string {
	p, _ := os.Getwd()
	return &p
}

// Get current directory name. Not full path.
//   - Return string pointer.
func CurrentDirBase() *string {
	p := path.Base(*CurrentPath())
	return &p
}

// Return full path of path provided.
//   - If <workPathP> is empty/nil, use current path.
//   - <workPathP> not modified.
//   - Return string pointer.
func FullPath(workPathP *string) *string {
	var p string
	if workPathP == nil || *workPathP == "" {
		// Empty return current path
		p = *CurrentPath()
	} else if (*workPathP)[0] == '/' {
		// Path start with / already full path, return it
		p = *workPathP
	} else {
		// Add to current path
		p = path.Join(*CurrentPath(), *workPathP)
	}
	return &p
}

// Return full path of path provided.
//   - If <workPath> is empty, use current path.
//   - Return string pointer.
func FullPathStr(workPath string) *string {
	return FullPath(&workPath)
}

// Check workPath is regular file
//
// Return false on stat() error
func IsRegularFile(workPath string) (result bool) {
	fileInfo, err := os.Stat(workPath)
	if err == nil {
		result = fileInfo.Mode().IsRegular()
	}
	return result
}

// --- Directory

// Search for file in a directory
//   - filename path info removed automatically, only base is used
//   - case insensitive search
//   - return actually filename if found
//   - return empty if not found or error
//   - error is added to Errs
func InDir(dir, filename string) (result string) {
	fileBase := path.Base(filename)
	files, err := os.ReadDir(dir)
	if err == nil {
		for _, f := range files {
			if strings.EqualFold(fileBase, f.Name()) {
				// case insensitive matched, return real name
				return f.Name()
			}
		}
	}
	return result
}

// Check workPath is directory
func IsDir(workPath string) (result bool) {
	fileInfo, err := os.Stat(workPath)
	if err == nil {
		result = fileInfo.IsDir()
	}
	return result
}

// Check two paths have same parent directory
func SameDir(path1, path2 string) bool {
	return path.Dir(path1) == path.Dir(path2)
}

// --- Extension

// Check file has supplied extension
func ExtHas(name, ext string) bool {
	return strings.EqualFold(path.Ext(name), ext)
}

// Return filename/path with extension removed
func ExtRemove(filename string) string {
	return strings.TrimSuffix(filename, path.Ext(filename))
}

// --- Filename

// Apply following to filename:
//   - remove extension
//   - to lowercase
//   - remove -,_
func SimplifyName(filename string) string {
	filename = ExtRemove(filename)
	filename = strings.ToLower(filename)
	filename = strings.ReplaceAll(filename, "-", "")
	filename = strings.ReplaceAll(filename, "_", "")
	return filename
}

// Expand Linux `~` and environment variable in string
func TildeEnvExpand(strIn string) (strOut string) {
	if strIn == "~" {
		strOut = "$HOME"
	} else {
		re := regexp.MustCompile(`^~/`)
		strOut = re.ReplaceAllString(strIn, "$$HOME/")
	}
	return os.ExpandEnv(strOut)
}

// --- Array

// Read file into []string.
//
// Lines are split by "\n".
func ArrayRead(filePath string) (strArray []string, err error) {
	var byteRead []byte
	byteRead, err = os.ReadFile(filePath)
	if err == nil {
		strArray = strings.Split(string(byteRead), "\n")
	}
	return strArray, err
}

// Write []string into file
func ArrayWrite(filePath string, strArray []string, perm os.FileMode) (err error) {
	err = os.WriteFile(filePath, []byte(strings.Join(strArray, "\n")), perm)
	return err
}
