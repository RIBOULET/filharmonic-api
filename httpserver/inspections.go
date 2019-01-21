package httpserver

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (server *HttpServer) listInspections(c *gin.Context) {
	inspections, err := server.service.ListInspections(server.retrieveUserContext(c))
	if err != nil {
		log.Error().Err(err).Msg("Bad service response")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, inspections)
}

func (server *HttpServer) getInspection(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	inspection, err := server.service.GetInspection(server.retrieveUserContext(c), id)
	if err != nil {
		log.Error().Err(err).Msg("Bad service response")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if inspection == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not_found",
		})
		return
	}
	c.JSON(http.StatusOK, inspection)
}
