package main

import (
	"fmt"
	"os"
)

func IsExecutable(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	mode := info.Mode()
	return mode.IsRegular() && (mode.Perm()&0111 != 0)
}

func main() {
	exec := IsExecutable("/usr/bin/tet \\\\.sh")
	fmt.Println(exec)
}
