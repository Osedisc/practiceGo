package getProduct

import (
	"errors"
	"strconv"

	"github.com/Osedisc/practiceGo/provider"
	"github.com/gin-gonic/gin"
)

type service struct {
	provider provider.Product
}

func NewService(provider provider.Product) *service {
	return &service{provider: provider}
}

func (s *service) Execute(ctx *gin.Context) (*Result, error) {
	id, _ := ctx.Params.Get("id")
	res, err := provider.GetProductById(id, ctx)
	if err != nil {
		return nil, err
	}
	if strconv.Itoa(res.Id) != id {
		return nil, errors.New("product not found")
	}
	return &Result{
		Id:          res.Id,
		Price:       res.Price,
		Description: res.Description,
		Title:       res.Title,
	}, nil
}
