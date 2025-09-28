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

// exec.Cmd wrapper
package cmd

import (
	"bytes"
	"os/exec"
	"sync"

	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/file"
)

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

// Setup and return cmd pointer.
//   - If <workPathP> is empty/nil, current directory is used.
func (c *Cmd) New(cmdName string, argsP *[]string, workPathP *string) *Cmd {
	c.ArgsP = argsP
	c.CmdName = cmdName
	c.ExitCode = 0
	c.WorkDir = *file.FullPath(workPathP)
	return c
}

// A exec.Cmd.Run() wrapper.
func (c *Cmd) Run() *Cmd {
	execCmd := exec.Command(c.CmdName, *c.ArgsP...)
	execCmd.Stdout = &c.Stdout
	execCmd.Stderr = &c.Stderr
	execCmd.Dir = c.WorkDir
	c.CmdLn = execCmd.String()
	c.Err = execCmd.Run()
	c.Ran = true

	if exitErr, ok := c.Err.(*exec.ExitError); ok {
		// fmt.Println("*** self.Err ok, get cmd exit code ***")
		c.ExitCode = exitErr.ExitCode()
	} else if c.Err != nil {
		// For simplicity, force exit code to 1
		// fmt.Println("*** self.Err != nil ***")
	}

	prefix := "cmd.Run"
	ezlog.Debug().Name(prefix).Msg(&c).Out()
	if c.Stderr.Len() > 0 {
		ezlog.Debug().Name(prefix).Name("Stderr").Msg(c.Stderr).Out()
	}
	if c.Stdout.Len() > 0 {
		ezlog.Debug().Name(prefix).Name("Stdout").Msg(c.Stdout).Out()
	}
	ezlog.Debug().Name(prefix).Name("Err").Msg(c.Err).Out()
	return c
}

// run func wrapper with sync.WaitGroup support.
//   - If <workPathP> is empty/nil, current directory is used.
//   - Print if <output> is true
func (c *Cmd) RunWg(title *string, wgP *sync.WaitGroup, output bool) *Cmd {
	if wgP != nil {
		defer wgP.Done()
	}
	c.Run()
	if output {
		if c.Stderr.Len() > 0 {
			ezlog.Log().Name(title).Name("Stderr").Msg(c.Stderr).Out()
		}
		if c.Stdout.Len() > 0 {
			ezlog.Log().Name(title).Name("Stdout").Msg(c.Stdout).Out()
		}
	}
	return c
}

func (c *Cmd) Error() error { return c.Err }

// Setup and return MyCmd pointer.
//   - If <workPathP> is empty/nil, current directory is used.
func New(cmdName string, argsP *[]string, workPathP *string) *Cmd {
	return new(Cmd).New(cmdName, argsP, workPathP)
}

// MyCmd run func wrapper.
//   - If <workPathP> is empty/nil, current directory is used.
func Run(cmdName string, argsP *[]string, workPathP *string) *Cmd {
	return RunWg(cmdName, argsP, workPathP, nil, nil, false)
}

// MyCmd run func wrapper with sync.WaitGroup support.
//   - If <workPathP> is empty/nil, current directory is used.
//   - Print if <output> is true
func RunWg(cmdName string, argsP *[]string, workPathP *string, title *string, wgP *sync.WaitGroup, output bool) *Cmd {
	return New(cmdName, argsP, workPathP).RunWg(title, wgP, output)
}
