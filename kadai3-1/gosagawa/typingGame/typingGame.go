package typingGame

import "fmt"

func Start() {
	fmt.Println("start")

	word := chooseWord()
	displayWord(word)
	input := inputWord()
	isCorrect := check(word, input)
	displayResult(isCorrect)
}

func chooseWord() string {
	return "test"
}

func displayWord(word string) {
	fmt.Println(word)
}

func inputWord() string {
	var stdin string
	fmt.Scan(&stdin)
	return stdin
}

func check(word string, input string) bool {
	if word == input {
		return true
	}
	return false
}

func displayResult(isCorrect bool) {
	if isCorrect {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Miss...")
	}
}
