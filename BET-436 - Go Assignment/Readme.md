# BET-436 - Go Assignement

Language: Go
Frameworks: Gin (Go)
Task code: 

Task details:
> In this task you'll create a simple Restful API for the following scenario: Suppose we want to create a simple API to help shops sell and manage products. A Shop should first sign up and provide the following information the shop's name and address. Shops can add their products, a product has the following information: the shop it belongs to, a name, a description and categories (A product can have multiple categories) the categories must be predefined and inserted to the database directly. In addition to these use cases we want endpoints: to list all shops, products, one shops details, one products details, edit shops and products (including their categories), delete shops and products. Authorization and authentication are good to have but not mandatory only use them if you have enough time.

Endpoints:
* /shops
	* `GET` - List all shops
* shops/:id
	* `GET` - List shop details by ID
	* `PUT` - Edit shop details
	* `DELETE` - Delete shop
* /login
	* `POST` - Log in with name and pass, uses JWT tokens
* /signup
	* `POST` - Add a shop
* /products
	* `GET` - List all products
	* `POST` - Add product
* /product/:id
	* `GET` - List product details by ID
	* `PUT` - Update product details
	* `DELETE` - Delete product
* /categories
	* `GET` - List all categories