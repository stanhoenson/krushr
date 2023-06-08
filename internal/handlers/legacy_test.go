package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/stanhoenson/krushr/internal/middleware"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLegacyRoutes(t *testing.T) {
	r := gin.Default()
	r.Use(middleware.Authorization())
	handlers.RegisterLegacyRoutes(r)
	database.InitializeDatabase("test.db", "test/", true)
	populateDatabaseWithDummyRoutesData()

	t.Run("legacy", func(t *testing.T) {
		t.Run("testGetAllLegacyRoutes", func(t *testing.T) {
			testGetAllLegacyRoutes(t, r)
		})
		t.Run("testGetLegacyRoute", func(t *testing.T) {
			testGetLegacyRoute(t, r)
		})
		t.Run("testGetLegacyMenu", func(t *testing.T) {
			testGetLegacyMenu(t, r)
		})
	})

	os.Remove("test/test.db")
	os.Remove("test")
}

// TODO maybe check if conversion went right?
func testGetAllLegacyRoutes(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/legacy/route", nil)
	r.ServeHTTP(w, req)

	var legacyRoutes []models.LegacyRoute
	err := json.Unmarshal(w.Body.Bytes(), &legacyRoutes)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 4, len(legacyRoutes))
}

// TODO maybe check if conversion went right?
func testGetLegacyRoute(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/legacy/route/1", nil)
	r.ServeHTTP(w, req)

	var legacyRoute models.LegacyRoute
	err := json.Unmarshal(w.Body.Bytes(), &legacyRoute)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Example Route 1", legacyRoute.RouteName)
}

func testGetLegacyMenu(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/legacy/menu", nil)
	r.ServeHTTP(w, req)

	var legacyMenus []models.LegacyMenu
	err := json.Unmarshal(w.Body.Bytes(), &legacyMenus)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, len(legacyMenus))
}
