package app

import (
	"github.com/Osedisc/practiceGo/getProduct"
	"github.com/Osedisc/practiceGo/provider"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	bindGetEmployee(r)

	r.Run()
}

func bindGetEmployee(r *gin.Engine) gin.IRoutes {
	provider := new(provider.Product)
	servicer := getProduct.NewService(*provider)
	handler := getProduct.NewHandler(servicer)
	return r.GET("/product/:id", handler.GetProductById)
}
