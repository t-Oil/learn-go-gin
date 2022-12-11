package controller

import (
	"net/http"
	"strconv"

	"learn/go-gin/model"

	"github.com/gin-gonic/gin"
)

func GetProducts(ctx *gin.Context) {
	products, err := model.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func FindProducts(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	product, err := model.FindProduct(id)

	if (err != nil || product == model.Product{}) {
		ctx.JSON(http.StatusNotFound, "Not Found")
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func StoreProduct(ctx *gin.Context) {
	var req model.ProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := model.StoreProduct(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "success")
}

func UpdateProduct(ctx *gin.Context) {
	var req model.ProductRequest

	id, _ := strconv.Atoi(ctx.Param("id"))

	findById, err := model.FindProduct(id)

	if (err != nil || findById == model.Product{}) {
		ctx.JSON(http.StatusNotFound, "Not Found")
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	product, err := model.UpdateProduct(id, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, product)
}
