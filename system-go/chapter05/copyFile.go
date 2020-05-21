package chapter05

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func myCopy(source, destination string) {
	sourceFileInfo, err := os.Stat(source)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}
	sourceMode := sourceFileInfo.Mode()
	if !sourceMode.IsRegular() {
		fmt.Println("error :", err)
		os.Exit(1)
	}
	sourceFile , err := os.Open(source)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}
	defer destinationFile.Close()
	nBytes, err := io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Println("error while copying : ", err)
		os.Exit(1)
	}
	fmt.Printf("copied : \n %s", string(nBytes))
}

func SimpleCopy() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Not enough arguments!!!")
		os.Exit(1)
	}

	sourceFile := args[1]
	destfile := args[2]

	myCopy(sourceFile, destfile)
}

func readAndCopy(sourceFileName, destFileName string) {
	sourceFileInfo, err := os.Stat(sourceFileName)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}
	destFile, err := os.Create(destFileName)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}
	defer destFile.Close()

	sourceFileMode := sourceFileInfo.Mode()
	if !sourceFileMode.IsRegular() {
		fmt.Println("error :", err)
		os.Exit(1)
	}

	input, err := ioutil.ReadFile(sourceFileName)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(destFileName, input, 0644)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}
}

func readAllAndCopy() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Not enough arguments!!!")
		os.Exit(1)
	}

	sourceFile := args[1]
	destfile := args[2]

	readAndCopy(sourceFile, destfile)

}

// will be introducing a buffer size of fixed size
var  BUFFERSIZE int64

func propercopyUsingBuffer(sourceFileName, destFileName string, bufferSize int64) {
	sourceFileInfo, err := os.Stat(sourceFileName)

	if err != nil {
		fmt.Println("Not enough arguments!!!")
		os.Exit(1)
	}
	 sourceFileMode := sourceFileInfo.Mode()
	 if !sourceFileMode.IsRegular() {
		 fmt.Println("Not enough arguments!!!")
		 os.Exit(1)
	 }

	 sourceFile, err := os.Open(sourceFileName)
	 if err != nil {
		 fmt.Println("Not enough arguments!!!")
		 os.Exit(1)
	 }
	 defer sourceFile.Close()
	 destFile, err := os.Create(destFileName)
	 if err != nil {
		 fmt.Println("Not enough arguments!!!")
		 os.Exit(1)
	 }
	defer destFile.Close()

	 buf := make([]byte, bufferSize)
	 for {
	 	n, err := sourceFile.Read(buf)
	 	if err != nil && err != io.EOF {
			fmt.Println("Not enough arguments!!!")
			os.Exit(1)
		}
		if n == 0 {
			break
		}
		_, err = destFile.Write(buf[:n])
		if err != nil {
			fmt.Println("Not enough arguments!!!")
			os.Exit(1)
		}
	 }
}

func betterCopy() {
	args := os.Args
	if len(args) != 4 {
		fmt.Println("Not enough arguments!!!")
		os.Exit(1)
	}
	sourceFileName := args[1]
	destFileName := args[2]
	bufferSize := args[3]

	BUFFERSIZE, err := strconv.ParseInt(bufferSize, 10,64) // setting a value for buffer size
	if err != nil {
		fmt.Println("Not enough arguments!!!")
		os.Exit(1)
	}
	propercopyUsingBuffer(sourceFileName, destFileName, BUFFERSIZE)
}

func Start05() {
	//SimpleCopy() //using io.Copy function to copy between file
	//readAllAndCopy() // it reads all the file and then copies it, not useful for large size files like 5GB or more, depending on your system!!!
	betterCopy()
}
