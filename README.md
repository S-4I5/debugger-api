# Debugger API

## 1. Description

CRUD API, which allows you to save free-form jsons and access them via given key with brotli compression.

## 2. ENVs

| Name      | Type     | Default value | Description |
|-----------|----------|---------------|-------------|
| HOST      | String   |               | Server host |
| PORT      | String   | 8080          | Servet port |
| BASE_PATH | String   | /api/v1       | Api prefix  |

## 3. Run
```
go run ./cmd/main.go
```

## 4. Deployment

### 4.1 Compose
```
docker-compose up
```

### 4.2 Build image
```
docker build -f ./deployment/debugger-api.dockerfile .
```

## 5. Swagger
```
http://localhost:8080/swagger/index.html#/
```

