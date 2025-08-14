# Go + Gin Backend Roadmap (Intermediate - 1 Week)

## **Day 1: Project Setup & Architecture**

-   [X] Khởi tạo dự án Gin (`go mod init` + cấu trúc thư mục Clean
    Architecture)
-   [X] Cấu hình `.env` & load config (Viper)
-   [X] Tích hợp logger (Zap hoặc Logrus)
-   [X] Kết nối PostgreSQL (GORM / sqlc)
-   [ ] Cấu hình Docker Compose cho Postgres & API

## **Day 2: Routing & Middleware**

-   [X] Định nghĩa route RESTful trong Gin
-   [X] Xây dựng middleware logging, recovery, request ID
-   [ ] CORS & rate-limiting
-   [ ] JWT Auth middleware
-   [ ] Unit test cho middleware

## **Day 3: Database Layer & Transaction**

-   [ ] Repository pattern cho CRUD
-   [ ] Transaction management (GORM `db.Transaction`, sqlc + manual
    transaction)
-   [ ] Index, constraint & migration (golang-migrate)
-   [ ] Test transaction rollback / commit

## **Day 4: Concurrency & Background Jobs**

-   [ ] Goroutines + WaitGroup
-   [ ] Context cancellation (timeout, deadline)
-   [ ] Worker pool pattern
-   [ ] **_Channel patterns (fan-in, fan-out)_**
-   [ ] **_Async job queue (RabbitMQ, Kafka, hoặc Redis Streams)_**

## **Day 5: API Business Logic & Validation**

-   [ ] DTOs & binding request body/query params
-   [ ] Validation với go-playground/validator
-   [ ] Service layer tách biệt
-   [ ] Error handling chuẩn REST API (HTTP status code + message)
-   [ ] Unit test cho service

## **Day 6: Observability & Metrics**

-   [ ] Tích hợp Prometheus metrics (ginprom)
-   [ ] Tracing với OpenTelemetry
-   [ ] Health check endpoint
-   [ ] Structured logging với correlation ID

## **Day 7: Deployment & Optimization**

-   [ ] Dockerfile multi-stage build
-   [ ] CI/CD pipeline (GitHub Actions hoặc GitLab CI)
-   [ ] Benchmark với `go test -bench`
-   [ ] Optimize DB queries
-   [ ] Triển khai lên cloud (Render, Railway, AWS)

------------------------------------------------------------------------

📚 **Tài liệu tham khảo:** - Gin: https://gin-gonic.com/docs/ - sqlc:
https://docs.sqlc.dev/ - GORM: https://gorm.io/docs/ - OpenTelemetry Go:
https://opentelemetry.io/docs/instrumentation/go/ - Golang-Migrate:
https://github.com/golang-migrate/migrate
