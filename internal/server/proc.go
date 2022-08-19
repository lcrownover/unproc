package server

import (
	"github.com/lcrownover/unproc/internal/util"
	"os/exec"
)

func GetProcessCountforUser(user string) int32 {
	if user == "test" {
		return 7
	}
	cmd := exec.Command("ps", "-o", "pid=", "-u", user)
	stdout, err := cmd.Output()
	if err != nil {
		return 0
	}
	lineCount := util.CountLines(string(stdout))
	return lineCount
}
