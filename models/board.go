package models

import (
  "fmt"
  "strings"
  "github.com/mikeboyle/go-hangman/utils"
)

type Board struct {
  BlankChar string
}

func (b Board) RenderKey(key string, correctGuesses []string) {
  rendered := ""
  for _, c := range key {
    letter := string(c)
    if utils.Contains(correctGuesses, string(c)) {
      rendered += strings.ToUpper(letter)
    } else {
      rendered += b.BlankChar
    }
    rendered += " "
  }
  fmt.Println("Secret word so far:", rendered)
}

func (b Board) RenderGuessStatus(wrongGuesses []string, remaining int) {
  fmt.Printf("Wrong guesses: %v\n", wrongGuesses)
  fmt.Printf("Guesses left: %v\n\n", remaining)
}
