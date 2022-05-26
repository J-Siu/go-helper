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

// A exec.Cmd wrapper
type MyCmd struct {
	ArgsP   *[]string    `json:"ArgsP"`   // In : Command args
	CmdLn   string       `json:"CmdLn"`   // Out: Command line
	CmdName string       `json:"CmdName"` // In : Command name
	Err     error        `json:"Err"`     // Out: run error
	Ran     bool         `json:"Ran"`     // Out: Set to true by Run()
	Stderr  bytes.Buffer `json:"Stderr"`  // Out: Stderr
	Stdout  bytes.Buffer `json:"Stdout"`  // Out: Stdout
	WorkDir string       `json:"WorkDir"` // In : Command working dir
}

// MyCmd run func wrapper.
//  - If <workPathP> is empty/nil, current directory is used.
func MyCmdRun(cmdName string, argsP *[]string, workPathP *string) *MyCmd {
	return MyCmdRunWg(cmdName, argsP, workPathP, nil, nil, false)
}

// MyCmd run func wrapper with sync.WaitGroup support.
//  - If <workPathP> is empty/nil, current directory is used.
//  - Print if <output> is true
func MyCmdRunWg(cmdName string, argsP *[]string, workPathP *string, title *string, wgP *sync.WaitGroup, output bool) *MyCmd {
	if wgP != nil {
		defer wgP.Done()
	}
	var self *MyCmd = MyCmdInit(cmdName, argsP, workPathP)
	self.Run()
	if output {
		Report(self.Stdout.String(), *title+":Stdout", true, false)
		Report(self.Stderr.String(), *title+":Stderr", true, false)
	}
	return self
}

// Setup and return MyCmd pointer.
//  - If <workPathP> is empty/nil, current directory is used.
func MyCmdInit(name string, argsP *[]string, workPathP *string) *MyCmd {
	var self MyCmd
	return self.Init(name, argsP, workPathP)
}

// Setup and return MyCmd pointer.
//  - If <workPathP> is empty/nil, current directory is used.
func (self *MyCmd) Init(name string, argsP *[]string, workPathP *string) *MyCmd {
	self.ArgsP = argsP
	self.CmdName = name
	self.WorkDir = *FullPath(workPathP)
	return self
}

// Reset MyCmd
func (self *MyCmd) Reset() *MyCmd {
	var myCmd MyCmd
	*self = myCmd
	return self
}

// A exec.Cmd.Run() wrapper.
func (self *MyCmd) Run() error {
	execCmd := exec.Command(self.CmdName, *self.ArgsP...)
	execCmd.Stdout = &self.Stdout
	execCmd.Stderr = &self.Stderr
	execCmd.Dir = self.WorkDir
	self.CmdLn = execCmd.String()
	self.Err = execCmd.Run()
	self.Ran = true
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
