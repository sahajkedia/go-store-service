# Go Store Service

A small production-minded **Go GraphQL API** to manage stores and products, built with:

- **Go** (1.22)
- **GraphQL** (gqlgen)
- **Postgres** (persistence)
- **Redis** (caching)
- **Docker & Docker Compose** (local orchestration)

---

##  Features

- **GraphQL API** for managing `Store` and `Product` entities.
- **Postgres** for persistent storage.
- **Redis caching** for faster repeated queries.
- **Clean architecture** separation: resolvers, services, and repositories.
- **Docker Compose** for one-command setup.
- **GraphQL Playground** enabled at `http://localhost:8080`.
- **Unit tests** for store and product services.

---

##  Getting Started

### Prerequisites

- Docker & Docker Compose installed.
- Go 1.22+ (if you want to run locally without Docker).

### Clone the repo

```bash
git clone https://github.com/your-username/go-store-service.git
cd go-store-service
```

### Run with Docker Compose

```bash
docker-compose up --build
```

This starts Postgres, Redis, and the API on port `8080`.

### Access Playground

Open [http://localhost:8080](http://localhost:8080) in your browser.

---

## 🔧 Example Queries

### Create a Store

```graphql
mutation {
	createStore(name: "My First Store") {
		id
		name
	}
}
```

### Add a Product

```graphql
mutation {
	addProduct(storeId: "1", name: "Sticker Pack", price: 9.99) {
		id
		name
		price
	}
}
```

### Query Products

```graphql
query {
	products(storeId: "1") {
		id
		name
		price
	}
}
```

---

##  Tests

Run unit tests with:

```bash
go test ./...
```

---

##  Why This Project

This project demonstrates a **scalable and maintainable** approach to building a commerce-style backend:

- **Separation of concerns** (models, resolvers, services).
- **Caching** for performance improvements.
- **Containerization** for easy local development.
- **GraphQL flexibility** for frontends.
- **Unit tests** for confidence in correctness.

It's not just functional — it's designed with **production-readiness** and **developer experience** in mind.

---

## License

MIT
