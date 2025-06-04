package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// fullTextSearch handles the full-text search for books by their name.
// It expects a query string parameter "query_str" which is the name of the book to search for.
// If the parameter is missing, it returns a 400 Bad Request error.
// If the search is successful, it returns a 200 OK response with the list of books found.
// If there is an error during the search, it returns a 500 Internal Server Error.
// Example request: GET /api/v1/books/full_text_search?query_str=some_book_name
// Example response: [{"id": "1", "name": "Some Book", "author": "Some Author", ...}, ...]
func (server *Server) fullTextSearch(c *gin.Context) {
	query_str := c.Query("query_str")
	if query_str == "" {
		c.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("name parameter is required")))
		return
	}

	res, err := server.esStore.FullTextSearch(c.Request.Context(), query_str)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("error searching for books by name: %s", err)))
		return
	}

	books := parseBooksTyped(res)
	c.JSON(http.StatusOK, books)
}
