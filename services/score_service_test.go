package services

import (
	"open-highscore/models"
	"testing"
)

func TestAddScore(t *testing.T) {
	scores = []models.Score{}

	newScore := models.NewScore("TestGame", "Alice", 500)
	AddScore(*newScore)

	if len(scores) != 1 {
		t.Errorf("Expected 1 score, got %d", len(scores))
	}

	if scores[0].Game != "TestGame" {
		t.Errorf("Expected game 'TestGame', got %s", scores[0].Game)
	}

	if scores[0].Player != "Alice" {
		t.Errorf("Expected player 'Alice', got %s", scores[0].Player)
	}

	if scores[0].Points != 500 {
		t.Errorf("Expected 500 points, got %d", scores[0].Points)
	}

	if scores[0].Timestamp.IsZero() {
		t.Error("Expected timestamp to be set, got zero time")
	}
}

func TestGetAllScores(t *testing.T) {
	scores = []models.Score{}

	AddScore(*models.NewScore("TestGame", "Alice", 500))
	AddScore(*models.NewScore("TestGame", "Bob", 400))

	retrievedScores := GetAllScores()

	if len(retrievedScores) != 2 {
		t.Errorf("Expected 2 scores, got %d", len(retrievedScores))
	}

	if retrievedScores[0].Player != "Alice" || retrievedScores[0].Points != 500 {
		t.Errorf("Expected Alice with 500 points, got %v", retrievedScores[0])
	}

	if retrievedScores[1].Player != "Bob" || retrievedScores[1].Points != 400 {
		t.Errorf("Expected Bob with 400 points, got %v", retrievedScores[1])
	}
}
