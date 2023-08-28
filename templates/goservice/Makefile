init:
	go mod init {{.Module}}
	go get -u github.com/spf13/cobra@latest
	go get -u github.com/spf13/viper
	mv env.sample .env
	mv gitignore .gitignore

toolsup:
	docker-compose -f tools/docker-compose.yml up -d

toolsdown:
	docker-compose -f tools/docker-compose.yml down

prepdev:
	brew install dbmate
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

migrate:
	dbmate up

apigen:
	rm -rf apigen
	mkdir apigen
	oapi-codegen -package apigen api/openapi.yaml > apigen/spec.gen.go

dbgen:
	rm -rf dbgen
	sqlc generate

genall: migrate prepdev apigen dbgen
	go mod tidy

dev:
	go run main.go serve

.PHONY: init dev prepdev apigen dbgen toolsup toolsdown