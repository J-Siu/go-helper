# File

Collection of file and directory helper functions.

## Installation

```sh
go get github.com/J-Siu/go-helper/v2
```

## Usage

```go
import "github.com/J-Siu/go-helper/v2/file"
```

## Types and Functions

### TxtFile

A very simple text file struct supporting read, write.

```go
type TxtFile struct {
  *basestruct.Base

  Content  string `json:"Content"`
  FilePath string `json:"FilePath"` //README path
}
```

```go
func (t *TxtFile) New(filePath string) *TxtFile
func (t *TxtFile) Read() *TxtFile
func (t *TxtFile) Write(permission os.FileMode) *TxtFile
```

### Package Functions

```go
func CurrentPath() *string
func CurrentDirBase() *string
func FullPath(workPathP *string) *string
func FullPathStr(workPath string) *string
func IsRegularFile(workPath string) (result bool)
func FileSame(file1 string, file2 string) (same bool)
func InDir(dir, filename string) (result string)
func IsDir(path string) (result bool)
func SameDir(path1, path2 string) bool
func GetDirFile(dir string) (*[]string, *[]string)
func FindFile(dir, filename string, caseSensitive bool) string
func ExtHas(name, ext string) bool
func ExtRemove(filename string) string
func SimplifyName(filename string) string
func TildeEnvExpand(strIn string) (strOut string)
func AppendByte(filePath string, bP *[]byte) (err error)
func AppendStr(filePath string, str *string) (err error)
func AppendStrArray(filePath string, strArray *[]string) (err error)
func ReadByte(filePath string) (*[]byte, error)
func ReadStr(filePath string) (*string, error)
func ReadStrArray(filePath string) (*[]string, error)
func WriteByte(filePath string, bP *[]byte, perm os.FileMode) error
func WriteStr(filePath string, str *string, perm os.FileMode) error
func WriteStrArray(filePath string, strArray *[]string, perm os.FileMode) error
```

## License

The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
