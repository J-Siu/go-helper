package helper

import (
	"fmt"
	"testing"

	"github.com/J-Siu/go-helper"
)

// Exit code should be 0
func Test_ExitCode_0(t *testing.T) {
	helper.Debug = true
	var wanted int = 0
	var cmd string = "ls"
	var args []string = []string{"."}
	var myCmd *helper.MyCmd = helper.MyCmdRun(cmd, &args, nil)
	var msg int = myCmd.ExitCode
	fmt.Printf("myCmd.ExitCode() = %d\n", msg)
	if msg != wanted {
		t.Fatalf("myCmd.ExitCode() = %d, not %d", msg, wanted)
	}
}

// Exit code should not be 0
func Test_ExitCode_Not_Equal_0(t *testing.T) {
	helper.Debug = true
	var notWanted int = 0
	var cmd string = "ls"
	var args []string = []string{"__This_Path_Does_Not_Exist__"}
	var myCmd *helper.MyCmd = helper.MyCmdRun(cmd, &args, nil)
	var msg int = myCmd.ExitCode
	fmt.Printf("myCmd.Err = %s\n", myCmd.Err.Error())
	fmt.Printf("myCmd.ExitCode() = %d\n", msg)
	if msg == notWanted {
		t.Fatalf("myCmd.ExitCode() = %d (Should not be %d)", msg, notWanted)
	}
}
