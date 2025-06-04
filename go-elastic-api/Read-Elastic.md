# ElasticSearch Architecture: A Comprehensive Guide

Elasticsearch is a powerful data search and analysis engine built on Apache Lucene. It is designed to handle structured and unstructured data, from text to numeric data, with horizontal scalability and near real-time search speed. This article will help you understand the architecture of Elasticsearch in detail, how it stores, processes, queries, and scales the system.

---

# 1. Elasticsearch Architecture Overview

Elasticsearch is a distributed system, allowing data storage, search, and analysis across multiple servers (nodes). It supports RESTful API, allows easy interaction with the system via HTTP, and operates in an asynchronous model, optimized for high performance.

Elasticsearch is developed in Java and leverages Apache Lucene as its core search library. Lucene is a powerful library for full-text search, but it is not a complete application. Elasticsearch wraps Lucene and extends it with distributed features, making it suitable for large-scale applications. While Lucene provides the core search capabilities, Elasticsearch adds the distribution, scalability, and ease-of-use features that make it a popular choice for modern applications.

## Key Features:
- **Distributed Nature:** Elasticsearch is designed to scale horizontally, allowing you to add more nodes to the cluster as your data grows.

- **Real-Time Search:** Elasticsearch provides near real-time search capabilities, making it ideal for applications that require fast query responses.

- **Schema-less:** Elasticsearch is schema-less by default, meaning you can start indexing documents without defining a schema. However, ***you can define mappings to control how data is indexed and searched.***

## Main components:

* **Cluster**: A set of nodes that work together. Each cluster has a unique name.

* **Node**: A running instance of Elasticsearch. A node can be a master, data, or ingest node.

* **Index**: A logical unit of data storage, similar to a "table" in a relational database.

* **Shard**: An index is split into multiple shards to distribute data.

* **Replica**: A replica of a shard to ensure availability and fault tolerance.

## Comparison to RDBMS:

| Elasticsearch | RDBMS       |
| ------------- | ----------- |
| Cluster       | Server      |
| Node          | Instance    |
| **Index**     | **Table**   |
| **Document**  | **Row**     |
| **Field**     | **Column**  |
| **Mapping**   | **Schema**  |
| \_id          | Primary Key |

---

# 2. Storage Model
<div align="center">
  <img src="imgs/Elastic Storage Model.png" alt="Image" width="600px">
</div>

Elasticsearch stores data in **indexes**, and each index consists of many **shards**. Each shard is an instance of Apache Lucene, capable of storing, searching, and analyzing data independently.

* **Primary Shard**: The primary shard, containing the original data.

* **Replica Shard**: A copy of the primary shard, used to increase availability.

Sharding helps Elasticsearch scale horizontally by adding nodes.

## Storage structure:
```
Cluster/node
 ├── Index (example: logs_2025)
     ├── Shard 0 (Primary + Replica): Document_000, Document_001, Document_002, ..
     ├── Shard 1 (Primary + Replica): Document_100, Document_101, Document_102, ..
     └── ...
```

---

# 3. Data Types and Mapping
Elasticsearch receives data in JSON format and performs a process called mapping to map incoming data to Elasticsearch's data types. Mapping defines how fields in the document should be indexed and searched.

## Common Data Types:
- **Keyword:** Used for exact ***matches***, such as IDs or tags.
- **Text:** Used for full-text search. ***Text fields undergo text analysis***, which includes `tokenization` and `normalization`.
- **Date:** Used for date and time values.
- **Numeric:** Used for numerical data, such as integers, floats, and doubles.

## Text Analysis:
When a text field is indexed, Elasticsearch performs text analysis, which involves:

- **Tokenization:** Breaking text into individual tokens (words).
- **Normalization:** Adding metadata, such as synonyms or translations, to make the text searchable.

---

# 4. Text Analysis and Inverted Indexes
Text analysis is a critical part of Elasticsearch's search capabilities. When a document is indexed, Elasticsearch creates an **inverted index**, which maps tokens to the documents that contain them. This allows for fast and efficient full-text search.

## Inverted Index:
Consider two documents:

1. `{"greeting": "Hello, WORLD"}`
2. `{"greeting": "Hello, Mate"}`

After text analysis, the inverted index might look like this:

| Token	Document | IDs |
|----------------|-----|
| hello | 1, 2 |
| world | 1 |
| mate | 2 |

When you search for "hello," Elasticsearch consults the inverted index and retrieves documents 1 and 2.

--- 

# 5. Building Blocks: Documents, Indexes, and Shards
## Documents:
- **Basic Unit:** A document is the basic unit of data in Elasticsearch. It is a JSON object that contains the data you want to index and search.
- **Fields:** Each document consists of fields, which are key-value pairs. Fields can be of different data types, such as text, keyword, date, or numeric.

## Indexes:
- **Logical Collection:** An index is a logical collection of documents. It is similar to a table in a relational database.
- **Shards:** Each index is divided into shards, which are physical instances of Apache Lucene. Shards are distributed across nodes for scalability and fault tolerance.

## Shards:
- **Primary Shards:** Each document belongs to a primary shard. The number of primary shards is fixed when the index is created.
- **Replicas:** Replicas are copies of primary shards. They provide redundancy and improve search performance by serving read requests.

