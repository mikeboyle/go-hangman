package main

import "github.com/mikeboyle/go-hangman/models"

func main() {
  gs, _ := models.NewGameState("jackpot", 6)
  i := models.NewGuessInput()
  g := models.Game{
    State: gs,
    Board: models.Board{ BlankChar: "_"},
    GuessInput: i,
  }

  for !g.IsOver() {
    guess, _ := g.GetGuess()
    g.HandleGuess(guess)
  }
}
