package chapter05

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func increment(value, inc int64) int64 {
	return value + inc
}

func countWords(line string) int64 {
	wordsPatterns := "[^\\s]+"
	r := regexp.MustCompile(wordsPatterns)
	countWords := int64(0)
	for range r.FindAllString(line, -1) {
		countWords++
	}
	return countWords
}

func scanLineBasedWC(){
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Not Enough Arguments!!!")
		os.Exit(1)
	}
	numberOfLines := int64(0)
	numberOfWords := int64(0)
	numberOfCharacters := int64(0)

	fileName := args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	lineDelimiter := byte('\n') // it needs to be a character!!!
	for {
		line, err := reader.ReadString(lineDelimiter)
		if err == io.EOF {
			break
		}else if err != nil {
			fmt.Println("error :", err)
			os.Exit(1)
		}
		numberOfLines = increment(numberOfLines, 1)
		numberOfWords = increment(numberOfWords, countWords(line))
		numberOfCharacters = increment(numberOfCharacters, int64(len(line)))
	}

	fmt.Println("L :", numberOfLines, "W :", numberOfWords, "C :", numberOfCharacters)
}

// we will read whole file nad read the buffer one chracter at a time
func characterByCharacter() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Not Enough Arguments!!!")
		os.Exit(1)
	}
	fileName := args[1]
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(buf)))
	/*
	“Here, ScanRunes is a split function that returns each character (rune) as a token. Then, the call to Scan() allows us to process each character one by one. There also exist ScanWords and ScanLines for getting words and lines, respectively.”

	Excerpt From: Mihalis Tsoukalos. “Go Systems Programming”. Apple Books.
	*/
	scanner.Split(bufio.ScanRunes)
	numberLines := int64(0)
	numberWords := int64(0)
	numberCharacters := int64(0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "\n" {
			numberLines++
		}
		if text == " " || text == "\n" {
			numberWords++
		}
		numberCharacters++

	}

	fmt.Println("L :", numberLines, "W :", numberWords, "C :", numberCharacters)
}

func Start06() {
	scanLineBasedWC()
	characterByCharacter()
}