<div align="center">
  <img src="imgs/Elastic Bulding Blocks.png" alt="Image" width="800px">
</div>

---

# 6. Indexing and Search Process

## Indexing:

1. Send JSON document to an index.

2. Elasticsearch moves document to a shard.

3. Data is analyzed, tokenized, and stored in Lucene (shard).

## Search:

1. Send query to index.

2. Elasticsearch distributes query to shards.

3. Results from shard are aggregated and returned to user.

---

# 7. Querying with Query DSL
Query DSL is a full-featured JSON-style query language that enables complex searching, filtering, and aggregations. It is the original and most powerful query language for Elasticsearch today.

The `_search` endpoint accepts queries written in Query DSL syntax.

## Search and filter with Query DSL
Query DSL support a wide range of search techniques, including the following:

- **Full-text search:** Search text that has been analyzed and indexed to support phrase or proximity queries, fuzzy matches, and more.
- **Keyword search:** Search for exact matches using keyword fields.
- **Semantic search:** Search semantic_text fields using dense or sparse vector search on embeddings generated in your Elasticsearch cluster.
- **Vector search:** Search for similar dense vectors using the kNN algorithm for embeddings generated outside of Elasticsearch.
- **Geospatial search:** Search for locations and calculate spatial relationships using geospatial queries.

You can also filter data using Query DSL. Filters enable you to include or exclude documents by retrieving documents that match specific field-level criteria. A query that uses the filter parameter indicates filter context.

## Aggregations - Analyze with Query DSL
**Aggregations** are the primary tool for analyzing Elasticsearch data using Query DSL. Aggregations enable you to build complex summaries of your data and gain insight into key metrics, patterns, and trends.

Because aggregations leverage the same data structures used for search, they are also very fast. This enables you to analyze and visualize your data in real time. You can search documents, filter results, and perform analytics at the same time, on the same data, in a single request. That means aggregations are calculated in the context of the search query.

The following aggregation types are available:

- **Metric:** Calculate metrics, such as a sum or average, from field values.
- **Bucket:** Group documents into buckets based on field values, ranges, or other criteria.
- **Pipeline:** Run aggregations on the results of other aggregations.

## Query context: Query vs. Filter
- **Query context:** In the query context, a query clause answers the question How well does this document match this query clause? Besides deciding whether or not the document matches, the query clause also calculates a relevance score in the `_score` metadata field.

- Query context is in effect whenever a query clause is passed to a query parameter, such as the query parameter in the search API.

- **Filter context:** A filter answers the binary question “Does this document match this query clause?”. The answer is simply "yes" or "no". Does not calculate a score. Often used for precise filter conditions and can be cached to improve performance.

### Example of query and filter contexts
Below is an example of query clauses being used in query and filter context in the search API. This query will match documents where all of the following conditions are met:

- The `title` field contains the word `search`.
- The `content` field contains the word `elasticsearch`.
- The `status` field contains the exact word `published`.
- The `publish_date` field contains a date from 1 Jan 2015 onwards.

```json
GET /_search
{
  "query": { (1)
    "bool": { (2)
      "must": [
        { "match": { "title":   "Search"        }},
        { "match": { "content": "Elasticsearch" }}
      ],
      "filter": [ (3)
        { "term":  { "status": "published" }},
        { "range": { "publish_date": { "gte": "2015-01-01" }}}
      ]
    }
  }
}
```

- (1): The `query` parameter indicates **query context**.
- (2): The `bool` and two `match` clauses are used in query context, which means that they are used to score how well each document matches.
- (3): The `filter` parameter indicates **filter context**. Its `term` and `range` clauses are used in filter context. They will filter out documents which do not match, but they will not affect the score for matching documents.


## Golang API for Query DSL 
**Official Go Client:** `elastic/go-elasticsearch`

### `NewClient`: Low-Level Client
The NewClient function creates a low-level client that provides direct access to Elasticsearch's REST API. With this client, you typically need to manually construct JSON payloads for queries, which can be error-prone and less convenient.

```go
client, err := elasticsearch.NewClient(elasticsearch.Config{
    CloudID: "<CloudID>",
    APIKey:  "<ApiKey>",
})
```

Using this client, you might perform a search like:

```go
query := `{ "query": { "match_all": {} } }`
client.Search(
    client.Search.WithIndex("my_index"),
    client.Search.WithBody(strings.NewReader(query)),
)
```

In this approach, you manually write the JSON query as a string, which can be cumbersome and error-prone, especially for complex queries.

### `NewTypedClient`: High-Level Typed Client
The **NewTypedClient** function introduces a higher-level, strongly-typed API that allows you to construct queries using Go structs, providing better compile-time checks and reducing the likelihood of errors.

```go
typedClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
    CloudID: "<CloudID>",
    APIKey:  "<ApiKey>",
})
```

With the typed client, you can build queries using Go structs:
```go
typedClient.Search().
  Index("my_index").
  Request(&search.Request{
      Query: &types.Query{MatchAll: &types.MatchAllQuery{}},
  }).
  Do(context.TODO())
```

This approach leverages Go's type system, making it easier to construct and maintain complex queries without manually crafting JSON strings.


