package scripts

import (
	"os"
	"path/filepath"
)

func GetAbsPathDir() string {
	execPath, err := os.Executable()
	if err != nil {
		panic("Cannot get executable path: " + err.Error() + "\n")
	}
	absExecPath, err := filepath.Abs(execPath)
	if err != nil {
		panic("Cannot get absolute path: " + err.Error() + "\n")
	}
	return filepath.Dir(absExecPath)
}
