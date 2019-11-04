package models

import (
  "fmt"
  "strings"
  "github.com/mikeboyle/go-hangman/utils"
)
// GameState handles guesses and the state of the game
type GameState struct {
  key string
  maxWrongGuesses int
  wrongGuesses []string
  correctGuesses []string
}

func NewGameState (key string, maxWrongGuesses int) *GameState {
  return &GameState{
    key: key,
    maxWrongGuesses: maxWrongGuesses,
    wrongGuesses: []string{},
    correctGuesses: []string{},
  }
}

func (g *GameState) Key() string {
  return g.key
}

func (g *GameState) MaxWrongGuesses() int {
  return g.maxWrongGuesses
}

func (g *GameState) WrongGuesses() []string {
  return g.wrongGuesses
}

func (g *GameState) CorrectGuesses() []string {
  return g.correctGuesses
}

// IsLost checks if there are no more Guesses
func (g *GameState) IsLost() bool {
  return len(g.wrongGuesses) >= g.maxWrongGuesses
}

// IsWon checks if the user has guessed the word
func (g *GameState) IsWon() bool {
  for _, c := range g.key {
    if !utils.Contains(g.correctGuesses, string(c)) {
      return false
    }
  }
  return true
}

// AddGuess adds a guess to the array and updates the game status.
// Returns true if guess was correct, else false
func (g *GameState) AddGuess(guess string) (bool, error) {
  if g.IsLost() {
    return false, fmt.Errorf("Cannot exceed max %v Guesses", g.maxWrongGuesses)
  }

  if len(guess) != 1 {
    return false, fmt.Errorf("Guess must be one letter only, got %v", guess)
  }

  if utils.Contains(g.correctGuesses, guess) || utils.Contains(g.wrongGuesses, guess) {
    return false, fmt.Errorf("%v has already been guessed", guess)
  }

  isCorrect := g.handleAddGuess(strings.ToUpper(guess))
  return isCorrect, nil
}

func (g *GameState) handleAddGuess(guess string) bool {
  if stringContains(g.key, guess) {
    g.correctGuesses = append(g.correctGuesses, guess)
    return true
  }

  g.wrongGuesses = append(g.wrongGuesses, guess)
  return false
}

func stringContains(str, item string) bool {
  for _, c := range str {
    if strings.ToUpper(string(c)) == strings.ToUpper(item) {
      return true
    }
  }
  return false
}
