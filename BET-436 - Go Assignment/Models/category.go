package Models

import db "Academy/Database"

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (category Category) FindAll() ([]Category, error) {
	var categories []Category

	rows, err := db.Connection.Query("SELECT * FROM Categories")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
