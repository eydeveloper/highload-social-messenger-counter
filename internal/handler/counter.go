package handler

import (
	"github.com/eydeveloper/highload-social-messenger-counter/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) incrementCounter(c *gin.Context) {
	requestId := c.MustGet("X-Request-ID").(string)

	logrus.Info("Handling increase messages counter request with ID: " + requestId)

	var message entity.Message

	if err := c.BindJSON(&message); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Counter.Increment(message)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"result": true})
}
