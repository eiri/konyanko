.DEFAULT_GOAL := all

PROJECT := konyanko
SRC := $(wildcard *.go ./ent/**/*.go)

.PHONY: all
all: build
	./$(PROJECT) -help

$(PROJECT): $(SRC)
	goimports -w main.go
	go build -o $(PROJECT) main.go

.PHONY: build
build: $(PROJECT)

.PHONY: migrate
migrate: $(PROJECT)
	./$< -c $@

nyaa.xml:
	curl -o $@ "https://nyaa.si/?page=rss&c=1_2"

.PHONY: import
import: $(PROJECT) nyaa.xml
	./$< -c $@

.PHONY: list
list: $(PROJECT)
	./$< -c $@

.PHONY: schema
schema: ENTITY := Episode
schema:
	go run -mod=mod entgo.io/ent/cmd/ent new $(ENTITY)

.PHONY: generate
generate:
	go generate ./ent
