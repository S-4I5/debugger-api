LOCAL_BIN:=$(CURDIR)/bin

install-go-deps:
	go get github.com/go-chi/chi/v5
	go get github.com/go-chi/render
	go get github.com/ilyakaznacheev/cleanenv
	go get github.com/swaggo/http-swagger/v2
	go get github.com/swaggo/swag
	go get gopkg.in/kothar/brotli-go.v0

generate:
	mkdir -p docs/swagger
	make generate-swagger

generate-swagger:
	swag init -d cmd
