package main

import (
	"bytes"
	"fmt"
	"go-elastic-api/api"
	"go-elastic-api/util"
	"io"
	"log"
	"strings"

	"go-elastic-api/es"

	"github.com/elastic/go-elasticsearch/v8"
	elastic "github.com/elastic/go-elasticsearch/v8"
)

func main() {
	// 1. Load config
	cfg, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config")
	}

	// 2. Create Elastic client
	esClientTyped, err := elastic.NewTypedClient(elastic.Config{
		Addresses: []string{cfg.ElasticsearchServerAddress},
	})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch typed client: %s", err)
	}
	esStore := es.NewClient(esClientTyped)

	// 3. Bulk insert mockdata into index "books"
	// bulkInsert(esClient)

	// 4. Initialize HTTP server
	server, err := api.NewServer(cfg, esStore)
	if err != nil {
		log.Fatalf("cannot create server")
	}

	// 5. Run server at port 8080
	fmt.Println("Starting server on :8000 …")
	if err := server.Start(cfg.HTTPServerAddress); err != nil {
		log.Fatalf("cannot start server")
	}
}

func deleteIndex(es *elasticsearch.Client, indexName string) error {
	res, err := es.Indices.Delete([]string{indexName})
	if err != nil {
		return fmt.Errorf("failed to delete index %s: %w", indexName, err)
	}
	defer res.Body.Close()

	if res.IsError() {
		if res.StatusCode == 404 {
			return nil
		}
		return fmt.Errorf("error deleting index %s: %s", indexName, res.String())
	}
	return nil
}

// bulkInsert pre-inserts a few book documents into the "books" index.
func bulkInsert(es *elasticsearch.Client) {
	if err := deleteIndex(es, "books"); err != nil {
		log.Fatalf("Delete index failed: %v", err)
	}

	// 3.1. Prepare NDJSON for Bulk API
	var buf bytes.Buffer
	bulkBody := `
	{ "index": { "_id": "9780553351927" }}
	{ "name": "Snow Crash", "author": "Neal Stephenson", "release_date": "1992-06-01", "page_count": 470 }
	{ "index": { "_id": "9780441017225" }}
	{ "name": "Revelation Space", "author": "Alastair Reynolds", "release_date": "2000-03-15", "page_count": 585 }
	{ "index": { "_id": "9780451524935" }}
	{ "name": "1984", "author": "George Orwell", "release_date": "1985-06-01", "page_count": 328 }
	{ "index": { "_id": "9781451673319" }}
	{ "name": "Fahrenheit 451", "author": "Ray Bradbury", "release_date": "1953-10-15", "page_count": 227 }
	{ "index": { "_id": "9780060850524" }}
	{ "name": "Brave New World", "author": "Aldous Huxley", "release_date": "1932-06-01", "page_count": 268 }
	{ "index": { "_id": "9780385490818" }}
	{ "name": "The Handmaid's Tale", "author": "Margaret Atwood", "release_date": "1985-06-01", "page_count": 311 }
	`
	buf.WriteString(strings.TrimSpace(bulkBody) + "\n")

	// 3.2. Call Bulk API
	res, err := es.Bulk(
		bytes.NewReader(buf.Bytes()),
		es.Bulk.WithIndex("books"),
	)
	if err != nil {
		log.Fatalf("Failed to execute bulk insert: %v", err)
	}
	defer res.Body.Close()

	// 3.3. Read response body and print
	b, _ := io.ReadAll(res.Body)
	fmt.Println("Bulk insert response:", string(b))
}
