package scripts

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func replaceInFile(filePath, oldString, newString string) error {
	input, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	output := strings.ReplaceAll(string(input), oldString, newString)

	err = ioutil.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}

func RenameAllGoCodeFragment(from, to string) {
	execPath, err := os.Executable()
	if err != nil {
		fmt.Printf("Cannot get executable path: %v\n", err)
		return
	}

	absExecPath, err := filepath.Abs(execPath)
	if err != nil {
		fmt.Printf("Cannot get absolute path: %v\n", err)
		return
	}

	root := filepath.Dir(absExecPath)
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			err := replaceInFile(path, from, to)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", root, err)
	} else {
		fmt.Println("Replacement completed successfully.")
	}
}
