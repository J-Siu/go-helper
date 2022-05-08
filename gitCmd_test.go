package helper

import (
	"fmt"
	"testing"
)

func TestGitRoot_hasSubmodule(t *testing.T) {
	Debug = true

	// testDir should be set to a git submodule directory.
	var testDir string = "/Users/js/code/go/src/github.com/J-Siu/t2/t3/t4"
	var msg string = GitRoot(&testDir)
	fmt.Println("Git root of " + testDir + " is `" + msg + "`")
	if msg == "" {
		t.Fatalf(`GitRoot(rtSp(%s) = ""`, testDir)
	}
}

func TestGitRoot_fail(t *testing.T) {
	// Assuming root directory is not a git directory
	var testDir string = "/"
	var msg string = GitRoot(&testDir)
	fmt.Println("Git root of " + testDir + " is `" + msg + "`")
	if msg != "" {
		t.Fatalf(`GitRoot(rtSp(%s) = ""`, testDir)
	}
}
