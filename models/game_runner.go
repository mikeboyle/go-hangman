package models

import (
  "math/rand"
  "time"

  "github.com/mikeboyle/go-hangman/words"
)

// GameRunner holds configuration and runs the game
type GameRunner struct {
  MaxWrongGuesses int
  BlankChar string
}

func (gr *GameRunner) Play() {
  g := Game{
    State: NewGameState(randomWord(), gr.MaxWrongGuesses),
    Board: NewBoard(gr.BlankChar),
    GuessInput: NewGuessInput(),
  }

  for !g.IsOver() {
    guess, _ := g.GetGuess()
    g.HandleGuess(guess)
  }
}

func randomWord() string {
  rand.Seed(time.Now().Unix())
  return words.Words[rand.Intn(len(words.Words))]
}
