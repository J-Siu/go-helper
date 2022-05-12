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
	//   git submodule add https://github.com/J-Siu/submodule_test_1
	args = []string{"submodule", "add", "https://github.com/J-Siu/submodule_test_1"}
	helper.MyCmdRun("git", &args, nil)
	//   git submodule update --recursive --init
	args = []string{"submodule", "update", "--recursive", "--init"}
	helper.MyCmdRun("git", &args, nil)

	// testDir should be set to a git submodule directory.
	var testDir string = "submodule_test_1/submodule_test_2/submodule_test_3/submodule_test_4/submodule_test_5"
	var wanted string = "go-helper"
	var msg string = path.Base(helper.GitRoot(&testDir))
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
