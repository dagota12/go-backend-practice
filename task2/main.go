package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var text string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a text to display the frequencies: ")

	text, _ = reader.ReadString('\n')
	wordFrequency := WordFrequency(text)

	fmt.Println("Word Frequency:")
	for word, count := range wordFrequency {
		fmt.Printf("%s: %d\n", word, count)
	}
	fmt.Print("Enter a text to check if it is a palindrome: ")
	text, _ = reader.ReadString('\n')
	isPalindrome := IsPalindrome(text)

	fmt.Println("Is palindrome:", isPalindrome)
}
func WordFrequency(s string) map[string]int {
	//splits the string s on every special character and spaces
	words := strings.Fields(strings.ToLower(stripPunctuation(s)))

	// Create a frequency map
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[stripPunctuation(word)]++
	}

	return frequency
}
func stripPunctuation(str string) string {
	// Remove punctuation from start and end of string
	return strings.Trim(str, "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
}

func IsPalindrome(s string) bool {

	s = stripPunctuation(strings.ToLower(s))
	i, j := 0, len(s)-1
	for ; i < j; i, j = i+1, j-1 {

		if s[i] != s[j] {
			return false
		}
	}
	return true
}
func IsAlphaNumeric(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsDigit(s)
}
