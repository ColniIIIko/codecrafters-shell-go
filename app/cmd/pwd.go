package cmd

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

func Pwd(args []string, ctx utils.Shell) error {
	pwd, err := os.Getwd()

	if err != nil {
		fmt.Printf("pwd error: %s\n", err)
		return err
	}

	fmt.Printf("%s\n", pwd)

	return nil
}
