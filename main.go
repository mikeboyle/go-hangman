package main

import "github.com/mikeboyle/go-hangman/models"

func main() {
  g := models.Game{
    State: models.NewGameState("jackpot", 6),
    Board: models.Board{ BlankChar: "_"},
    GuessInput: models.NewGuessInput(),
  }

  for !g.IsOver() {
    guess, _ := g.GetGuess()
    g.HandleGuess(guess)
  }
}
