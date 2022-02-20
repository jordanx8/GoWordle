package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"

	"github.com/TwiN/go-color"
)

type guesses struct {
	guess1 string
	guess2 string
	guess3 string
	guess4 string
	guess5 string
	guess6 string
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func isGuessInWordList(guess string, wordlist []string) bool {
	index := sort.SearchStrings(wordlist, guess)
	if index < len(wordlist) && wordlist[index] == guess {
		fmt.Printf("Found %s at index %d in the %s.\n", guess, index, "wordlist")
		return true
	} else {
		fmt.Printf("%s is not in the word list.\nTry again: ", guess)
		return false
	}
}

func checkWord(guess string, answer string) []int {
	var guessColors [5]int
	guessbyte := []byte(guess)
	answerbyte := []byte(guess)
	wrong := 0
	wrongspot := 1
	correct := 2
	for i := 0; i < 5; i++ {
		//correct spot, correct letter
		if string(guessbyte[i]) == string(answerbyte[i]) {
			guessColors[i] = correct
			answerbyte[i] = nil
		}
		//wrong spot, correct letter
		else if {
			index := sort.SearchChars()
		}
	}
}

func main() {
	var guesses [6]string
	var letterColors [6][5]int
	//store word list into an array
	wordlist, err := readLines("words.txt")
	sort.Strings(wordlist)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	answer := wordlist[rand.Intn(len(wordlist))]

	scanner := bufio.NewScanner(os.Stdin)

	for numberofguesses := 0; numberofguesses < 6; numberofguesses++ {
		fmt.Println("Answer =", answer)
		scanner.Scan()
		guess := strings.ToLower(scanner.Text())
		for !isGuessInWordList(guess, wordlist) {
			scanner.Scan()
			guess = strings.ToLower(scanner.Text())
		}
		guesses[numberofguesses] = guess

		//checkWord()

		fmt.Println(color.Colorize(color.Red, string([]byte(guess)[numberofguesses])))
	}

}
