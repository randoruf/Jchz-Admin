package system

import (
	"fmt"
	"os"
)

var (
	SysInfo   = "[System Info]:"
	CrashInfo = "[Crash Info]:"
	ErrInfo   = "[Error Info]:"
)

func PrintSysInfo(msg string) {
	fmt.Println(SysInfo, msg)
}
func PrintError(msg error) {
	fmt.Println(ErrInfo, msg)
}

func PrintCrash(msg error) {
	fmt.Println(CrashInfo, msg)
	os.Exit(1)
}
