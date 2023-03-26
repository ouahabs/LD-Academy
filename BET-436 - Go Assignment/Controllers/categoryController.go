package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "Academy/Database"
	"Academy/Models"
)

func GetCategories(c *gin.Context) {
	var categories []Models.Category

	rows, err := db.Connection.Query("SELECT * FROM Categories")

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again."})
	}

	defer rows.Close()

	for rows.Next() {
		var category Models.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again."})
			return
		}
		categories = append(categories, category)
	}

	c.IndentedJSON(http.StatusOK, categories)
}
