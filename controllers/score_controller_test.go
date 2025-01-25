package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"open-highscore/models"
	"open-highscore/services"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetScores(t *testing.T) {
	services.AddScore(*models.NewScore("TestGame", "Alice", 500))
	services.AddScore(*models.NewScore("TestGame", "Bob", 400))

	r := gin.Default()

	r.GET("/scores", GetScores)

	req, _ := http.NewRequest("GET", "/scores", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string][]models.Score
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to parse JSON response: %v", err)
	}

	if len(response["scores"]) != 2 {
		t.Errorf("Expected 2 scores, got %d", len(response["scores"]))
	}

	if response["scores"][0].Player != "Alice" || response["scores"][0].Points != 500 {
		t.Errorf("Expected Alice with 500 points, got %v", response["scores"][0])
	}

	if response["scores"][1].Player != "Bob" || response["scores"][1].Points != 400 {
		t.Errorf("Expected Bob with 400 points, got %v", response["scores"][1])
	}
}

func TestAddScore(t *testing.T) {
	r := gin.Default()

	r.POST("/scores", AddScore)

	jsonStr := `{"game": "TestGame", "player": "Charlie", "points": 300}`

	req, _ := http.NewRequest("POST", "/scores", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to parse JSON response: %v", err)
	}

	if response["message"] != "score added successfully" {
		t.Errorf("Expected success message 'Score added successfully', got '%s'", response["message"])
	}
}
