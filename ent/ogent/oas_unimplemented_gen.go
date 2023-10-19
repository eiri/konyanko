// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateAnime implements createAnime operation.
//
// Creates a new Anime and persists it to storage.
//
// POST /api/v1/animes
func (UnimplementedHandler) CreateAnime(ctx context.Context, req *CreateAnimeReq) (r CreateAnimeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateEpisode implements createEpisode operation.
//
// Creates a new Episode and persists it to storage.
//
// POST /api/v1/episodes
func (UnimplementedHandler) CreateEpisode(ctx context.Context, req *CreateEpisodeReq) (r CreateEpisodeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateItem implements createItem operation.
//
// Creates a new Item and persists it to storage.
//
// POST /api/v1/items
func (UnimplementedHandler) CreateItem(ctx context.Context, req *CreateItemReq) (r CreateItemRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateReleaseGroup implements createReleaseGroup operation.
//
// Creates a new ReleaseGroup and persists it to storage.
//
// POST /api/v1/release-groups
func (UnimplementedHandler) CreateReleaseGroup(ctx context.Context, req *CreateReleaseGroupReq) (r CreateReleaseGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteAnime implements deleteAnime operation.
//
// Deletes the Anime with the requested ID.
//
// DELETE /api/v1/animes/{id}
func (UnimplementedHandler) DeleteAnime(ctx context.Context, params DeleteAnimeParams) (r DeleteAnimeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteEpisode implements deleteEpisode operation.
//
// Deletes the Episode with the requested ID.
//
// DELETE /api/v1/episodes/{id}
func (UnimplementedHandler) DeleteEpisode(ctx context.Context, params DeleteEpisodeParams) (r DeleteEpisodeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteItem implements deleteItem operation.
//
// Deletes the Item with the requested ID.
//
// DELETE /api/v1/items/{id}
func (UnimplementedHandler) DeleteItem(ctx context.Context, params DeleteItemParams) (r DeleteItemRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteReleaseGroup implements deleteReleaseGroup operation.
//
// Deletes the ReleaseGroup with the requested ID.
//
// DELETE /api/v1/release-groups/{id}
func (UnimplementedHandler) DeleteReleaseGroup(ctx context.Context, params DeleteReleaseGroupParams) (r DeleteReleaseGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListAnime implements listAnime operation.
//
// List Animes.
//
// GET /api/v1/animes
func (UnimplementedHandler) ListAnime(ctx context.Context, params ListAnimeParams) (r ListAnimeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListAnimeEpisodes implements listAnimeEpisodes operation.
//
// List attached Episodes.
//
// GET /api/v1/animes/{id}/episodes
func (UnimplementedHandler) ListAnimeEpisodes(ctx context.Context, params ListAnimeEpisodesParams) (r ListAnimeEpisodesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListEpisode implements listEpisode operation.
//
// List Episodes.
//
// GET /api/v1/episodes
func (UnimplementedHandler) ListEpisode(ctx context.Context, params ListEpisodeParams) (r ListEpisodeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListItem implements listItem operation.
//
// List Items.
//
// GET /api/v1/items
func (UnimplementedHandler) ListItem(ctx context.Context, params ListItemParams) (r ListItemRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListReleaseGroup implements listReleaseGroup operation.
//
// List ReleaseGroups.
//
// GET /api/v1/release-groups
func (UnimplementedHandler) ListReleaseGroup(ctx context.Context, params ListReleaseGroupParams) (r ListReleaseGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListReleaseGroupEpisodes implements listReleaseGroupEpisodes operation.
//
// List attached Episodes.
//
// GET /api/v1/release-groups/{id}/episodes
func (UnimplementedHandler) ListReleaseGroupEpisodes(ctx context.Context, params ListReleaseGroupEpisodesParams) (r ListReleaseGroupEpisodesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadAnime implements readAnime operation.
//
// Finds the Anime with the requested ID and returns it.
//
// GET /api/v1/animes/{id}
func (UnimplementedHandler) ReadAnime(ctx context.Context, params ReadAnimeParams) (r ReadAnimeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadEpisode implements readEpisode operation.
//
// Finds the Episode with the requested ID and returns it.
//
// GET /api/v1/episodes/{id}
func (UnimplementedHandler) ReadEpisode(ctx context.Context, params ReadEpisodeParams) (r ReadEpisodeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadEpisodeItem implements readEpisodeItem operation.
//
// Find the attached Item of the Episode with the given ID.
//
// GET /api/v1/episodes/{id}/item
func (UnimplementedHandler) ReadEpisodeItem(ctx context.Context, params ReadEpisodeItemParams) (r ReadEpisodeItemRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadEpisodeReleaseGroup implements readEpisodeReleaseGroup operation.
//
// Find the attached ReleaseGroup of the Episode with the given ID.
//
// GET /api/v1/episodes/{id}/release-group
func (UnimplementedHandler) ReadEpisodeReleaseGroup(ctx context.Context, params ReadEpisodeReleaseGroupParams) (r ReadEpisodeReleaseGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadEpisodeTitle implements readEpisodeTitle operation.
//
// Find the attached Anime of the Episode with the given ID.
//
// GET /api/v1/episodes/{id}/title
func (UnimplementedHandler) ReadEpisodeTitle(ctx context.Context, params ReadEpisodeTitleParams) (r ReadEpisodeTitleRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadItem implements readItem operation.
//
// Finds the Item with the requested ID and returns it.
//
// GET /api/v1/items/{id}
func (UnimplementedHandler) ReadItem(ctx context.Context, params ReadItemParams) (r ReadItemRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadItemEpisode implements readItemEpisode operation.
//
// Find the attached Episode of the Item with the given ID.
//
// GET /api/v1/items/{id}/episode
func (UnimplementedHandler) ReadItemEpisode(ctx context.Context, params ReadItemEpisodeParams) (r ReadItemEpisodeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadReleaseGroup implements readReleaseGroup operation.
//
// Finds the ReleaseGroup with the requested ID and returns it.
//
// GET /api/v1/release-groups/{id}
func (UnimplementedHandler) ReadReleaseGroup(ctx context.Context, params ReadReleaseGroupParams) (r ReadReleaseGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateAnime implements updateAnime operation.
//
// Updates a Anime and persists changes to storage.
//
// PATCH /api/v1/animes/{id}
func (UnimplementedHandler) UpdateAnime(ctx context.Context, req *UpdateAnimeReq, params UpdateAnimeParams) (r UpdateAnimeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateEpisode implements updateEpisode operation.
//
// Updates a Episode and persists changes to storage.
//
// PATCH /api/v1/episodes/{id}
func (UnimplementedHandler) UpdateEpisode(ctx context.Context, req *UpdateEpisodeReq, params UpdateEpisodeParams) (r UpdateEpisodeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateItem implements updateItem operation.
//
// Updates a Item and persists changes to storage.
//
// PATCH /api/v1/items/{id}
func (UnimplementedHandler) UpdateItem(ctx context.Context, req *UpdateItemReq, params UpdateItemParams) (r UpdateItemRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateReleaseGroup implements updateReleaseGroup operation.
//
// Updates a ReleaseGroup and persists changes to storage.
//
// PATCH /api/v1/release-groups/{id}
func (UnimplementedHandler) UpdateReleaseGroup(ctx context.Context, req *UpdateReleaseGroupReq, params UpdateReleaseGroupParams) (r UpdateReleaseGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}
