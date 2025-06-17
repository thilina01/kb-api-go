# 📘 KB API (Go + MongoDB)

An internal **Knowledge Base (KB) backend API** built using Go and MongoDB. This project is containerized with Docker and compatible with VS Code Dev Containers.

---

## ✨ Features

- 📝 CRUD for Articles (`title`, `content`, `tags`, timestamps)
- 🏷️ Tag Management (`create`, `list`)
- 🔗 Article-Tag linking with validation
- 🔍 Filter Articles by tag + pagination support
- 📦 MongoDB storage, accessed via Go MongoDB Driver
- 🐳 Fully Dockerized with Compose (API + MongoDB)
- 🧪 HTTP test suite (.http file)
- 💻 VS Code + DevContainer support

---

## 🚀 Getting Started

### 1. Clone & Run with Docker Compose

```bash
git clone https://github.com/thilina01/kb-api-go.git
cd kb-api-go
docker compose up --build
```

### 2. Health Check

```bash
http GET :8080/ping
# Response: { "status": "ok" }
```

---

## 📚 API Reference

### Tags

| Method | Endpoint   | Description      |
|--------|------------|------------------|
| GET    | /tags      | List all tags    |
| POST   | /tags      | Create a new tag |

### Articles

| Method | Endpoint         | Description                          |
|--------|------------------|--------------------------------------|
| GET    | /articles        | List all articles (with filters)     |
| POST   | /articles        | Create a new article                 |
| GET    | /articles/{id}   | Get a single article                 |
| PUT    | /articles/{id}   | Update article by ID                 |
| DELETE | /articles/{id}   | Delete article by ID                 |

---

## 🧪 Testing

You can test endpoints using the included `.http` file in VS Code or using HTTPie.

```bash
http POST :8080/tags name="go"
http GET :8080/tags
http POST :8080/articles title="Go" content="Great" tags:='["<tag_id>"]'
```

Or use `api-tests.http` with REST Client in VS Code.

---

## 🧰 Dev Setup

### Local Development

```bash
go run main.go
# Default Mongo URI: mongodb://localhost:27017
```

### VS Code Debugging

`.vscode/launch.json` supports debugging with:
```json
"env": { "MONGO_URI": "mongodb://localhost:27017" }
```

---

## 💼 Dev Container (VS Code)

1. Ensure Docker + Dev Containers extension is installed.
2. Open project → `Cmd+Shift+P` → “Dev Containers: Reopen in Container”
3. MongoDB and dependencies auto-mounted.

---

## 🧪 Seeder

Uncomment `commands.SeedTags()` in `main.go` to prefill sample tags:

```go
// commands.SeedTags()
```

---

## 📁 Folder Structure

```
kb-api-go/
├── config/         # MongoDB connection
├── controllers/    # Article & Tag handlers
├── models/         # Data models
├── routes/         # Route registrations
├── commands/       # Seeder utilities
├── .devcontainer/  # VS Code container support
├── .vscode/        # Debug tasks
├── Dockerfile
├── docker-compose.yml
├── api-tests.http
├── main.go
```

---

## 📜 License

MIT License © [thilina01](https://github.com/thilina01)