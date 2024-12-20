package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type listItems struct {
	ID   string `json:"id"`
	Item string `json:"item"`
}

var items = []listItems{
	{ID: "1", Item: "First Item"},
	{ID: "2", Item: "Second Item"},
}

// Return list of items
func getList(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, items)
}

// Add to list
func addItem(context *gin.Context) {
	var newItem listItems

	if err := context.BindJSON(&newItem); err != nil {
		return
	}

	items = append(items, newItem)
	context.IndentedJSON(http.StatusCreated, newItem)

}

// Get path parameter
func getItemId(context *gin.Context) {
	id := context.Param("id")
	item, err := getSingleItem(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, item)

}

// Get specific item in list
func getSingleItem(id string) (*listItems, error) {
	for i, t := range items {
		if t.ID == id {
			return &items[i], nil
		}
	}

	return nil, errors.New("item not found")
}

// GIN Server
func main() {
	router := gin.Default()            //this is the server
	router.GET("/list", getList)       //this is the GET using func getList
	router.GET("/list/:id", getItemId) //this is the GET for single item
	router.POST("/add", addItem)       //this is the POST using func addItem
	router.Run("localhost:7676")       //this is the server
}
