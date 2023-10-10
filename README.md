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

## Schema

Anime
---
- title:str

ReleaseGroup
---
- name:str

Episode
---
- anime_id:id
- number:int
- view_url:str
- download_url:str
- release_group_id:id
- file_name:str
- file_size:int
- resolution:string
- video_codec:string
- audio_codec:string

Irregular
---
- view_url:str
- download_url:str
- file_name:str
- file_size:int

Preferred
---
- anime_id:id
- release_group_id:id
- resolution:string

## Dev runflow

```bash
$ make generate
$ make migrate
$ make import
$ make describe
```
