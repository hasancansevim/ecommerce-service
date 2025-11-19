# 🛍️ GO E-COMMERCE BACKEND API

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen.svg)](https://github.com/yourusername/go-ecommerce-service)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

Kurumsal seviyede Go dilinde geliştirilmiş e-ticaret backend API'si. Clean Architecture, kapsamlı test suite ve modern development practices ile geliştirilmiştir.

## 🏗️ ARCHITECTURE
go-ecommerce-service/

├── 📁 common/ # Shared utilities

├── 📁 config/ # Configuration management

├── 📁 controller/ # HTTP handlers (Echo framework)

├── 📁 domain/ # Business entities

├── 📁 internal/ # Private packages

├── 📁 middleware/ # Middlewares

├── 📁 persistence/ # Repository pattern

├── 📁 pkg/ # Shared libraries

├── 📁 service/ # Business logic layer

├── 📁 test/ # Comprehensive test suite

        └── 📁 integration/ # Database integration tests
        
        └── 📁 unit/ # Business logic unit tests
        
        └── 📁 mocks/ # Mock implementations



## 🚀 IMPLEMENTED FEATURES

**Authentication & Authorization**
- [x] JWT token based authentication
- [x] Password hashing with bcrypt
- [x] Register/Login endpoints
- [x] Auth middleware for protected routes

**Business Domains**
- [x] User Management: Registration, login, profiles
- [x] Product Management: CRUD operations, store filtering
- [x] Order Management: Order creation, status tracking
- [x] Cart Management: Shopping cart operations
- [x] OrderItem Management: Order line items

**Technical Excellence**
- [x] Clean Architecture implementation
- [x] Comprehensive unit & integration testing
- [x] Custom error handling & validation
- [x] PostgreSQL + pgx integration
- [x] Echo web framework
- [x] Mock-based testing infrastructure

**Test Coverage - %100 SUCCESS**
- ✅ Auth Controller Tests
- ✅ Product Controller Tests  
- ✅ Order Controller Tests
- ✅ Cart Controller Tests
- ✅ CartItem Controller Tests
- ✅ OrderItem Controller Tests

## 🛠️ TECHNOLOGY STACK

- **Language**: Go 1.21+
- **Framework**: Echo v4
- **Database**: PostgreSQL + pgx/v4
- **Authentication**: JWT + bcrypt
- **Testing**: Testify + custom mocks
- **Validation**: Custom functional-options validator

## 📄 Licence

  [MIT](https://choosealicense.com/licenses/mit/)

