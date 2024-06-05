package scripts

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
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
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Cannot get current file info")
	}

	root := filepath.Dir(filename)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
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
