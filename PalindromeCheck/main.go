package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	sentence := "A man, a plan, a canal, PanaMa!"
	if(palindromeCheck(sentence)){
		fmt.Println(sentence, "---> is a plaindrome")
	}else {
		fmt.Println(sentence, "---> is not a palindrome")
	}
}

func palindromeCheck(sentence string) bool {
	cleanedString := strings.ToLower(strings.Map(func(r rune) rune {
		if unicode.IsLetter(r){
			return r
		}
		return -1
	}, sentence))

	for i := 0; i < len(cleanedString) / 2; i++{
		if(cleanedString[i] != cleanedString[len(cleanedString) - i - 1]){
			return false
		}
	}

	return true
}