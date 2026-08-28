package cmd

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

func Pwd(args []string, ctx utils.Shell) (string, error) {
	pwd, err := os.Getwd()

	if err != nil {
		return "", fmt.Errorf("pwd error: %s\n", err)
	}

	return pwd, nil
}
