package helper

import (
	"fmt"
	"testing"
)

// Exit code should be 0
func Test_ExitCode_0(t *testing.T) {
	Debug = true
	var wanted int = 0
	var cmd string = "cd"
	var args []string = []string{"."}
	var myCmd *MyCmd = MyCmdRun(cmd, &args, nil)
	var msg int = myCmd.ExitCode()
	fmt.Printf("myCmd.ExitCode() = %d\n", msg)
	if msg != wanted {
		t.Fatalf("myCmd.ExitCode() = %d, not %d", msg, wanted)
	}
}

// Exit code should be 1
func Test_ExitCode_1(t *testing.T) {
	Debug = true
	var wanted int = 1
	var cmd string = "cd"
	var args []string = []string{"__This_Path_Does_Not_Exist__"}
	var myCmd *MyCmd = MyCmdRun(cmd, &args, nil)
	var msg int = myCmd.ExitCode()
	fmt.Printf("myCmd.ExitCode() = %d\n", msg)
	if msg != wanted {
		t.Fatalf("myCmd.ExitCode() = %d, not %d", msg, wanted)
	}
}
