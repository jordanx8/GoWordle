package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

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
		//fmt.Printf("Found %s at index %d in the %s.\n", guess, index, "wordlist")
		return true
	} else {
		//fmt.Printf("%s is not in the word list.\nTry again: ", guess)
		return false
	}
}

func amtOfALetter(answer []byte, guess byte) int {
	amt := 0
	for _, letter := range answer {
		if guess == letter {
			amt++
		}
	}
	return amt
}

func checkWord(guess string, answer string) [5]int {
	var guessColors [5]int
	guessbyte := []byte(guess)
	answerbyte := []byte(answer)
	notinword := 0
	wrongspot := 1
	correct := 2
	for i := 0; i < 5; i++ {
		if string(guessbyte[i]) == string(answerbyte[i]) {
			//correct spot, correct letter
			guessColors[i] = correct
			answerbyte[i] = byte(0)
		}
	}
	for i := 0; i < 5; i++ {
		if guessColors[i] != correct {
			if amtOfALetter(answerbyte, guessbyte[i]) > 0 {
				//wrong spot, correct letter
				guessColors[i] = wrongspot
				answerbyte[i] = byte(0)
			} else {
				//wrong
				guessColors[i] = notinword
			}
		}
	}
	return guessColors
}

func main() {
	var guesses [6]string
	var lettercolors [6][5]int
	//store word list into an array
	wordlist, err := readLines("words.txt")
	sort.Strings(wordlist)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	rand.Seed(time.Now().UnixNano())

	answer := wordlist[rand.Intn(len(wordlist))]

	scanner := bufio.NewScanner(os.Stdin)

out:
	for numberofguesses := 0; numberofguesses < 6; numberofguesses++ {
		//fmt.Println("Answer =", answer)
		fmt.Println("Make a guess = ")
		scanner.Scan()
		guess := strings.ToLower(scanner.Text())
		for !isGuessInWordList(guess, wordlist) {
			fmt.Printf("%s is not in the word list. \n", guess)
			fmt.Print("Try again = ")
			scanner.Scan()
			guess = strings.ToLower(scanner.Text())
		}
		guesses[numberofguesses] = guess

		lettercolors[numberofguesses] = checkWord(guess, answer)
		fmt.Println()

		//only works on windows; clears terminal
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		for i := 0; i <= numberofguesses; i++ {
			correct := 0
			for j := 0; j < 5; j++ {
				if lettercolors[i][j] == 0 {
					fmt.Print(color.Colorize(color.Red, string([]byte(guesses[i])[j])))
				} else if lettercolors[i][j] == 1 {
					fmt.Print(color.Colorize(color.Yellow, string([]byte(guesses[i])[j])))
				} else {
					fmt.Print(color.Colorize(color.Green, string([]byte(guesses[i])[j])))
					correct++
					if correct == 5 {
						fmt.Println()
						fmt.Println("Correct! It only took you", numberofguesses+1, "guesses.")
						break out
					}
				}
			}
			fmt.Println()
		}
	}

}
