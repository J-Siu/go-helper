# Cmd

`exec.Command` shell wrapper.

## Installation

```sh
go get github.com/J-Siu/go-helper/v2
```

## Usage

```go
import "github.com/J-Siu/go-helper/v2/cmd"
```

## Types and Functions

### Structure

```go
type Cmd struct {
  ArgsP    *[]string    `json:"ArgsP"`    // In : Command args
  CmdLn    string       `json:"CmdLn"`    // Out: Command line
  CmdName  string       `json:"CmdName"`  // In : Command name
  Err      error        `json:"Err"`      // Out: run error
  ExitCode int          `json:"ExitCode"` // Out: Exit Code
  Ran      bool         `json:"Ran"`      // Out: Set to true by Run()
  Stderr   bytes.Buffer `json:"Stderr"`   // Out: Stderr
  Stdout   bytes.Buffer `json:"Stdout"`   // Out: Stdout
  WorkDir  string       `json:"WorkDir"`  // In : Command working dir
}
```

```go
func (t *Cmd) New(cmdName string, argsP *[]string, workPathP *string) *Cmd
func (t *Cmd) Run() *Cmd
// run func wrapper with sync.WaitGroup support.
func (t *Cmd) RunWg(title *string, wgP *sync.WaitGroup, output bool) *Cmd
func (t *Cmd) Error() error
 ```

### Package Functions

```go
func New(cmdName string, argsP *[]string, workPathP *string) *Cmd
// MyCmd run func wrapper.
func Run(cmdName string, argsP *[]string, workPathP *string) *Cmd
// MyCmd run func wrapper with sync.WaitGroup support.
func RunWg(cmdName string, argsP *[]string, workPathP *string, title *string, wgP *sync.WaitGroup, output bool) *Cmd
```

## License

The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
