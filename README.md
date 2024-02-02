# go-health-check-api

This is a simple health check API that can be used to check the health of an server. It is implemented with Gin web framework in Golang, PostgreSQL for the db and Gorm for ORM.

## Install & update packages

```sh
go get -u github.com/gin-gonic/gin

go get -u gorm.io/gorm

go get -u gorm.io/driver/postgres

go get -u github.com/joho/godotenv
```

## How to test the API

```sh
# Test GET method - should return 200 OK
curl -4 -vvvv http://localhost:8080/healthz

# Test payload not allowed - should return 400 Bad Request
curl -4 -d "data=payload" -vvvv http://localhost:8080/healthz

# Test PUT method not allowed - should return 405 Method Not Allowed
curl -4 -vvvv -XPUT http://localhost:8080/healthz

# Test DB connection - Modify the dsn,
# Or kill the psql server:
sudo lsof -i :5432
sudo kill -15 PID
# should return 503 Service Unavailable when the DB is not available
curl -4 -vvvv http://localhost:8080/healthz
```