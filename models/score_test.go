package models

import (
	"testing"
)

func TestNewScore(t *testing.T) {
	score := NewScore("TestGame", "Alice", 500)

	if score.Game != "TestGame" {
		t.Errorf("Expected game 'TestGame', got '%s'", score.Game)
	}

	if score.Player != "Alice" {
		t.Errorf("Expected player 'Alice', got '%s'", score.Player)
	}

	if score.Points != 500 {
		t.Errorf("Expected 500 points, got %d", score.Points)
	}

	if score.Timestamp.IsZero() {
		t.Error("Expected timestamp to be set, got zero time")
	}
}
