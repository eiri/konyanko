.DEFAULT_GOAL := all

PROJECT := konyanko
SRC := $(wildcard ./*.go ./cmd/*.go ./ent/**/*.go)

.PHONY: all
all: build
	./$(PROJECT) --help

$(PROJECT): $(SRC)
	fd -e go -X goimports -w
	go build -o $(PROJECT) ./cmd/...

.PHONY: build
build: $(PROJECT)

.PHONY: migrate
migrate: $(PROJECT)
	./$< $@

nyaa.xml:
	curl -o $@ "https://nyaa.si/?page=rss&c=1_2"

.PHONY: import
import: $(PROJECT) nyaa.xml
	./$< $@ -f $(word 2,$^)

.PHONY: list-titles
list-titles: $(PROJECT)
	@./$< list titles

.PHONY: list-anime
list-anime: $(PROJECT)
	@./$< list anime --day 2023-10-10

.PHONY: list-anime-gron
list-anime-gron:
	$(MAKE) list-anime 2>&1 | tail -n +1 | gron

.PHONY: list-irregular
list-irregular: $(PROJECT)
	@./$< list irregular

.PHONY: server
server: $(PROJECT)
	./$< $@

.PHONY: schema
schema: ENTITY := Episode
schema:
	go run -mod=mod entgo.io/ent/cmd/ent new $(ENTITY)
	go mod tidy

.PHONY: graphql
graphql:
	go run -mod=mod github.com/99designs/gqlgen
	go mod tidy

.PHONY: cli-command
cli-command: COMMAND := queries
cli-command:
	go run -mod=mod github.com/spf13/cobra-cli add $(COMMAND) -p 'listCmd'
	go mod tidy

.PHONY: describe
describe:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema
	go mod tidy

.PHONY: generate
generate:
	go generate .
	go mod tidy
