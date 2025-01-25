package controllers

import (
	"net/http"
	"open-highscore/models"
	"open-highscore/services"

	"github.com/gin-gonic/gin"
)

// GetScores retrieves a list of all high scores.
//
// It calls the service layer to fetch all the scores and returns them in a JSON format.
// The response includes a "scores" key with the list of scores.
func GetScores(c *gin.Context) {
	scores := services.GetAllScores()

	c.IndentedJSON(http.StatusOK, gin.H{"scores": scores})
}

// AddScore adds a new high score to the system.
//
// The function expects a JSON request body with the player's name and score.
// If the request is valid, the score is added to the system. Otherwise, a BadRequest error is returned.
func AddScore(c *gin.Context) {
	var newScore models.Score
	if err := c.ShouldBindJSON(&newScore); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.AddScore(newScore)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "score added successfully"})
}
