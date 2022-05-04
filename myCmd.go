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
	"os/exec"
	"sync"
)

type MyCmd struct {
	ArgsP  *[]string    `json:"ArgsP"`  // Command args
	Name   string       `json:"Name"`   // Command name
	CmdLn  string       `json:"CmdLn"`  // Out: Command Line
	Err    error        `json:"Err"`    // Out: run error
	Stdout bytes.Buffer `json:"Stdout"` // Out: Stdout
	Stderr bytes.Buffer `json:"Stderr"` // Out: Stderr
}

func MyCmdRun(name string, argsP *[]string) *MyCmd {
	var self MyCmd
	self.ArgsP = argsP
	self.Name = name
	self.Run()
	return &self
}

func MyCmdRunWg(name string, argsP *[]string, title *string, wgP *sync.WaitGroup, output bool) *MyCmd {
	var self MyCmd
	self.ArgsP = argsP
	self.Name = name
	self.Run()
	if output {
		Report(self.Stdout.String(), *title+":StdOut", true)
		Report(self.Stderr.String(), *title+":StdErr", true)
	}
	return &self
}

func (self *MyCmd) Run() error {
	execCmd := exec.Command(self.Name, *self.ArgsP...)
	execCmd.Stdout = &self.Stdout
	execCmd.Stderr = &self.Stderr
	self.CmdLn = execCmd.String()
	self.Err = execCmd.Run()
	ReportDebug(&self, "myCmdP:", false)
	ReportDebug(self.Stdout.String(), "myCmdP.Stdout", false)
	ReportDebug(self.Stderr.String(), "myCmdP.Stderr", false)
	return self.Err
}
