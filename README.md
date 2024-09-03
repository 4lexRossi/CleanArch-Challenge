# Order Service

## Setup

### Prerequisites

- Docker
- Docker Compose

### Running the Application

1. **Start the Database:**

  ```bash
   docker-compose -f docker/docker-compose.yaml up
  ```
2. **Run the Application:**

  ```bash
   go run internal/main.go wire_gen.go
  ```

  > The REST API will be available at http://localhost:8080/order.

  > The GRPC service will be available at localhost:50051.

  > The GraphQL endpoint will be available at http://localhost:8080/graphql.

### Testing

  > Use the api/api.http file to test the different endpoints.

### Migrations

> To apply the database migrations, connect to the PostgreSQL container and run:
  ```bash
   psql -U user -d orderdb -f migrations/001_create_orders_table.sql
  ```