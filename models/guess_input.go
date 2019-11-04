package models

import (
  "bufio"
  "fmt"
  "os"
)

type GuessInput struct {
  Scanner *bufio.Scanner
}

func NewGuessInput() *GuessInput {
  return &GuessInput{
    Scanner: bufio.NewScanner(os.Stdin),
  }
}

func (i *GuessInput) GetGuess() (string, error) {
  fmt.Println("Guess a letter in the secret word:")
  i.Scanner.Scan()
  text := i.Scanner.Text()
  return text, nil
}
