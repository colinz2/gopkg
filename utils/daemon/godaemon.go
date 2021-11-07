package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GoDaemon() {
	args := os.Args[1:]
	i := 0
	for ; i < len(args); i++ {
		if strings.HasPrefix(args[i], "-d=") || args[i] == "-d" {
			args[i] = ""
			break
		}
	}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	//cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	err := cmd.Start()
	if err == nil {
		fmt.Println("[PID]", cmd.Process.Pid)
		cmd.Process.Release()
		os.Exit(0)
	}
}
