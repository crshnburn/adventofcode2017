package main

import (
	"fmt"
	"sort"
	"strings"
	"os"
	"bufio"
	"log"
)

func CheckPassPhrase(passphrase string) bool {
	words := strings.Split(passphrase, " ")
	sort.Strings(words)
	for i, v := range(words) {
		if i < len(words)-1 && v == words[i+1] {
			return false
		}
	}
	return true
}

func CheckPassPhraseAnagram(passphrase string) bool {
	words := strings.Split(passphrase, " ")
	for i, v := range(words) {
		chars := strings.Split(v, "")
		sort.Strings(chars)
		words[i] = strings.Join(chars, "")
	}
	return CheckPassPhrase(strings.Join(words, " "))
}

func main() {
	file, err := os.Open("./passphrases.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	successfulPhrases := 0
    for scanner.Scan() {
        if CheckPassPhraseAnagram(scanner.Text()){
			successfulPhrases++
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	fmt.Println("Number of successful pass phrases ", successfulPhrases)
}