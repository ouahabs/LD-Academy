package main

import (
	"github.com/gin-gonic/gin"

	"Academy/Controllers"
	db "Academy/Database"
	"Academy/Middleware"
)

func main() {
	router := gin.Default()
	db.ConnectToDB()

	router.POST("/signup", Controllers.SignUp)
	router.POST("/login", Controllers.LogIn)

	router.GET("/shops", Middleware.VerifyAuth(), Controllers.GetShops)
	router.GET("/shops/:id", Controllers.GetShopById)
	router.PUT("/shops/:id", Middleware.VerifyAuth(), Controllers.UpdateShop)
	router.DELETE("/shops/:id", Middleware.VerifyAuth(), Controllers.DeleteShop)

	router.GET("/products", Controllers.GetProducts)
	router.POST("/products", Middleware.VerifyAuth(), Controllers.AddProduct)

	router.GET("/products/:id", Controllers.GetProductById)
	router.PUT("/products/:id", Middleware.VerifyAuth(), Controllers.UpdateProduct)
	router.DELETE("/products/:id", Middleware.VerifyAuth(), Controllers.DeleteProduct)

	router.GET("/categories", Controllers.GetCategories)

	router.Run("localhost:8080")
}
