package es

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"go-elastic-api/util"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/stretchr/testify/require"
)

func TestFilterBooks(t *testing.T) {
	// Create a random book
	book := createRandomBook()

	// Add the book to Elasticsearch
	_, err := testClient.AddBook(context.Background(), book)
	require.NoError(t, err)
	defer testClient.DeleteBook(context.Background(), book.ID)

	// Wait a bit to ensure indexing (optional)
	time.Sleep(1 * time.Second)

	// Filter books by a specific field (e.g., author)
	filter := map[string]any{
		"term": map[string]any{
			"author": book.Author,
		},
	}

	res, err := testClient.FilterBooks(context.Background(), filter)
	require.NoError(t, err)

	// Check if we got results
	require.NotEmpty(t, res.Hits.Hits)

	// Verify the first hit matches the book we added
	checkResult(t, res, book)
}

// TestFullTextSearch tests the full-text search functionality
func TestFullTextSearch(t *testing.T) {
	// Create a random book
	book := createRandomBook()

	// Add the book to Elasticsearch
	_, err := testClient.AddBook(context.Background(), book)
	require.NoError(t, err)
	defer testClient.DeleteBook(context.Background(), book.ID)

	// Wait a bit to ensure indexing (optional)
	time.Sleep(1 * time.Second)

	// Perform full-text search on the book's name
	query := fmt.Sprintf("name:%s", book.Name)
	res, err := testClient.FullTextSearch(context.Background(), query)
	require.NoError(t, err)

	// Check if we got results
	require.NotEmpty(t, res.Hits.Hits)

	// Verify the first hit matches the book we added
	checkResult(t, res, book)
}

func TestAddBook(t *testing.T) {
	book := createRandomBook()

	// Add book
	res, err := testClient.AddBook(context.Background(), book)
	require.NoError(t, err)
	defer testClient.DeleteBook(context.Background(), book.ID)

	require.Equal(t, "created", res.Result.String())
	require.Equal(t, book.ID, res.Id_)

	// Wait a bit to ensure indexing (optional)
	time.Sleep(1 * time.Second)

	// Get book back
	tes, ok := testClient.(*ESClient)
	require.True(t, ok, "testClient is not of type *ESClient")

	getRes, err := tes.client.Get("books", book.ID).Do(context.Background())
	require.NoError(t, err)
	require.True(t, getRes.Found)

	var got Book
	err = json.Unmarshal(getRes.Source_, &got)
	require.NoError(t, err)

	// Compare the book
	require.Equal(t, book.ID, got.ID)
	require.Equal(t, book.Name, got.Name)
	require.Equal(t, book.Author, got.Author)
	require.Equal(t, book.ReleaseDate, got.ReleaseDate)
	require.Equal(t, book.PageCount, got.PageCount)
}

func TestDeleteBook(t *testing.T) {
	book := createRandomBook()

	// Add book first
	_, err := testClient.AddBook(context.Background(), book)
	require.NoError(t, err)

	// Delete book
	res, err := testClient.DeleteBook(context.Background(), book.ID)
	require.NoError(t, err)
	require.Equal(t, "deleted", res.Result.String())

	// Wait a bit to ensure deletion (optional)
	time.Sleep(1 * time.Second)

	// Try to get the deleted book
	tes, ok := testClient.(*ESClient)
	require.True(t, ok, "testClient is not of type *ESClient")

	getRes, err := tes.client.Get("books", book.ID).Do(context.Background())
	require.NoError(t, err)
	require.False(t, getRes.Found)
}

func TestGetBook(t *testing.T) {
	// Create a random book
	book := createRandomBook()

	// Add the book to Elasticsearch
	_, err := testClient.AddBook(context.Background(), book)
	require.NoError(t, err)
	defer testClient.DeleteBook(context.Background(), book.ID)

	// Wait a bit to ensure indexing (optional)
	time.Sleep(1 * time.Second)

	// Get the book back
	res, err := testClient.GetBook(context.Background(), book.ID)
	require.NoError(t, err)

	// Check if we got results
	require.NotEmpty(t, res.Hits.Hits)

	// Verify the first hit matches the book we added
	checkResult(t, res, book)
}

func createRandomBook() Book {
	book := Book{
		ID:          fmt.Sprintf("%d", rand.Int63()),
		Name:        util.RandomString(10),
		Author:      util.RandomAuthor(),
		Edition:     util.RandomEdition(),
		Publisher:   util.RandomPublisher(),
		ReleaseDate: util.RandomReleaseDate(),

		Description: util.RandomDescription(),
		PageCount:   util.RandomPageCount(),
		Content:     util.RandomContent(),

		Categories:  util.RandomCategories(),
		Tags:        util.RandomTags(),
		Rating:      util.RandomRating(),
		ReviewCount: util.RandomReviewCount(),
	}
	return book
}

func checkResult(t *testing.T, res *search.Response, book Book) {
	found := false
	for _, hit := range res.Hits.Hits {
		var got Book
		err := json.Unmarshal(hit.Source_, &got)
		require.NoError(t, err)
		fmt.Println("Got book:", got, "from search:", book.ID)

		if got.ID == book.ID {
			found = true
			require.Equal(t, book.ID, got.ID)
			require.Equal(t, book.Name, got.Name)
			require.Equal(t, book.Author, got.Author)
			require.Equal(t, book.ReleaseDate, got.ReleaseDate)
			require.Equal(t, book.PageCount, got.PageCount)
			require.Equal(t, book.Edition, got.Edition)
			require.Equal(t, book.Publisher, got.Publisher)
			require.Equal(t, book.Description, got.Description)
			require.Equal(t, book.Content, got.Content)
			require.ElementsMatch(t, book.Categories, got.Categories)
			require.ElementsMatch(t, book.Tags, got.Tags)
			require.Equal(t, book.Rating, got.Rating)
			require.Equal(t, book.ReviewCount, got.ReviewCount)
			break
		}
	}
	require.True(t, found, "Expected book not found in search results")
}
