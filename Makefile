LOCAL_BIN:=$(CURDIR)/bin

install-go-deps:
	go get github.com/go-chi/chi/v5
	go get github.com/go-chi/render
	go get github.com/ilyakaznacheev/cleanenv
	go get github.com/swaggo/http-swagger/v2
	go get github.com/swaggo/swag
	go get gopkg.in/kothar/brotli-go.v0
	go get github.com/go-redis/redis/v8
	go get github.com/pressly/goose/v3/cmd/goose@latest
	go get github.com/jackc/pgx/v5/stdlib
	go get github.com/Masterminds/squirrel
	go get github.com/stretchr/testify/require
	go get github.com/go-playground/validator/v10

generate-swagger:
	swag init -d cmd
