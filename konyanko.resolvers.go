package konyanko

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/eiri/konyanko/ent"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Animes is the resolver for the animes field.
func (r *queryResolver) Animes(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.AnimeOrder) (*ent.AnimeConnection, error) {
	return r.client.Anime.Query().Paginate(ctx, after, first, before, last,
		ent.WithAnimeOrder(orderBy),
	)
}

// Episodes is the resolver for the episodes field.
func (r *queryResolver) Episodes(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*ent.EpisodeOrder) (*ent.EpisodeConnection, error) {
	return r.client.Episode.Query().Paginate(ctx, after, first, before, last,
		ent.WithEpisodeOrder(orderBy),
	)
}

// Items is the resolver for the items field.
func (r *queryResolver) Items(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*ent.ItemOrder) (*ent.ItemConnection, error) {
	return r.client.Item.Query().Paginate(ctx, after, first, before, last,
		ent.WithItemOrder(orderBy),
	)
}

// ReleaseGroups is the resolver for the releaseGroups field.
func (r *queryResolver) ReleaseGroups(ctx context.Context) ([]*ent.ReleaseGroup, error) {
	return r.client.ReleaseGroup.Query().All(ctx)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
