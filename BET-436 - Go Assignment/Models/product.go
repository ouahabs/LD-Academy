package Models

type Product struct {
	ShopID      string   `json:"shopid"`
	Name        string   `json:"name"`
	Description string   `json:description`
	Categories  []string `json:categories`
}
