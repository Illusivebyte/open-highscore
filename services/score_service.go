package services

import (
	"open-highscore/models"
	"sync"
)

var scores = []models.Score{}
var mu sync.Mutex

// GetAllScores retrieves all high scores from the in-memory database.
//
// This function returns a slice of all the scores stored in the system.
func GetAllScores() []models.Score {
	mu.Lock()
	defer mu.Unlock()
	return scores
}

// AddScore adds a new high score to the in-memory database.
//
// The function takes a Score struct and appends it to the global scores list.
func AddScore(score models.Score) {
	mu.Lock()
	defer mu.Unlock()
	scores = append(scores, score)
}
