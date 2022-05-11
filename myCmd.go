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
	"errors"
	"os/exec"
	"sync"
)

// go exec.Cmd wrapper
type MyCmd struct {
	ArgsP   *[]string    `json:"ArgsP"`  // Command args
	Name    string       `json:"Name"`   // Command name
	WorkDir string       `json:"Dir"`    // Command working dir
	CmdLn   string       `json:"CmdLn"`  // Out: Command Line
	Err     error        `json:"Err"`    // Out: run error
	Stdout  bytes.Buffer `json:"Stdout"` // Out: Stdout
	Stderr  bytes.Buffer `json:"Stderr"` // Out: Stderr
}

// MyCmd run func wrapper with working directory support.
// If `workpath` is empty, current directory is used.
func MyCmdRun(name string, argsP *[]string, workpathP *string) *MyCmd {
	return MyCmdRunWg(name, argsP, workpathP, nil, nil, false)
}

// MyCmd run func wrapper with sync.WaitGroup and working directory support
// If `workpath` is empty, current directory is used.
func MyCmdRunWg(name string, argsP *[]string, workpathP *string, title *string, wgP *sync.WaitGroup, output bool) *MyCmd {
	if wgP != nil {
		defer wgP.Done()
	}
	var self MyCmd = *MyCmdInit(name, argsP, workpathP)
	self.Run()
	if output {
		Report(self.Stdout.String(), *title+":Stdout", true, false)
		Report(self.Stderr.String(), *title+":Stderr", true, false)
	}
	return &self
}

// Setup a MyCmd and return pointer to it
// If `workpath` is empty, current directory is used.
func MyCmdInit(name string, argsP *[]string, workpathP *string) *MyCmd {
	var self MyCmd
	self.ArgsP = argsP
	self.Name = name
	self.WorkDir = *FullPath(workpathP)
	return &self
}

// exec.Cmd.Run() wrapper
func (self *MyCmd) Run() error {
	execCmd := exec.Command(self.Name, *self.ArgsP...)
	execCmd.Stdout = &self.Stdout
	execCmd.Stderr = &self.Stderr
	execCmd.Dir = self.WorkDir
	self.CmdLn = execCmd.String()
	self.Err = execCmd.Run()
	ReportDebug(&self, "myCmd:", false, false)
	ReportDebug(self.Stderr.String(), "myCmd:Stderr", false, false)
	ReportDebug(self.Stdout.String(), "myCmd:Stdout", false, false)
	return self.Err
}

// Return exit code
func (self *MyCmd) ExitCode() int {
	var exitErr *exec.ExitError
	if errors.As(self.Err, &exitErr) {
		return exitErr.ExitCode()
	}
	// No error
	return 0
}
