# 📘 kb-api-go

An internal Knowledge Base backend API built with Go and MongoDB.

## Features

- CRUD for articles with `title`, `content`, `tags`, `createdAt`, `updatedAt`
- Separate tag collection with lookup and validation
- Filter articles by tags
- Pagination support
- MongoDB aggregation to enrich tag data
- Clean and idiomatic Go structure

## Setup

1. **Start MongoDB** (via Podman/macOS):

```bash
podman run --name mongo -d -p 27017:27017 mongo
```

2. **Run the server**

```bash
go run main.go
```

## API Endpoints

### Tags

| Method | Endpoint   | Description      |
|--------|------------|------------------|
| GET    | /tags      | List tags        |
| POST   | /tags      | Create a new tag |

### Articles

| Method | Endpoint            | Description                          |
|--------|---------------------|--------------------------------------|
| GET    | /articles           | List articles (filter, paginate)     |
| POST   | /articles           | Create article                       |
| GET    | /articles/{id}      | Get single article                   |
| PUT    | /articles/{id}      | Update article                       |
| DELETE | /articles/{id}      | Delete article                       |

## Example Article Payload

```json
{
  "title": "Intro to Go",
  "content": "Go is fast and fun.",
  "tags": ["<tag_id_1>", "<tag_id_2>"]
}
```
