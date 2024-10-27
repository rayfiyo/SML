package handles

import (
	"log"
	"math/rand/v2"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayfiyo/simple-language-model/models"
)

func Message(c *gin.Context) {
	var msg models.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response string

	switch msg.Mode {
	case "query":
		if !contains(models.Storage.Query, msg.Content) {
			models.Storage.Query = append(models.Storage.Query, msg.Content)
			log.Printf("Added new query: %s", msg.Content)
		} else {
			log.Printf("Duplicate query ignored: %s", msg.Content)
		}
		if len(models.Storage.Reply) > 0 {
			response = models.Storage.Reply[rand.N(len(models.Storage.Reply))]
		}
	case "reply":
		if !contains(models.Storage.Reply, msg.Content) {
			models.Storage.Reply = append(models.Storage.Reply, msg.Content)
			log.Printf("Added new reply: %s", msg.Content)
		} else {
			log.Printf("Duplicate reply ignored: %s", msg.Content)
		}
		if len(models.Storage.Query) > 0 {
			response = models.Storage.Query[rand.N(len(models.Storage.Query))]
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mode"})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		Mode:    msg.Mode,
		Content: response,
	})
}

// 重複チェック用の関数
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
