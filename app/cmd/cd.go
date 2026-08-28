package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

func Cd(args []string, ctx utils.Shell) (string, error) {
	arg := args[0]

	var resPath string

	if strings.HasPrefix(arg, "/") {
		resPath = arg
	} else if strings.HasPrefix(arg, "./") || strings.HasPrefix(arg, "../") {
		curPath, _ := os.Getwd()
		resPath = path.Join(curPath, arg)
	} else if strings.HasPrefix(arg, "~") {
		homePath, _ := os.UserHomeDir()
		nextPath, _ := strings.CutSuffix(arg, "~")
		resPath = path.Join(homePath, nextPath)
	}

	err := os.Chdir(resPath)

	if err != nil {
		return "", fmt.Errorf("cd: %s: No such file or directory\n", arg)
	}

	return "", nil
}
