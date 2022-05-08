package helper

import (
	"fmt"
	"testing"
)

func Test_CurrentDirBase(t *testing.T) {
	var msg string = CurrentDirBase()
	fmt.Println("CurrentDirBase() is `" + msg + "`")
	if msg != "go-helper" {
		t.Fatalf(`CurrentDirBase() = "%s"`, msg)
	}
}
