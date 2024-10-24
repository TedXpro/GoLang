package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	sentence := "aaa bbb bbb aaa! bbb. ccc? ddd aaa aaa   bbb"
	freq := map[string]int{};

	countWords(sentence, freq)

	fmt.Println(sentence)
	for key, val := range freq{
		fmt.Println(key, "------> has", val, "number of occurences in the sentence")
	}
}

func countWords(sentence string, freq map[string]int) {
	currWord := ""
	for index := 0; index < len(sentence); index++ {
		currChar := rune(sentence[index])
		
		if currChar == ' ' && currWord != "" {
			currWord = strings.TrimSpace(strings.ToLower(currWord))
			freq[currWord]++
			currWord = ""
		} else if (currChar != ' ' && !unicode.IsPunct(currChar)) {
			currWord += string(currChar)
		} else if(unicode.IsPunct(currChar) && currWord != ""){
			currWord = strings.TrimSpace(strings.ToLower(currWord))
			freq[currWord]++
			currWord = ""
		}
	}

	currWord = strings.TrimSpace(strings.ToLower(currWord))
	if currWord != "" {
		freq[currWord]++

		
	}
}