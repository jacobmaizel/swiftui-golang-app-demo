package services

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

func AddServicesRoutes(r *gin.RouterGroup) {
	rg := r.Group("/services")

	rg.GET("/search", Search)
}

func Search(c *gin.Context) {

	search_term := utils.Sfq(c, "term")
	if search_term == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "term query param required"})
		return
	}

	search_term = strings.TrimSpace(search_term)

	for_type := utils.Sfq(c, "type")
	if for_type == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type query param required"})
		return
	}

	count, res, err := generateSearchResults(c, for_type, search_term)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// log.Println("res", res, "count", count)

	c.JSON(http.StatusOK, gin.H{"count": count, "results": res})

}
