package chapter04

import (
	"fmt"
	"os"
	"path/filepath"
)

func Start02() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("please provide an argument!")
		os.Exit(1)
	}

	filename := args[1]
	fileInfo, err := os.Lstat(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if fileInfo.Mode()*os.ModeSymlink != 0 {
		fmt.Println(filename, "is a symbolic link")
		realPath, err := filepath.EvalSymlinks(filename)
		if err == nil {
			fmt.Println(realPath)
		}
	}
}
