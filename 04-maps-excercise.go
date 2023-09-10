package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	text := "leon is a good man and his name is leon"
	returnedWorkCountMap := WordCountFunction(text)

	fmt.Println("Word Counts : ")
	for word, count := range returnedWorkCountMap {
		fmt.Println(word, count)
	}
}

// this function will take a string input : s
// return a map with key of type string
// value of type int
func WordCountFunction(s string) map[string]int {

	//split the string based on space characters
	words := strings.Fields(s)
	log.Println(words) //[leon is a good man and his name is leon]

	wordCountMap := make(map[string]int) //declares and initialize map
	log.Println(wordCountMap)            //map[]

	for _, word := range words {
		wordCountMap[word]++
		log.Println(wordCountMap) //final iternation would be --> map[a:1 and:1 good:1 his:1 is:2 leon:2 man:1 name:1]
	}
	return wordCountMap
}
