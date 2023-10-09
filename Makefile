.DEFAULT_GOAL := all

.PHONY: all
all:
	echo "ohai"

.PHONY: schema
schema: ENTITY := Episode
schema:
	go run -mod=mod entgo.io/ent/cmd/ent new $(ENTITY)

.PHONY: generate
generate:
	go generate ./ent

.PHONY: migrate
migrate:
	go run migration.go
