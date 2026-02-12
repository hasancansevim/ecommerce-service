# 🛍️ Go E-Commerce Service

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go&logoColor=white)](https://golang.org)
[![Echo](https://img.shields.io/badge/Echo-v4-3B5998)](https://echo.labstack.com)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A production-ready e-commerce backend API built with Go, following **Clean Architecture** principles. Features PostgreSQL, Redis, Elasticsearch, RabbitMQ, JWT authentication, and comprehensive test coverage.

---

## 📐 Architecture Overview

This project implements **Clean Architecture** (Hexagonal/Ports & Adapters) with clear separation of concerns. Dependencies point inward: outer layers depend on inner layers, never the reverse.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           PRESENTATION LAYER                                  │
│  Controllers (HTTP)  │  Request/Response DTOs  │  Middleware (Auth, Errors)   │
└─────────────────────────────────────────────────────────────────────────────┘
                                        │
                                        ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                            APPLICATION LAYER                                  │
│  Services (Use Cases)  │  DTOs  │  Business Rules  │  Interfaces (Ports)      │
└─────────────────────────────────────────────────────────────────────────────┘
                                        │
                                        ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                              DOMAIN LAYER                                     │
│  Entities (Product, Order, User, Cart...)  │  Pure business logic             │
└─────────────────────────────────────────────────────────────────────────────┘
                                        │
                                        ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                         INFRASTRUCTURE LAYER                                  │
│  Repositories (DB)  │  PostgreSQL  │  Redis  │  Elasticsearch  │  RabbitMQ   │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 🏗️ Project Structure (Clean Architecture)

```
go-ecommerce-service/
│
├── main.go                    # Application entry point, wiring, graceful shutdown
├── config/                    # Configuration (env-based)
│   └── config.go              # Config struct, Load() from env
│
├── controller/                # PRESENTATION - HTTP handlers
│   ├── base_controller.go     # ParseIdParam, Success, BadRequest, Created
│   ├── auth_controller.go     # Register, Login (public)
│   ├── product_controller.go  # Product CRUD, search, sync
│   ├── order_controller.go    # Order CRUD, status, total price
│   ├── cart_controller.go     # Cart operations
│   ├── cart_item_controller.go
│   ├── order_item_controller.go
│   ├── category_controller.go
│   ├── store_controller.go
│   ├── user_controller.go
│   ├── request/               # HTTP request DTOs, ToModel() mappers
│   │   └── request.go
│   └── response/              # API response wrapper
│       └── response.go
│
├── domain/                    # DOMAIN - Pure entities
│   ├── product.go
│   ├── order.go
│   ├── order_item.go
│   ├── cart.go
│   ├── cart_item.go
│   ├── user.go
│   ├── category.go
│   └── store.go
│
├── service/                   # APPLICATION - Use cases, business logic
│   ├── product_service.go     # IProductService, Redis cache, ES search
│   ├── order_service.go       # IOrderService, RabbitMQ events
│   ├── auth_service.go        # AuthService (Register, Login, JWT)
│   ├── cart_service.go
│   ├── cart_item_service.go
│   ├── order_item_service.go
│   ├── category_service.go
│   ├── store_service.go
│   ├── user_service.go
│   ├── jwt_service.go
│   ├── interface/             # Port interfaces (inversion)
│   │   ├── auth.go
│   │   └── jwt_manager.go
│   ├── model/                 # Internal request models
│   │   └── model.go
│   ├── validation/            # Functional validator (legacy, rules preferred)
│   └── worker/                # Background worker (consumes RabbitMQ)
│       └── order_worker.go
│
├── persistence/               # INFRASTRUCTURE - Data access
│   ├── product_repository.go  # IProductRepository, PostgreSQL + Elasticsearch
│   ├── order_repository.go
│   ├── cart_repository.go
│   ├── cart_item_repository.go
│   ├── order_item_repository.go
│   ├── user_repository.go
│   ├── category_repository.go
│   ├── store_repository.go
│   ├── common/                # Errors, constants
│   └── helper/                # GenericScanner[T], scan functions
│       ├── generic_scanner.go
│       ├── scan_functions.go
│       └── interfaces/
│
├── internal/                  # Private application packages
│   ├── dto/                   # Data Transfer Objects
│   │   ├── product_dto.go
│   │   ├── order_dto.go
│   │   ├── cart_dto.go
│   │   └── ...
│   ├── rules/                 # Business validation rules
│   │   ├── base_rules.go      # ValidateStructure (go-playground/validator)
│   │   ├── product_rules.go   # Price >= 0, Discount >= 0
│   │   ├── order_rules.go
│   │   └── ...
│   ├── auth/                  # Password hashing (bcrypt)
│   │   └── password.go
│   └── jwt/                   # JWT token generation/validation
│       └── helper.go
│
├── infrastructure/            # External systems
│   ├── elasticsearch/
│   │   └── client.go          # Elasticsearch client, retry logic
│   └── rabbitmq/
│       └── client.go          # IRabbitMQClient, Publish, queue declaration
│
├── common/                    # Shared infra utilities
│   └── postgresql/
│       └── connection.go      # pgxpool connection
│
├── pkg/                       # Reusable packages
│   ├── errors/                # AppError, NewBadRequest, NewNotFound...
│   ├── logger/                # Zerolog initialization
│   ├── middleware/            # AuthMiddleware, CustomHTTPErrorHandler
│   ├── util/                  # GenerateSlug, GenerateUniqueSlug
│   └── validation/            # ValidateStruct (go-playground/validator)
│
├── docs/                      # Swagger documentation
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
│
├── test/                      # Tests
│   ├── controller/            # Product, Order controller tests
│   ├── unit/service/          # Product, Order service unit tests
│   ├── mock/                  # Mocks (gomock)
│   │   ├── repository/
│   │   ├── service/
│   │   └── infrastructure/
│   └── scripts/
│       └── test_db.ps1        # Test PostgreSQL setup
│
├── docker-compose.yml         # App + PostgreSQL + Redis + RabbitMQ + Elasticsearch
├── Dockerfile                 # Multi-stage build
├── init.sql                   # Database schema
└── manage.ps1                 # Scripts: run, up, down, swagger, infra
```

---

## 🔄 Data Flow (Request → Response)

### Example: Create Order

```
1. HTTP POST /api/v1/orders
   └─ AuthMiddleware validates JWT → userId
   
2. OrderController.CreateOrder
   └─ Bind AddOrderRequest → ToModel() → dto.CreateOrderRequest
   
3. OrderService.CreateOrder
   └─ OrderRules.ValidateStructure (total_price > 0)
   └─ OrderRepository.CreateOrder (PostgreSQL)
   └─ Publish to RabbitMQ "order_created_queue"
   
4. OrderWorker (background)
   └─ Consume from "order_created_queue"
   └─ UpdateOrderStatus(orderId, "Shipped")
   
5. Response: OrderResponse JSON
```

### Example: Get Product by ID (with Redis cache)

```
1. HTTP GET /api/v1/products/:id
   
2. ProductController.GetProductById
   └─ ParseIdParam("id")
   
3. ProductService.GetProductById
   └─ Redis Get "product:{id}" → cache hit? return
   └─ ProductRepository.GetProductById (PostgreSQL)
   └─ Redis Set "product:{id}", 10min TTL
   
4. Response: ProductResponse JSON
```

### Example: Search Products (Elasticsearch)

```
1. HTTP GET /api/v1/products/search?q=laptop
   
2. ProductController.SearchProducts
   
3. ProductService.SearchProducts
   └─ ProductRepository.SearchProducts
      └─ Elasticsearch: multi_match (fuzzy), wildcard (name, slug)
   
4. Response: []ProductResponse
```

---

## 🛠️ Technology Stack

| Layer | Technology |
|-------|------------|
| **Language** | Go 1.24 |
| **Web Framework** | Echo v4 |
| **Database** | PostgreSQL 14 (pgx/v4) |
| **Cache** | Redis (go-redis) |
| **Search** | Elasticsearch 8 |
| **Message Queue** | RabbitMQ (amqp091-go) |
| **Auth** | JWT (golang-jwt/jwt/v4), bcrypt |
| **Validation** | go-playground/validator/v10 |
| **Logging** | Zerolog |
| **Config** | envconfig (12-factor) |
| **API Docs** | Swagger (swaggo) |
| **Testing** | testify, gomock, redismock |

---

## 📦 Domain Entities

| Entity | Key Fields |
|--------|------------|
| **Product** | Id, Name, Slug, Price, BasePrice, Discount, StockQuantity, StoreId, CategoryId |
| **Order** | Id, UserId, TotalPrice, Status, CreatedAt, UpdatedAt |
| **OrderItem** | OrderId, ProductId, Quantity, Price |
| **Cart** | Id, UserId |
| **CartItem** | CartId, ProductId, Quantity |
| **User** | Id, FirstName, LastName, Email, PasswordHash |
| **Category** | Id, Name, Description, IsActive |
| **Store** | Id, Name, Slug, Description, ContactEmail |

---

## 🚀 API Endpoints

### Public (no auth)
| Method | Path | Description |
|--------|------|-------------|
| POST | `/api/v1/auth/register` | Register user |
| POST | `/api/v1/auth/login` | Login, returns JWT |
| GET | `/api/v1/products` | List all products |
| GET | `/api/v1/products/search?q=` | Search products (Elasticsearch) |
| GET | `/api/v1/products/:id` | Get product by ID |

### Protected (Bearer token)
| Method | Path | Description |
|--------|------|-------------|
| POST | `/api/v1/products` | Add product |
| PUT | `/api/v1/products/:id` | Update product |
| DELETE | `/api/v1/products/:id` | Delete product |
| POST | `/api/v1/products/sync` | Sync products to Elasticsearch |
| POST | `/api/v1/orders` | Create order |
| GET | `/api/v1/orders/:id` | Get order |
| GET | `/api/v1/orders/get-orders-by-user-id?user_id=` | Orders by user |
| GET | `/api/v1/orders/get-all-orders` | All orders |
| PUT | `/api/v1/orders/update-order-status/:id?status=` | Update status |
| PUT | `/api/v1/orders/:id?total_price=` | Update total price |
| DELETE | `/api/v1/orders/:id` | Delete order |
| GET | `/api/v1/orders/?status=` | Orders by status |
| ... | Cart, CartItem, OrderItem, Category, Store, User | CRUD operations |

**Swagger UI:** `http://localhost:8080/swagger/index.html`

---

## ⚙️ Configuration (Environment Variables)

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | localhost | PostgreSQL host |
| `DB_PORT` | 6432 | PostgreSQL port |
| `DB_USERNAME` | postgres | DB user |
| `DB_PASSWORD` | 123456 | DB password |
| `DB_NAME` | ecommerce | Database name |
| `DB_MAX_CONNECTIONS` | 10 | Connection pool size |
| `REDIS_HOST` | localhost | Redis host |
| `REDIS_PORT` | 6379 | Redis port |
| `RABBITMQ_HOST` | localhost | RabbitMQ host |
| `RABBITMQ_PORT` | 5672 | RabbitMQ port |
| `ELASTICSEARCH_HOST` | localhost | Elasticsearch host |
| `ELASTICSEARCH_PORT` | 9200 | Elasticsearch port |
| `SERVER_PORT` | 8080 | HTTP server port |
| `JWT_SECRET` | akaimpkminik3 | JWT signing secret |
| `JWT_DURATION` | 24h | Token expiry |

> **Note:** In `docker-compose.yml`, `DB_USER` is set but config expects `DB_USERNAME`. For Docker, add `DB_USERNAME=postgres` or align variable names.

---

## 🏃 Running the Project

### Prerequisites
- Go 1.24+
- PostgreSQL, Redis, RabbitMQ, Elasticsearch (or use Docker)

### Local run (with Docker infra)
```powershell
# Start infrastructure only
.\manage.ps1 infra

# Run app (Swagger init + go run)
.\manage.ps1 run
```

### Full stack with Docker
```powershell
.\manage.ps1 up
# or detached: .\manage.ps1 up-d
```

### Manual
```bash
swag init
go run .
```

### Build
```bash
go build -o main .
```

---

## 🧪 Testing

```bash
# Unit tests
go test ./test/unit/... -v

# Controller tests
go test ./test/controller/... -v

# All tests
go test ./test/... -v
```

**Test coverage:**
- Product service (Redis cache, validation)
- Order service (RabbitMQ publish, validation)
- Product controller (suite)
- Order controller (suite)

**Mock generation (if interfaces change):**
```bash
mockgen -source=persistence/product_repository.go -destination=test/mock/repository/product_repository.go -package=repository
mockgen -source=persistence/order_repository.go -destination=test/mock/repository/order_repository.go -package=repository
mockgen -source=infrastructure/rabbitmq/client.go -destination=test/mock/infrastructure/rabbitmq_mock.go -package=mock_infra
```

---

## 🐳 Docker

### Services
| Service | Port | Description |
|---------|------|-------------|
| app | 8080 | Go API |
| postgres | 5432 | PostgreSQL |
| redis | 6379 | Redis |
| rabbitmq | 5672 (AMQP), 15672 (Management UI) | RabbitMQ |
| elasticsearch | 9200 | Elasticsearch |

### Volumes
- `postgres_data` – PostgreSQL data
- `elastic_data` – Elasticsearch indices

---

## 📋 Event-Driven Flow (Order → RabbitMQ → Worker)

```
OrderService.CreateOrder
    │
    ├─► OrderRepository.CreateOrder (PostgreSQL)
    │
    └─► RabbitMQ.Publish("order_created_queue", payload)
            │
            ▼
        OrderWorker.Start()
            │
            ├─► Consume from "order_created_queue"
            │
            └─► OrderRepository.UpdateOrderStatus(orderId, "Shipped")
```

Payload: `{"order_id": 1, "user_id": 1, "message": "...", "total": 15000}`

---

## 📄 License

MIT License – see [LICENSE](LICENSE) for details.

---

## 👤 Author

**Hasan Can Sevim**
