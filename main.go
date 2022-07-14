package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)
var dictionary = []string{"Zombie", "Gopher", "United States of America", "Indonesia", "Nazism", "Apple", "Programming"}

func main() {
	rand.Seed(time.Now().UnixNano())
	targetWord := getRandomWord()
	guessedLetters := initializeGuessWords(targetWord)
	hangmanState := 0
	for !isGameOver(targetWord, guessedLetters, hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("Invalid input. Please use letters only...")
			continue
		}
		letter := rune(input[0])
		if isCorrectGuess(targetWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}
	}
	printGameState(targetWord, guessedLetters, hangmanState)
	fmt.Print("Game Over... ")
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You win!")
	} else if isHangmanComplete(hangmanState) {
		fmt.Println("You loose!")
	} else {
		panic("Invalid state. Game is over there is no winner!")
	}
}

func initializeGuessWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true
	return guessedLetters
}

func isGameOver(targetWord string, guessedLetters map[rune]bool, hangmanState int) bool {
	return isWordGuessed(targetWord, guessedLetters) || isHangmanComplete(hangmanState)
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}
	return true
}
func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return (targetWord)
}

func printGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	if hangmanState != 0 {
		fmt.Println()
	}
	fmt.Println(getHangmanDrawing(hangmanState))
}

func getWordGuessingProgress(targetWord string, guessedLetters map[rune]bool) string {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += (" ")
		} else if guessedLetters[unicode.ToLower(ch)] == true {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += ("_")
		}

	}
	return result
}

func getHangmanDrawing(state int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("states/hangman%d", state))
	if err != nil {
		panic(err)
	}
	return string(data)
}

func readInput() string {
	fmt.Print("> ")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(input)
}

func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}
