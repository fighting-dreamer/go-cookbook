package chapter04

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

/*
	path is the complete path of file
	info is os.FileInfo of this file
	err is nil

	we call os.Stat() again coz many things may happen in-between info is created by filepath.Walk() and walkfunc() being called and executed
*/
func WalkFunc(path string, info os.FileInfo, err error) error {
	_, err = os.Stat(path)
	if err != nil {
		return err
	}
	fmt.Println(path)
	return nil
}

func WalkOnlyDir(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	mode := fileInfo.Mode()
	if mode.IsDir() {
		fmt.Println(path)
	}
	return nil
}

func WalkRegularOrDirFunc(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	mode := fileInfo.Mode()
	if mode.IsDir() || mode.IsRegular() {
		fmt.Println(path)
	}
	return nil
}

func SimpleWalk() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Not enoufh arguments!!!")
		os.Exit(1)
	}

	path := args[1]
	err := filepath.Walk(path, WalkFunc)
	if err != nil {
		fmt.Println("go error :", err)
		os.Exit(1)
	}
}

func WalkWithFlags(flagMap map[string]*bool) filepath.WalkFunc {
	minusS := flagMap["s"]
	minusP := flagMap["p"]
	minusSL := flagMap["sl"]
	minusD := flagMap["d"]
	minusF := flagMap["f"]

	printAll := false
	// when all flags are in arguments
	if *minusS && *minusP && *minusSL && *minusD && *minusF {
		printAll = true
	}
	// when no flags is in arguments
	if !(*minusS || *minusP || *minusSL || *minusD || *minusF) {
		printAll = true
	}
	return func(path string, info os.FileInfo, err error) error {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}

		if printAll {
			fmt.Println(path)
			return nil
		}
		mode := fileInfo.Mode()
		if *minusF && mode.IsRegular() {
			fmt.Println(path)
			return nil
		}

		if *minusD && mode.IsDir() {
			fmt.Println(path)
			return nil
		}

		if *minusSL && (mode&os.ModeSymlink != 0) {
			fmt.Println(path)
			return nil
		}

		if *minusP && (mode&os.ModeNamedPipe != 0) {
			fmt.Println(path)
			return nil
		}

		if *minusS && (mode&os.ModeSocket != 0) {
			fmt.Println(path)
			return nil
		}
		return nil
	}
}

func FindWithFlags() {
	flagMap := make(map[string]*bool)
	minusS := flag.Bool("s", false, "Sockets")
	minusP := flag.Bool("p", false, "Pipes")
	minusSL := flag.Bool("sl", false, "Symbolic Links")
	minusD := flag.Bool("d", false, "Directories")
	minusF := flag.Bool("f", false, "files")

	flag.Parse()
	flagMap["s"] = minusS
	flagMap["p"] = minusP
	flagMap["sl"] = minusSL
	flagMap["d"] = minusD
	flagMap["f"] = minusF
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Println("not enough Arguments!!!")
		os.Exit(1)
	}

	path := flags[0]
	err := filepath.Walk(path, WalkWithFlags(flagMap))

	if err != nil {
		fmt.Println("go error :", err)
		os.Exit(1)
	}
}


func excludeNames(path, excluded string) bool {
	if excluded == "" {
		return false
	}

	if filepath.Base(path) == excluded {
		fmt.Println(path, ":", excluded)
		return true
	}

	return false
}

func WalkWithFlagsWithExclude(flagMap map[string]interface{}) filepath.WalkFunc {
	minusS := flagMap["s"].(*bool)
	minusP := flagMap["p"].(*bool)
	minusSL := flagMap["sl"].(*bool)
	minusD := flagMap["d"].(*bool)
	minusF := flagMap["f"].(*bool)
	minusX := flagMap["x"].(*string)

	printAll := false
	// when all flags are in arguments
	if *minusS && *minusP && *minusSL && *minusD && *minusF {
		printAll = true
	}
	// when no flags is in arguments
	if !(*minusS || *minusP || *minusSL || *minusD || *minusF) {
		printAll = true
	}
	return func(path string, info os.FileInfo, err error) error {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}

		if excludeNames(path, *minusX) {
			return nil
		}

		if printAll {
			fmt.Println(path)
			return nil
		}
		mode := fileInfo.Mode()
		if *minusF && mode.IsRegular() {
			fmt.Println(path)
			return nil
		}

		if *minusD && mode.IsDir() {
			fmt.Println(path)
			return nil
		}

		if *minusSL && (mode&os.ModeSymlink != 0) {
			fmt.Println(path)
			return nil
		}

		if *minusP && (mode&os.ModeNamedPipe != 0) {
			fmt.Println(path)
			return nil
		}

		if *minusS && (mode&os.ModeSocket != 0) {
			fmt.Println(path)
			return nil
		}
		return nil
	}
}

func FindWithFlagsWithExclude() {
	flagMap := make(map[string]interface{})
	minusS := flag.Bool("s", false, "Sockets")
	minusP := flag.Bool("p", false, "Pipes")
	minusSL := flag.Bool("sl", false, "Symbolic Links")
	minusD := flag.Bool("d", false, "Directories")
	minusF := flag.Bool("f", false, "files")
	minusX := flag.String("x", "", "Files")

	flag.Parse()
	flagMap["s"] = minusS
	flagMap["p"] = minusP
	flagMap["sl"] = minusSL
	flagMap["d"] = minusD
	flagMap["f"] = minusF
	flagMap["x"] = minusX

	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Println("not enough Arguments!!!")
		os.Exit(1)
	}

	path := flags[0]
	err := filepath.Walk(path, WalkWithFlagsWithExclude(flagMap))

	if err != nil {
		fmt.Println("go error :", err)
		os.Exit(1)
	}
}

func Start07() {
	FindWithFlagsWithExclude()
}
