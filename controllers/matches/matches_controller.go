package matches

import (
	"CricRadio/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListMatches(c *gin.Context) {

	matches, err := services.MatchesService.List()
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, matches.Marshall())
}
