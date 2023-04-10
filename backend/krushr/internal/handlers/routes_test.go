package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// func TestGetRouteByID(t *testing.T) {
// 	r := gin.Default()
// 	r.GET("/routes/:id", handlers.GetRouteByID)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/routes/:1", nil)
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, "hello", w.Body.String())
// }

func TestGetRouteByIDWithInvalidID(t *testing.T) {
	r := gin.Default()
	r.GET("/routes/:id", handlers.GetRouteByID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/routes/:3", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"Invalid ID parameter\"}", w.Body.String())
}

// func TestDeleteRouteByID(t *testing.T) {
// 	r := gin.Default()
// 	r.DELETE("/routes/:id", handlers.DeleteRouteByID)

// 	route := models.Route{
// 		Title:    "test",
// 		StatusID: 1,
// 		UserID:   1,
// 	}
// 	repositories.CreateEntity(&route)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/routes/:1", nil)
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, "hello", w.Body.String())
// }

func TestDeleteRouteByIDWithInvalidID(t *testing.T) {
	r := gin.Default()
	r.DELETE("/routes/:id", handlers.DeleteRouteByID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/routes/:1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"Invalid ID parameter\"}", w.Body.String())
}

// func TestPutRoute(t *testing.T) {
// 	r := gin.Default()
// 	r.PUT("/routes", handlers.PutRoute)

// 	updatedRoute := models.Route{}
// 	payload, _ := json.Marshal(updatedRoute)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("PUT", "/routes", bytes.NewBuffer(payload))
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, "hello", w.Body.String())
// }
