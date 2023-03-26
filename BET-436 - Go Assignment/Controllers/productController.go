package Controllers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	db "Academy/Database"
	"Academy/Models"
)

func CategoriesExist(categories []string, dbCategories []Models.Category) bool {
	set := make(map[string]bool)

	for _, v := range dbCategories {
		set[v.Name] = true
	}

	for _, v := range categories {
		if _, ok := set[strings.TrimSpace(v)]; !ok {
			return false
		}
	}

	return true
}

func GetProducts(c *gin.Context) {
	var products []Models.Product

	rows, err := db.Connection.Query("SELECT * FROM Products")

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again."})
	}

	defer rows.Close()

	for rows.Next() {
		var product Models.Product
		if err := rows.Scan(&product.ID, &product.ShopID, &product.Name, &product.Description, &product.Categories); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again."})
			return
		}
		products = append(products, product)
	}

	c.IndentedJSON(http.StatusOK, products)
}

func AddProduct(c *gin.Context) {
	var product Models.Product
	var category Models.Category

	if err := c.BindJSON(&product); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse JSON, please check formatting"})
		return
	}

	row := db.Connection.QueryRow("SELECT * FROM Products WHERE name = ?", product.Name)

	if err := row.Scan(&product.ID, &product.ShopID, &product.Name, &product.Description, &product.Categories); err != nil {
		if err != sql.ErrNoRows {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "This product already exists."})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again."})
	}

	categories := product.Categories
	dbCategories, _ := category.FindAll()

	if ok := CategoriesExist(categories, dbCategories); !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Category is not defined."})
		return
	}

	insert, err := db.Connection.Exec("INSERT INTO Products (shopid, name, description, categories) VALUES (?, ?, ?, ?)", product.ShopID, product.Name, product.Description, product.Categories)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again."})
	}

	id, err := insert.LastInsertId()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get product ID"})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"Shop ID": id, "message": "Success (Create)"})
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	var product Models.Product

	row := db.Connection.QueryRow("SELECT * FROM Product WHERE ID = ?", id)

	if err := row.Scan(&product.ID, &product.ShopID, &product.Name, &product.Description, &product.Categories); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found."})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again."})
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {

	id := c.Param("id")
	var updatedProduct Models.Product
	var category Models.Category

	if err := c.BindJSON(&updatedProduct); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse JSON, please check formatting"})
		return
	}

	categories := updatedProduct.Categories
	dbCategories, _ := category.FindAll()

	if ok := CategoriesExist(categories, dbCategories); !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Category is not defined."})
		return
	}

	if _, err := db.Connection.Exec("UPDATE Products SET name=?, Description=?, Categories=? WHERE id=?", updatedProduct.Name, updatedProduct.Description, updatedProduct.Categories, id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success (Update)"})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if _, err := db.Connection.Exec("DELETE FROM Shops WHERE id=?", id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success (Delete)"})
}
