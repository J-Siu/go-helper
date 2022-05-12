package helper

import (
	"fmt"
	"testing"

	"github.com/J-Siu/go-helper"
)

func Test_CurrentDirBase(t *testing.T) {
	var wanted string = "test"
	var msg *string = helper.CurrentDirBase()
	fmt.Println("CurrentDirBase() is `" + *msg + "`")
	if *msg != wanted {
		t.Fatalf(`CurrentDirBase() = "%s", not "%s"`, *msg, wanted)
	}
}
