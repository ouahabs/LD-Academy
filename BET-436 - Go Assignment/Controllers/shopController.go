package Controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	db "Academy/Database"
	"Academy/Models"
)

func GetShops(c *gin.Context) {
	var shops []Models.Shop

	rows, err := db.Connection.Query("SELECT * FROM Shops")

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again."})
	}

	defer rows.Close()

	for rows.Next() {
		var shop Models.Shop
		if err := rows.Scan(&shop.ID, &shop.Name, &shop.Address, &shop.PasswordHash); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again."})
			return
		}
		shops = append(shops, shop)
	}

	c.IndentedJSON(http.StatusOK, shops)
}

func GetShopById(c *gin.Context) {
	id := c.Param("id")
	var shop Models.Shop

	row := db.Connection.QueryRow("SELECT * FROM Shops WHERE ID = ?", id)

	if err := row.Scan(&shop.ID, &shop.Name, &shop.Address, &shop.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Shop not found."})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again."})
		return
	}

	c.IndentedJSON(http.StatusOK, shop)
}

func UpdateShop(c *gin.Context) {

	id := c.Param("id")
	var updatedShop Models.Shop

	if err := c.BindJSON(&updatedShop); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Couldn't bind JSON, please check formatting"})
		return
	}

	if _, err := db.Connection.Exec("UPDATE Shops SET name=?, address=? WHERE id=?", updatedShop.Name, updatedShop.Address, id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Shop not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success (Update)"})
	return
}

func DeleteShop(c *gin.Context) {
	id := c.Param("id")

	if _, err := db.Connection.Exec("DELETE FROM Shops WHERE id=?", shop.ID); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Shop not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success (Delete)"})
	return
}
