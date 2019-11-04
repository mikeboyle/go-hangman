package main

import (
  "flag"

  "github.com/mikeboyle/go-hangman/models"
)

func main() {
  guesses := flag.Int("guesses", 8, "Max number of wrong guesses allowed")
  blank := flag.String("blank", "_", "Character used to show a blank letter in the secret word")

  flag.Parse()

  gr := models.GameRunner{
    MaxWrongGuesses: *guesses,
    BlankChar: *blank,
  }

  gr.Play()
}
