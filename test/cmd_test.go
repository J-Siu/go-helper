package helper

import (
	"fmt"
	"testing"

	"github.com/J-Siu/go-helper/v2/cmd"
)

// Exit code should be 0
func Test_ExitCode_0(t *testing.T) {
	var (
		wanted int    = 0
		cli    string = "ls"
		args          = []string{"."}
		myCmd         = cmd.New(cli, &args, nil).Run()
		msg    int    = myCmd.ExitCode
	)
	fmt.Printf("myCmd.ExitCode() = %d\n", msg)
	if msg != wanted {
		t.Fatalf("myCmd.ExitCode() = %d, not %d", msg, wanted)
	}
}

// Exit code should not be 0
func Test_ExitCode_Not_Equal_0(t *testing.T) {
	var (
		notWanted int    = 0
		cli       string = "ls"
		args             = []string{"__This_Path_Does_Not_Exist__"}
		myCmd            = cmd.New(cli, &args, nil).Run()
		msg       int    = myCmd.ExitCode
	)
	fmt.Printf("myCmd.Err = %s\n", myCmd.Err.Error())
	fmt.Printf("myCmd.ExitCode() = %d\n", msg)
	if msg == notWanted {
		t.Fatalf("myCmd.ExitCode() = %d (Should not be %d)", msg, notWanted)
	}
}
