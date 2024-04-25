package getProduct

import (
	"github.com/gin-gonic/gin"
)

type Servicer interface {
	Execute(c *gin.Context) (*Result, error)
}

type handler struct {
	service Servicer
}

func NewHandler(servicer Servicer) *handler {
	return &handler{service: servicer}
}

func (h handler) GetProductById(c *gin.Context) {
	res, err := h.service.Execute(c)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	c.JSON(200, res)
}
