package cmd

import (
	"fmt"
	"strings"
)

func Echo(args []string, ctx any) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}
