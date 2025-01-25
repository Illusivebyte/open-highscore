package models

import "time"

// Score represents a player's high score.
//
// It includes the player's name, their score (points), and the timestamp when the score was recorded.
type Score struct {
	Game      string    `json:"game" binding:"required"`
	Player    string    `json:"player" binding:"required"`
	Points    int       `json:"points" binding:"required"`
	Timestamp time.Time `json:"timestamp"`
}

// NewScore creates and returns a new Score instance with the current timestamp.
//
// It initializes a Score with the provided player name and points and sets the timestamp to the current time.
func NewScore(game string, player string, points int) *Score {
	return &Score{
		Game:      game,
		Player:    player,
		Points:    points,
		Timestamp: time.Now(),
	}
}
