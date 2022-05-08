package helper

import (
	"fmt"
	"testing"
)

func Test_CurrentDirBase(t *testing.T) {
	var wanted string = "go-helper"
	var msg *string = CurrentDirBase()
	fmt.Println("CurrentDirBase() is `" + *msg + "`")
	if *msg != wanted {
		t.Fatalf(`CurrentDirBase() = "%s", not "%s"`, *msg, wanted)
	}
}
