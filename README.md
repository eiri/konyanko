# Konyanko

A nyaa.si aggregator.

## Roadmap

- [x] Schema with entgo
- [x] RSS reader-exporter
- [x] Anime name parser
- [x] Iterative cycles of two above till completed
- [x] Openapi specs from ent schema
- [ ] Server in go and client in js from openapi specs
- [ ] API server
- [ ] Vue UI
- [ ] Crontab RSS puller in server
- [ ] Docker
- [ ] Goreleaser
- [ ] Myanimelist integration

## Schema

```mermaid
---
title: Konyanko schema
---
erDiagram
    Anime ||--o{ Episode : edge
    Anime {
        int    id
        string title
    }
    ReleaseGroup ||--o{ Episode : edge
    ReleaseGroup {
        int    id
        string name
    }
    Episode ||--|| Item : edge
    Episode {
        int id
        int episode_number
        int anime_season
        string resolution
        string video_codec
        string audio_codec
    }
    Item {
        int id
        string view_url
        string download_url
        string file_name
        int file_size
        datetime publish_date
    }
```

## OpenAPI specs

[openapi.json](https://petstore.swagger.io/?url=https://github.com/eiri/konyanko/blob/main/ent/openapi.json)

## Dev runflow

```bash
$ make generate
$ make migrate
$ make import
$ make describe
```
