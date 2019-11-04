package models

import "fmt"

type GameStateInterface interface {
  Key() string
  IsWon() bool
  IsLost() bool
  MaxWrongGuesses() int
  WrongGuesses() []string
  CorrectGuesses() []string
  AddGuess(string) (bool, error)
}

type BoardInterface interface {
  RenderKey(string, []string)
  RenderGuessStatus([]string, int)
}

type GuessInputInterface interface {
  GetGuess() (string, error)
}

type Game struct {
  State GameStateInterface
  Board BoardInterface
  GuessInput GuessInputInterface
}

func (g *Game) RenderKey() {
  gs := g.State
  g.Board.RenderKey(gs.Key(), gs.CorrectGuesses())
}

func (g *Game) RenderGuessStatus() {
  gs := g.State
  wrongGuesses := gs.WrongGuesses()
  remaining := gs.MaxWrongGuesses() - len(wrongGuesses)
  g.Board.RenderGuessStatus(wrongGuesses, remaining)
}

func (g *Game) GetGuess() (string, error) {
  return g.GuessInput.GetGuess()
}

func (g *Game) IsOver() bool {
  gs := g.State
  return gs.IsWon() || gs.IsLost()
}

func (g *Game) HandleGuess(guess string) error {
  gs := g.State
  isCorrect, err := gs.AddGuess(guess)
  if err != nil {
    fmt.Println("Error adding guess:", err)
    return err
  }

  fmt.Printf("You guessed: %v\n", guess)
  fmt.Printf("Correct: %v\n", isCorrect)
  g.RenderKey()
  if gs.IsWon() {
    fmt.Println("You won!")
    return nil
  }

  if gs.IsLost() {
    fmt.Println("You lost...")
    fmt.Printf("The secret word was %v.\n", gs.Key())
    return nil
  }

  g.RenderGuessStatus()
  return nil
}
