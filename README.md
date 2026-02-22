# taskgraph

Async Graph Task Executor


## Tech stack

- Go
- Postgres
- Docker

### Libs & Tools
- `net/http` - Web server
- GORM - DB ORM
- goose - DB migrations
- swag - Swagger docs


## How to run

1. Copy `.env.sample` to `.env`:

```shell
cp .env.sample .env
```

2. Start DB container:

```shell
make run-db
```

3. Run migrations:

```shell
make goose-up
```

4. Run server:

```shell
make run
```

5. Access Swagger docs:

```
http://127.0.0.1:8000/swagger
```
