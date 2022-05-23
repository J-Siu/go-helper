package helper

import (
	"fmt"
	"path"
	"testing"

	"github.com/J-Siu/go-helper"
)

func TestGitRoot_hasSubmodule(t *testing.T) {
	// Debug = true

	// Pull submodules
	var args []string
	var myCmd *helper.MyCmd
	//		git submodule add -f https://github.com/J-Siu/submodule_test_1
	args = []string{"submodule", "add", "-f", "https://github.com/J-Siu/submodule_test_1"}
	myCmd = helper.MyCmdRun("git", &args, nil)
	helper.Report(myCmd, "myCmd", false, false)
	helper.Report(myCmd.Stderr, "myCmd.Stderr", true, false)
	helper.Report(myCmd.Stdout, "myCmd.Stdout", true, false)
	//   git submodule update --recursive --init
	args = []string{"submodule", "update", "--recursive", "--init"}
	myCmd = helper.MyCmdRun("git", &args, nil)
	helper.Report(myCmd, "myCmd", false, false)
	helper.Report(myCmd.Stderr, "myCmd.Stderr", true, false)
	helper.Report(myCmd.Stdout, "myCmd.Stdout", true, false)

	// Test
	//		testDir should be set to a git submodule directory.
	var testDir string = "submodule_test_1/submodule_test_2/submodule_test_3/submodule_test_4/submodule_test_5"
	var wanted string = "go-helper"
	var msg string = path.Base(helper.GitRoot(&testDir))

	// Clean up
	//		git rm -rf submodule_test_1 ../.gitmodules
	args = []string{"rm", "-rf", "submodule_test_1", "../.gitmodules"}
	myCmd = helper.MyCmdRun("git", &args, nil)

	// Result
	fmt.Println("Git root of " + testDir + " is `" + msg + "`")
	if msg != wanted {
		t.Fatalf(`GitRoot(rtSp(%s) = ""`, testDir)
	}
}

func TestGitRoot_fail(t *testing.T) {
	// Assuming root directory is not a git directory
	var testDir string = "/"
	var msg string = helper.GitRoot(&testDir)
	fmt.Println("Git root of " + testDir + " is `" + msg + "`")
	if msg != "" {
		t.Fatalf(`GitRoot(rtSp(%s) = ""`, testDir)
	}
}
