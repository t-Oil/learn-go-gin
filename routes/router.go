package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"learn/go-gin/controller"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
}

func SetupRouter() *gin.Engine {

	Router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "aa")
	})

	Router.GET("/products", controller.GetProducts)
	Router.GET("/products/:id", controller.FindProducts)
	Router.POST("/products", controller.StoreProduct)
	Router.PUT("/products/:id", controller.UpdateProduct)

	Router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "NOT FOUND")
	})

	return Router
}
