package cookbook

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"io/ioutil"
	util "nipun.io/tools"
	"strings"
)

const newLineSeperator = "\n"
const wordSeperator = " "
const stopWordsListString = ",the,of,to,and,a,in,is,it,you,that"

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
	wordCounter := 0
	for _, fileName := range fileNames {
		fileContent := readFile(fileName)
		for _, line := range fileContent {
			words := strings.Split(line, wordSeperator)
			for _, word := range words {
				if stopWordsMap[word] == false {
					redisCli.SAdd(word, fileName)
					wordCounter = wordCounter + 1
				}
			}
		}
	}
	fmt.Println("word Counter :", wordCounter)
}

func SearcFiles(redisCli *redis.Client, search string) []string {
	searchWords := strings.Split(search, wordSeperator)
	for _, word := range searchWords {
		fileNames := redisCli.SMembers(word)
		fmt.Println(word, ":", fileNames)
	}
	result, err := redisCli.SInter(searchWords...).Result()
	if err != nil {
		fmt.Println("couldn't find the files!!!")
		return []string{}
	}
	return result
}

func InvertedIndexesExample(redisCli *redis.Client) {
	CreateInvertedIndexes(redisCli)
	fileNames := SearcFiles(redisCli, "support our functionality")
	fmt.Println("files that most probably contains this string :", strings.Join(fileNames, ","))
}

func CreateInvertedIndexesWithScore(redisCli *redis.Client) {
	fileNames := getFileNames()
	stopWords := strings.Split(stopWordsListString, ",")
	stopWordsMap := make(map[string]bool)
	for _, stopWord := range stopWords {
		stopWordsMap[stopWord] = true
	}
	wordCounter := 0
	for _, fileName := range fileNames {
		fileContent := readFile(fileName)
		for _, line := range fileContent {
			words := strings.Split(line, wordSeperator)
			for _, word := range words {
				if stopWordsMap[word] == false {
					redisCli.ZIncrBy(word, 1, fileName)
					wordCounter = wordCounter + 1
				}
			}
		}
	}
	fmt.Println("word Counter :", wordCounter)
}

func SearcFilesWithScore(redisCli *redis.Client, search string) []string {
	searchWords := strings.Split(search, wordSeperator)
	searchId := util.GetUniqueID()
	redisCli.ZInterStore(searchId, &redis.ZStore{
			Keys:      searchWords,})

	result, err := redisCli.ZRevRange(searchId, 0, -1).Result()
	if err != nil {
		fmt.Println("couldn't find the files!!!")
		return []string{}
	}
	return result
}

func InvertedIndexesWithScoreExample(redisCli *redis.Client) {
	CreateInvertedIndexesWithScore(redisCli)
	fileNames := SearcFilesWithScore(redisCli, "channel names are not")
	fmt.Println("score wise files that most probably contains this string :", strings.Join(fileNames, ","))
}
