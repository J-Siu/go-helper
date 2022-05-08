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
	"os/exec"
	"strings"
)

// Check git executable exit
func GitExecExist() bool {
	return GitExecPath() != ""
}

// Check git executable path
func GitExecPath() string {
	path, err := exec.LookPath("git")
	if err != nil {
		return ""
	}
	return path
}

// execute "git init" using MyCmd
func GitInit(workpathP *string) *MyCmd {
	return MyCmdRun("git", &[]string{"init"}, workpathP)
}

// execute "git push" using MyCmd
func GitPush(workpathP *string, optionsP []string) *MyCmd {
	args := []string{"push"}
	if optionsP != nil {
		args = append(args, optionsP...)
	}
	return MyCmdRun("git", &args, workpathP)
}

// execute "git remote" using MyCmd
func GitRemote(workpathP *string, v bool) *[]string {
	var args []string
	if v {
		args = []string{"remote", "-v"}
	} else {
		args = []string{"remote"}
	}
	output := MyCmdRun("git", &args, workpathP).Stdout.String()
	return StrPToArrayP(&output)
}

// execute "git remote add" using MyCmd
func GitRemoteAdd(workpathP *string, name string, git string) *MyCmd {
	return MyCmdRun("git", &[]string{"remote", "add", name, git}, workpathP)
}

// execute "git remote exit" using MyCmd
func GitRemoteExist(workpathP *string, name string) bool {
	r := GitRemote(workpathP, false)
	return StrArrayPtrContain(r, &name)
}

// execute "git remote remove" using MyCmd
func GitRemoteRemove(workpathP *string, name string) *MyCmd {
	return MyCmdRun("git", &[]string{"remote", "remove", name}, workpathP)
}

// execute "git remote remove" all git remote using MyCmd
func GitRemoteRemoveAll(workpathP *string) {
	gr := GitRemote(workpathP, false)
	for _, r := range *gr {
		GitRemoteRemove(workpathP, r)
	}
}

// Get git root from current directory.
// Return empty string if not a git dir.
func GitRoot(workpathP *string) string {
	if *workpathP == "" {
		CurrentPath()
	}
	// Check submodule path repeatly
	var submodulePath string = *workpathP
	var currPath string = *workpathP
	for submodulePath != "" {
		submodulePath = GitRootSubmodule(&submodulePath)
		if submodulePath != "" {
			currPath = submodulePath
		}
	}
	// Check git root
	var opts []string = []string{"rev-parse", "--show-toplevel"}
	var myCmd *MyCmd = MyCmdRun("git", &opts, &currPath)
	if myCmd.Err != nil {
		return ""
	}
	return strings.TrimSpace(myCmd.Stdout.String())
}

// Get git submodule root from `workpath`.
// If `workpath` is empty, current directory is used.
// Return empty string if not a submodule dir.
func GitRootSubmodule(workpathP *string) string {
	var opts []string = []string{"rev-parse", "--show-superproject-working-tree"}
	var myCmd *MyCmd = MyCmdRun("git", &opts, workpathP)
	if myCmd.Err != nil {
		return ""
	}
	return strings.TrimSpace(myCmd.Stdout.String())
}
