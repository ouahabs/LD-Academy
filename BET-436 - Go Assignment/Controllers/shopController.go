package Controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"Academy/Authentication"
	db "Academy/Database"
	"Academy/Models"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignUp(c *gin.Context) {
	var shop Models.Shop

	if err := c.BindJSON(&shop); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse JSON, please check formatting"})
		return
	}
	row := db.Connection.QueryRow("SELECT * FROM Shops WHERE name = ?", shop.Name)

	if err := row.Scan(&shop.ID, &shop.Name, &shop.Address); err != nil {
		if err != sql.ErrNoRows {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "This shop already exists."})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again."})
	}

	hash, err := hashPassword(shop.Password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again."})
	}

	shop.PasswordHash = hash

	insert, err := db.Connection.Exec("INSERT INTO Shops (name, address, password, hash) VALUES (?, ?, ?, ?)", shop.Name, shop.Address, shop.Password, shop.PasswordHash)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again."})
	}

	id, err := insert.LastInsertId()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get shop ID"})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"Shop ID": id, "message": "Success (Create)"})
}

func LogIn(c *gin.Context) {
	var shop, shopRequest Models.Shop

	if err := c.BindJSON(&shopRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse JSON, please check formatting."})
		return
	}

	ok, err := shop.Find(shopRequest.Name)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Error."})
		return
	}

	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Shop not found."})
		return
	}

	ok = checkPasswordHash(shopRequest.Password, shop.PasswordHash)

	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Wrong password, please try again."})
		return
	}

	token, err := Authentication.GetToken(shopRequest)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Could not create your token, please try again."})
	}

	c.Header("Token", token)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success (login)", "data": shopRequest.Name})

}

func GetShops(c *gin.Context) {
	var shops []Models.Shop

	rows, err := db.Connection.Query("SELECT * FROM Shops")

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again."})
	}

	defer rows.Close()

	for rows.Next() {
		var shop Models.Shop
		if err := rows.Scan(&shop.ID, &shop.Name, &shop.Address); err != nil {
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse JSON, please check formatting"})
		return
	}

	if _, err := db.Connection.Exec("UPDATE Shops SET name=?, address=? WHERE id=?", updatedShop.Name, updatedShop.Address, id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Shop not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success (Update)"})
}

func DeleteShop(c *gin.Context) {
	id := c.Param("id")

	if _, err := db.Connection.Exec("DELETE FROM Shops WHERE id=?", id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Shop not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success (Delete)"})
}
