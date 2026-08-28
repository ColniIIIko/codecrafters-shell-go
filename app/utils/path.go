package utils

import (
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	PATH_ENV = "PATH"
)

func ExecutablePath(executable string) (string, error) {
	if IsExecutable(executable) {
		return executable, nil
	}

	pathEnvValue, exists := os.LookupEnv(PATH_ENV)

	if !exists {
		return "", fmt.Errorf("PATH env not found")
	}

	for pathValue := range strings.SplitSeq(pathEnvValue, string(os.PathListSeparator)) {
		fullPath := path.Join(pathValue, executable)

		if IsExecutable(fullPath) {
			return fullPath, nil
		}
	}

	return "", fmt.Errorf("%s: not found", executable)
}

func IsExecutable(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	mode := info.Mode()
	return mode.IsRegular() && (mode.Perm()&0111 != 0)
}

func ResolvePath(p string) string {
	resPath := ""

	if strings.HasPrefix(p, "/") {
		resPath = p
	} else if strings.HasPrefix(p, "./") || strings.HasPrefix(p, "../") {
		curPath, _ := os.Getwd()
		resPath = path.Join(curPath, p)
	} else if strings.HasPrefix(p, "~") {
		homePath, _ := os.UserHomeDir()
		nextPath, _ := strings.CutSuffix(p, "~")
		resPath = path.Join(homePath, nextPath)
	}

	return resPath
}
