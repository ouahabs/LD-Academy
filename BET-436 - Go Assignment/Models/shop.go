package Models

import (
	db "Academy/Database"
	"database/sql"
)

type Shop struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Address      string `json:address`
	Password     string `json:password`
	PasswordHash string `json:hash`
}

func (shop *Shop) Find(name string) (bool, error) {

	row := db.Connection.QueryRow("SELECT * FROM Shops WHERE name = ?", name)

	if err := row.Scan(&shop.ID, &shop.Name, &shop.Address, &shop.Password, &shop.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
