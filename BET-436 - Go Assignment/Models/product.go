package Models

type Product struct {
	ID          string   `json:"id"`
	ShopID      string   `json:"shopid"`
	Name        string   `json:"name"`
	Description string   `json:description`
	Categories  []string `json:categories`
}
