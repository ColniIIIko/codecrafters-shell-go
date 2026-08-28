package cmd

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

func Cd(args []string, ctx utils.Shell) (string, error) {
	arg := args[0]
	resPath := utils.ResolvePath(arg)

	err := os.Chdir(resPath)

	if err != nil {
		return "", fmt.Errorf("cd: %s: No such file or directory\n", arg)
	}

	return "", nil
}
