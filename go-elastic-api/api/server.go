package api

import (
	"encoding/json"
	"go-elastic-api/util"

	"go-elastic-api/es"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config  util.Config
	esStore es.Client
	router  *gin.Engine
}

func NewServer(cfg util.Config, esStore es.Client) (*Server, error) {
	server := &Server{
		config:  cfg,
		esStore: esStore,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// Define routes
	// 1. Search by full_text_search: /search/full_text_search?query_str=random_string
	router.GET("/search/full_text_search", server.fullTextSearch)

	// 2. Filters books based on a JSON body.
	router.POST("/filter/books", server.filterBooks)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func parseBooksTyped(res *search.Response) []es.Book {
	var books []es.Book
	for _, hit := range res.Hits.Hits {
		var book es.Book
		if err := json.Unmarshal(hit.Source_, &book); err == nil {
			book.ID = *hit.Id_
			books = append(books, book)
		}
	}
	return books
}
