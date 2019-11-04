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
  RenderGuessResult(guess string, correct, won, lost bool)
  RenderGuessStatus([]string, int)
  RenderWonLost(key string, isWon, isLost bool)
}

type GuessInputInterface interface {
  GetGuess() (string, error)
}

type Game struct {
  State GameStateInterface
  Board BoardInterface
  GuessInput GuessInputInterface
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

  g.renderGuessResult(guess, isCorrect)
  g.renderKey()
  g.renderWonLost()
  g.renderGuessStatus()
  return nil
}

func (g *Game) renderKey() {
  gs := g.State
  g.Board.RenderKey(gs.Key(), gs.CorrectGuesses())
}

func (g *Game) renderGuessResult(guess string, correct bool) {
  gs := g.State
  g.Board.RenderGuessResult(guess, correct, gs.IsWon(), gs.IsLost())
}

func (g *Game) renderGuessStatus() {
  gs := g.State
  wrongGuesses := gs.WrongGuesses()
  remaining := gs.MaxWrongGuesses() - len(wrongGuesses)
  g.Board.RenderGuessStatus(wrongGuesses, remaining)
}

func (g *Game) renderWonLost() {
  gs := g.State
  g.Board.RenderWonLost(gs.Key(), gs.IsWon(), gs.IsLost())
}
