# Elasticsearch

##  Query Example

### 1. Find books by `name`
- curl -X GET "http://localhost:8000/search/name?name=1984"

- curl -X GET "http://localhost:8000/search/name?name=Brave%20New%20World"
- curl -X GET "http://localhost:8000/search/name?name=Brave%20Ne"

- curl -X GET "http://localhost:8000/search/name?name=Snow%20Crash"
- curl -X GET "http://localhost:8000/search/name?name=Snow"

### 2. Filter books by `author`
- curl -X GET "http://localhost:8000/filter/author?author=George%20Orwell"
- curl -X GET "http://localhost:8000/filter/author?author=Aldous%20Huxley"
- curl -X GET "http://localhost:8000/filter/author?author=Neal%20Stephenson"

### 3. Find books `release_date/after`: "YYYY-MM-DD"
- curl -X GET "http://localhost:8000/search/release_date?after=1980-01-01"

# Reference
- [Quickstart - Elastic.](https://www.elastic.co/docs/solutions/search/elasticsearch-basics-quickstart)
- [Querying and filterin - Elastic.](https://www.elastic.co/docs/explore-analyze/query-filter)
- [Query DSL - Elastic.](https://www.elastic.co/docs/explore-analyze/query-filter/languages/querydsl)
- https://viblo.asia/p/mot-so-cau-query-hay-su-dung-trong-elasticsearch-1VgZv0vY5Aw
- https://www.elastic.co/docs/solutions/search/querying-for-search
- https://discuss.elastic.co/t/best-practices-on-generating-queries-with-sdk-vs-storing-queries-as-files/303328
- https://opster.com/guides/elasticsearch/glossary/elasticsearch-query-syntax/