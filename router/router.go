package router

import (
	"tugas10/controllers"
	"tugas10/middlewares"

	"github.com/gin-gonic/gin"
)


func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/",middlewares.UserAuthorization(),controllers.GetProducts)
		productRouter.GET("/:productId", middlewares.ProductAuthorization(),controllers.GetProduct)
		productRouter.PUT("/:productId", middlewares.UserAuthorization(),controllers.UpdateProduct)
		productRouter.DELETE("/:productId", middlewares.UserAuthorization(),controllers.DeleteProduct)
	}
	return r
}