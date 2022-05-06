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

// execute "git remote" using MyCmd
func GitRemote(v bool) *[]string {
	var args []string
	if v {
		args = []string{"remote", "-v"}
	} else {
		args = []string{"remote"}
	}
	output := MyCmdRun("git", &args).Stdout.String()
	return StrPToArrayP(&output)
}

// execute "git init" using MyCmd
func GitInit() *MyCmd {
	return MyCmdRun("git", &[]string{"init"})
}

// execute "git push" using MyCmd
func GitPush(optionsP []string) *MyCmd {
	args := []string{"push"}
	if optionsP != nil {
		args = append(args, optionsP...)
	}
	return MyCmdRun("git", &args)
}

// execute "git remote add" using MyCmd
func GitRemoteAdd(name string, git string) *MyCmd {
	return MyCmdRun("git", &[]string{"remote", "add", name, git})
}

// execute "git remote exit" using MyCmd
func GitRemoteExist(name string) bool {
	r := GitRemote(false)
	return StrArrayPtrContain(r, &name)
}

// execute "git remote remove" using MyCmd
func GitRemoteRemove(name string) *MyCmd {
	return MyCmdRun("git", &[]string{"remote", "remove", name})
}

// execute "git remote remove" all git remote using MyCmd
func GitRemoteRemoveAll() {
	gr := GitRemote(false)
	for _, r := range *gr {
		GitRemoteRemove(r)
	}
}
