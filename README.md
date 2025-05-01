# 📦 Datastore Server (Go + PostgreSQL)

This is a scalable Go server that connects to a PostgreSQL database and provides two endpoints:

1. ✅ `/health` — Health check for DB connection  
2. 📥 `/run-query` — Accepts a SQL query in the request body and returns the result

---

## 🚀 Features

- ⚙️ Environment-based configuration using .env
- 🗄 PostgreSQL connection pooling with pgx
- 🔐 Clean folder structure: internal/config, db, handlers
- 📡 RESTful API via Gorilla Mux
- 🧪 Modular for easy testing and expansion
