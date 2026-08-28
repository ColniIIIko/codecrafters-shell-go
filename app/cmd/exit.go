package cmd

import "os"

func Exit(args []string, ctx any) (string, error) {
	os.Exit(0)
	return "", nil
}
