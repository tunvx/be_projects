package es

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/delete"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type Book struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Edition     string `json:"edition"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"release_date"`

	// Content and description
	Description string `json:"description,omitempty"`
	PageCount   int    `json:"page_count"`
	Content     string `json:"content,omitempty"`

	// Classification and evaluation
	Categories  []string `json:"categories,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Rating      float32  `json:"rating,omitempty"`
	ReviewCount int      `json:"review_count,omitempty"`
}

type BookInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Edition     string `json:"edition"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"release_date"`

	// Content and description
	Description string `json:"description,omitempty"`
	PageCount   int    `json:"page_count"`

	// Classification and evaluation
	Categories  []string `json:"categories,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Rating      float32  `json:"rating,omitempty"`
	ReviewCount int      `json:"review_count,omitempty"`
}

func (es *ESClient) AddBook(ctx context.Context, book Book) (*index.Response, error) {
	res, err := es.client.Index("books").
		Id(book.ID).
		Request(book).
		Do(ctx)
	return res, err
}

func (es *ESClient) DeleteBook(ctx context.Context, bookID string) (*delete.Response, error) {
	res, err := es.client.Delete("books", bookID).Do(ctx)
	return res, err
}

func (es *ESClient) GetBook(ctx context.Context, bookID string) (*search.Response, error) {
	return es.client.Search().
		Index("books").
		Request(&search.Request{
			Query: &types.Query{
				Term: map[string]types.TermQuery{
					"_id": {
						Value: bookID,
					},
				},
			},
			Size: func(i int) *int { return &i }(1), // Limit to 1 result
		}).Do(ctx)
}

func (es *ESClient) FilterBooks(ctx context.Context, filter map[string]any) (*search.Response, error) {
	var mustClauses []types.Query

	if author, ok := filter["author"].(string); ok {
		mustClauses = append(mustClauses, types.Query{
			Match: map[string]types.MatchQuery{
				"author": {Query: author},
			},
		})
	}

	if publisher, ok := filter["publisher"].(string); ok {
		mustClauses = append(mustClauses, types.Query{
			Term: map[string]types.TermQuery{
				"publisher.keyword": {Value: publisher},
			},
		})
	}

	if categories, ok := filter["categories"].([]string); ok && len(categories) > 0 {
		mustClauses = append(mustClauses, types.Query{
			Terms: &types.TermsQuery{
				TermsQuery: map[string]types.TermsQueryField{
					"categories.keyword": categories,
				},
			},
		})
	}

	if releaseDateAfter, ok := filter["release_after"].(string); ok {
		mustClauses = append(mustClauses, types.Query{
			Range: map[string]types.RangeQuery{
				"release_date": &types.DateRangeQuery{
					Gte: &releaseDateAfter,
				},
			},
		})
	}

	return es.client.Search().
		Index("books").
		Request(&search.Request{
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Must: mustClauses,
				},
			},
		}).
		Do(ctx)
}

func (es *ESClient) FullTextSearch(ctx context.Context, query string) (*search.Response, error) {
	return es.client.Search().
		Index("books").
		Request(&search.Request{
			Query: &types.Query{
				QueryString: &types.QueryStringQuery{
					Query: query,
				},
			},
		}).
		Do(ctx)
}
