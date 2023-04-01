# BET-438 - HTTP and DB with Go

Languages: Go, MySQL  
Frameworks: Gin (Go)  
Task code: https://github.com/ouahabs/LD-Academy/tree/master/BET-438%20-%20HTTP%20and%20DB%20with%20Go  

Task details: 
* A REST API for a jazz record store. 
* 2 Endpoints
	* /albums
		* GET - get a list of albums, json response
		* POST - add new album, json request
	* /albums/:id
		* GET - get album with ID, json response



### Notes
* *`json:"name"` is how a certain column or data field going to be marshalised in Go
* `gin.Context` (https://pkg.go.dev/github.com/gin-gonic/gin#Context) is a struct that allows for JSON rendering, has http.request information, etc
* More tutorials here (https://gowebexamples.com/)