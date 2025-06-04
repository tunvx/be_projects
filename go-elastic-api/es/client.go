package es

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/delete"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
)

type Client interface {
	AddBook(ctx context.Context, book Book) (*index.Response, error)
	DeleteBook(ctx context.Context, bookID string) (*delete.Response, error)
	GetBook(ctx context.Context, bookID string) (*search.Response, error)
	FilterBooks(ctx context.Context, filter map[string]any) (*search.Response, error)
	FullTextSearch(ctx context.Context, query string) (*search.Response, error)
}

type ESClient struct {
	client *elasticsearch.TypedClient
}

func NewClient(client *elasticsearch.TypedClient) Client {
	return &ESClient{
		client: client,
	}
}
