# Article Web Service

## Architecture

[![QSCmBx.jpg](https://i.im.ge/2021/09/03/QSCmBx.jpg)](https://im.ge/i/QSCmBx)

### Services

* **Write Service** is a service in which there is an endpoint to create a new article. This service will save the article data to the main db (Postgres) and then publish the article data to the redis pub sub on the `article_created` topic.
* **Article Created Subscriber** is a service that subscribes to an `article_created` topic from the redis pub sub. After that this service will save the article data to elasticsearch.
* **Read Service** is a service in which there is an endpoint to get list of articles.

## Installation

### Prerequisite

* Docker version 19.03.13 or higher
* Docker Compose version 1.27.4 or higher

### How To Run

Run Write Service

```
docker-compose up --build write-service
```

Run Article Created Subscriber

```
docker-compose up --build article-created-subscriber
```

Run Read Service

```
docker-compose up --build read-service
```

## API Documentation

### Create New Article

```http
POST http://localhost:8080/v1/articles
```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `author` | `string` | **Required**. The author of article |
| `title` | `string` | **Required**. The title of article |
| `body` | `string` | **Required**. The body of article |

Example Request:

```curl
curl --location --request POST 'http://localhost:8080/v1/articles' \
--header 'Accept-Language: id' \
--header 'Content-Type: application/json' \
--data-raw '{
    "author": "John Doe",
    "title": "Lorem Ipsum",
    "body": "Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum"
}'
```

### Get List of Articles

```http
GET http://localhost:8082/v1/articles?author=&query=
```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `author` | `string` | **Optional**. To search keyword in article title and body |
| `query` | `string` | **Optional**. Filter by author name |

Example Request:

```curl
curl --location --request GET 'http://localhost:8082/v1/articles?author=John%20Doe&query=world' \
--header 'Accept-Language: id'
```
