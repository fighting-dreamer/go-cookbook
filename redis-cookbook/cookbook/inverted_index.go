package cookbook

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"io/ioutil"
	"strings"
)

const newLineSeperator = "\n"
const wordSeperator = " "
const stopWordsListString = "the,of,to,and,a,in,is,it,you,that"

func readFile(fileName string) []string {
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []string{}
	}
	text := string(fileContent)
	return strings.Split(text, newLineSeperator)
}

func CreateInvertedIndexes(redisCli *redis.Client) {
	fileNames := getFileNames()
	stopWords := strings.Split(stopWordsListString, ",")
	stopWordsMap := make(map[string]bool)
	for _, stopWord := range stopWords {
		stopWordsMap[stopWord] = true
	}

	for _, fileName := range fileNames {
		fileContent := readFile(fileName)
		for _, line := range fileContent {
			words := strings.Split(line, wordSeperator)
			for _, word := range words {
				if stopWordsMap[word] == false {
					redisCli.SAdd(word, fileName)
				}
			}
		}
	}
}

func GetMostProbableFiles(redisCli *redis.Client, search string) []string {
	searchWords := strings.Split(search, wordSeperator)
	result, err := redisCli.SInter(searchWords...).Result()
	if err != nil {
		fmt.Println("couldn't find the files!!!")
		return []string{}
	}
	return result
}

func InvertedIndexesExample(redisCli *redis.Client) {
	CreateInvertedIndexes(redisCli)
	GetMostProbableFiles(redisCli, "just testing")
}