# Konyanko

A nyaa.si aggregator.

## Roadmap

- [x] Schema with entgo
- [x] RSS reader-exporter
- [x] Anime name parser
- [x] Iterative cycles of two above till completed
- [ ] Openapi specs from ent schema
- [ ] Server in go and client in js from openapi specs
- [ ] API server
- [ ] Vue UI
- [ ] Crontab RSS puller in server
- [ ] Docker
- [ ] Goreleaser
- [ ] Myanimelist integration

## Schama

Anime
---
- title:str
- preferred:bool

Subber
---
- name:str
- preferred:bool

Episode
---
- anime_id:id
- number:int
- url:str
- file_name:str
- subber_id:id
- resolution:enum
- codec:enum
- size:int
