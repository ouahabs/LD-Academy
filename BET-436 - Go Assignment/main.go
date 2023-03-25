package main

import (
	"github.com/gin-gonic/gin"

	"Academy/Controllers"
	db "Academy/Database"
)

func main() {
	router := gin.Default()
	db.ConnectToDB()

	//router.GET("/signup", Controllers.SignUp)
	// router.GET("/login", Controllers.LogIn)

	router.GET("/shops", Controllers.GetShops)
	router.GET("/shops/:id", Controllers.GetShopById)
	router.PUT("/shops/:id", Controllers.UpdateShop)
	router.DELETE("/shops/:id", Controllers.DeleteShop)

	// router.GET("/products", Controllers.GetProducts)
	// router.POST("/products", Controllers.AddProduct)

	// router.GET("/products/:id", Controllers.GetProductById)
	// router.PUT("/products/:id", Controllers.UpdateProduct)
	// router.DELETE("/products/:id", Controllers.DeleteProduct)

	// router.GET("/categories", Controllers.GetCategories)

	router.Run("localhost:8080")
}
