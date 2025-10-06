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

package cmd

import (
	"bytes"
	"os/exec"
	"sync"

	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/file"
)

type Cmd struct {
	ArgsP    *[]string    `json:"ArgsP,omitempty"`    // In : Command args
	CmdLn    string       `json:"CmdLn,omitempty"`    // Out: Command line
	CmdName  string       `json:"CmdName,omitempty"`  // In : Command name
	Err      error        `json:"Err,omitempty"`      // Out: run error
	ExitCode int          `json:"ExitCode,omitempty"` // Out: Exit Code
	Ran      bool         `json:"Ran,omitempty"`      // Out: Set to true by Run()
	Stderr   bytes.Buffer `json:"Stderr,omitempty"`   // Out: Stderr
	Stdout   bytes.Buffer `json:"Stdout,omitempty"`   // Out: Stdout
	WorkDir  string       `json:"WorkDir,omitempty"`  // In : Command working dir
}

// Setup and return cmd pointer.
//   - If <workPathP> is empty/nil, current directory is used.
func (t *Cmd) New(cmdName string, argsP *[]string, workPathP *string) *Cmd {
	t.ArgsP = argsP
	t.CmdName = cmdName
	t.ExitCode = 0
	t.WorkDir = *file.FullPath(workPathP)
	return t
}

// A exec.Cmd.Run() wrapper.
func (t *Cmd) Run() *Cmd {
	execCmd := exec.Command(t.CmdName, *t.ArgsP...)
	execCmd.Stdout = &t.Stdout
	execCmd.Stderr = &t.Stderr
	execCmd.Dir = t.WorkDir
	t.CmdLn = execCmd.String()
	t.Err = execCmd.Run()
	t.Ran = true

	if exitErr, ok := t.Err.(*exec.ExitError); ok {
		// fmt.Println("*** self.Err ok, get cmd exit code ***")
		t.ExitCode = exitErr.ExitCode()
	} else if t.Err != nil {
		// For simplicity, force exit code to 1
		// fmt.Println("*** self.Err != nil ***")
	}

	prefix := "cmd.Run"
	ezlog.Debug().Name(prefix).Msg(&t).Out()
	if t.Stdout.Len() > 0 {
		ezlog.Debug().Name(prefix).Name("Stdout").Msg(t.Stdout).Out()
	}
	if t.Stderr.Len() > 0 {
		ezlog.Debug().Name(prefix).Name("Stderr").Msg(t.Stderr).Out()
	}
	ezlog.Debug().Name(prefix).Name("Err").Msg(t.Err).Out()
	return t
}

// run func wrapper with sync.WaitGroup support.
//   - If <workPathP> is empty/nil, current directory is used.
//   - Print if <output> is true
func (t *Cmd) RunWg(title *string, wgP *sync.WaitGroup, output bool) *Cmd {
	if wgP != nil {
		defer wgP.Done()
	}
	t.Run()
	if output {
		if t.Stdout.Len() > 0 {
			ezlog.Log().Name(title).Name("Stdout").Msg(t.Stdout).Out()
		}
		if t.Stderr.Len() > 0 {
			ezlog.Log().Name(title).Name("Stderr").Msg(t.Stderr).Out()
		}
	}
	return t
}

func (t *Cmd) Error() error { return t.Err }
