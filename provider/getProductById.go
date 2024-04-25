package provider

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product interface {
	GetProductById(id string, ctx *gin.Context) (*Result, error)
}

func GetProductById(id string, ctx *gin.Context) (*Result, error) {
	response, err := http.Get("https://dummyjson.com/products/" + id)
	if err != nil {
		return nil, errors.New("InternalServer Error")
	}
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("cannot Read Response Body")
	}
	result := Result{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.New("cannot unmarshal struct")
	}
	return &result, nil
}
