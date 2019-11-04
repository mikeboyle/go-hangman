package models

import (
  "fmt"
  "strings"
  "github.com/mikeboyle/go-hangman/utils"
)

type Board struct {
  blankChar string
}

func NewBoard(blankChar string) *Board {
  return &Board{ blankChar: blankChar,}
}

func (b Board) RenderKey(key string, correctGuesses []string) {
  rendered := ""
  for _, c := range key {
    letter := string(c)
    if utils.Contains(correctGuesses, string(c)) {
      rendered += strings.ToUpper(letter)
    } else {
      rendered += b.blankChar
    }
    rendered += " "
  }
  fmt.Println("Secret word so far:", rendered)
}

func (b Board) RenderGuessResult(guess string, correct, won, lost bool) {
  fmt.Printf("You guessed: %v\n", strings.ToUpper(guess))
  fmt.Printf("Correct: %v\n", correct)
}

func (b Board) RenderGuessStatus(wrongGuesses []string, remaining int) {
  fmt.Printf("Wrong guesses: %v\n", wrongGuesses)
  fmt.Printf("Guesses left: %v\n\n", remaining)
}

func (b Board) RenderWonLost(key string, isWon, isLost bool) {
  if isWon {
    fmt.Println("You won!")
  }

  if isLost {
    fmt.Println("You lost. :(")
    fmt.Printf("The secret word was %v.\n", strings.ToUpper(key))
  }
}
