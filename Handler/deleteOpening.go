package handler

import (
	"fmt"
	"net/http"

	"github.com/ImHypna/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	//find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusFound, fmt.Sprintf("opening with id:  %s not found", id))
		return
	}
	//Delete
	if err := db.Delete(&opening).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
	}
	sendSuccess(c, "delete-opening", opening)
}
