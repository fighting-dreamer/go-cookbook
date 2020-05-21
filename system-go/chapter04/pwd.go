package chapter04

import (
	"fmt"
	"os"
	"path/filepath"
)

func Start03() {
	args := os.Args
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	if len(args) == 1 {
		fmt.Println(pwd)
		return
	}
	if args[1] == "-P" {
		fileInfo, err := os.Lstat(pwd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if fileInfo.Mode()&os.ModeSymlink != 0 {
			realPath, err := filepath.EvalSymlinks(pwd)
			if err == nil {
				fmt.Println(realPath)
			}
		}
	}
}
