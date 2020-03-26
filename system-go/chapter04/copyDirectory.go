package chapter04

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WalkFunction(permissions os.FileMode, Path, NewPath string, flagMap map[string]*bool) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		fileInfo, _ := os.Lstat(path)
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			fmt.Println("Skipping : ", path)
			return nil
		}

		fileInfo, err = os.Stat(path)
		if err != nil {
			fmt.Println("Error :", err)
			return err
		}

		mode := fileInfo.Mode()
		if mode.IsDir() {
			tempPath := strings.Replace(path, Path, "", 1)
			pathToCreate := NewPath + "/" + filepath.Base(path) + tempPath
			_, err := os.Stat(pathToCreate)
			if os.IsNotExist(err) {
				os.MkdirAll(pathToCreate, permissions)
			} else {
				fmt.Println("Did not create :", pathToCreate, ":", err)
			}
		}

		return nil
	}
}

func copy() {
	minusTest := flag.Bool("test", false, "Test run!!")
	flag.Parse()
	flags := flag.Args()

	if len(flags) < 2 {
		fmt.Println("not enough arguments!!!")
		os.Exit(1)
	}

	Path := flags[0]
	NewPath := flags[1]
	flagMap := make(map[string]*bool)
	flagMap["test"] = minusTest

	fmt.Println("minusTest :", *minusTest)
	fmt.Println("PATH :", Path)
	fmt.Println("NewPATH :", NewPath)
	permissions := os.ModePerm
	_, err := os.Stat(NewPath)
	 if os.IsNotExist(err) {
	 	os.MkdirAll(NewPath, permissions)
	 } else {
	 	fmt.Println(NewPath, "already exists!!!")
	 	os.Exit(1)
	 }

	 err = filepath.Walk(Path, WalkFunction(permissions, Path, NewPath, flagMap))
	 if err != nil {
		 fmt.Println(err)
		 os.Exit(1)
	 }
}

func Start08() {
	copy()
}
