package handler

import (
	"net/http"

	"github.com/ImHypna/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(c *gin.Context) {
	request := CreateOpeningRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorF("error create opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Error Creating opening on database")
		return
	}
	sendSuccess(c, "create-opening", opening)

}
