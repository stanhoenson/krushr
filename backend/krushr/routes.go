package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var routes = []Route{
	{ID: "0", Title: "Bernard's Route", StatusID: "0", UserID: "2"},
	{ID: "1", Title: "Het Pad der 7 Zonden", StatusID: "0", UserID: "1"},
	{ID: "2", Title: "Reflecties en Introspecties", StatusID: "1", UserID: "2"},
}

func GetRoutes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, routes)
}

func GetRouteByID(c *gin.Context) {
	id := c.Param("id")

	for _, r := range routes {
		if r.ID == id {
			c.IndentedJSON(http.StatusOK, r)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "route not found"})
}

func PostRoutes(c *gin.Context) {
	var newRoute Route

	if err := c.BindJSON(&newRoute); err != nil {
		return
	}

	routes = append(routes, newRoute)
	c.IndentedJSON(http.StatusCreated, newRoute)
}
