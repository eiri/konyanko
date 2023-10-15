.DEFAULT_GOAL := all

PROJECT := konyanko
SRC := $(wildcard ./cmd/*.go ./ent/**/*.go)

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

.PHONY: list
list: $(PROJECT)
	@./$< $@

.PHONY: gron-list
gron-list:
	$(MAKE) list 2>&1 | tail -n +1 | jq -cs . | gron

.PHONY: server
server: $(PROJECT)
	./$< $@

.PHONY: schema
schema: ENTITY := Episode
schema:
	go run -mod=mod entgo.io/ent/cmd/ent new $(ENTITY)

.PHONY: cli-command
cli-command: COMMAND := migrate
cli-command:
# 	go run -mod=mod github.com/spf13/cobra-cli add $(COMMAND) -p 'readCmd'
	go run -mod=mod github.com/spf13/cobra-cli add $(COMMAND)

.PHONY: describe
describe:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema

.PHONY: generate
generate:
	go generate ./ent/...
