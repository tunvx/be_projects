package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// filterBooks handles the filtering of books based on a JSON filter.
// It expects a JSON body with the filter criteria.
// If the body is invalid, it returns a 400 Bad Request error.
// If the filtering is successful, it returns a 200 OK response with the list of books found.
// If there is an error during filtering, it returns a 500 Internal Server Error.
// Example request body: {"term": {"author": "Some Author"}}
// Example response: [{"id": "1", "name": "Some Book", "author": "Some Author", ...}, ...]
func (server *Server) filterBooks(c *gin.Context) {
	filter := make(map[string]any)
	if err := c.BindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid filter format: %s", err)))
		return
	}

	res, err := server.esStore.FilterBooks(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("error filtering books: %s", err)))
		return
	}

	books := parseBooksTyped(res)
	c.JSON(http.StatusOK, books)
}
