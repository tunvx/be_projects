package es

import (
	"go-elastic-api/util"
	"log"
	"os"
	"testing"

	elastic "github.com/elastic/go-elasticsearch/v8"
)

var txKey = struct{}{}

var testClient Client

func TestMain(m *testing.M) {
	cfg, err := util.LoadConfig("..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	esClientTyped, err := elastic.NewTypedClient(elastic.Config{
		Addresses: []string{cfg.ElasticsearchServerAddress},
	})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch typed client: %s", err)
	}

	testClient = NewClient(esClientTyped)
	os.Exit(m.Run())
}
