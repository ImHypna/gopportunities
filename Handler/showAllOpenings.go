package handler

import (
	"net/http"

	"github.com/ImHypna/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

func ShowAllOpeningsHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(c, "list-openings", openings)
}
