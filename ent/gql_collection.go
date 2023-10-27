// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/eiri/konyanko/ent/anime"
	"github.com/eiri/konyanko/ent/episode"
	"github.com/eiri/konyanko/ent/item"
	"github.com/eiri/konyanko/ent/releasegroup"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *AnimeQuery) CollectFields(ctx context.Context, satisfies ...string) (*AnimeQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return a, nil
	}
	if err := a.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AnimeQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(anime.Columns))
		selectedFields = []string{anime.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "episodes":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&EpisodeClient{config: a.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, episodeImplementors)...); err != nil {
				return err
			}
			a.WithNamedEpisodes(alias, func(wq *EpisodeQuery) {
				*wq = *query
			})
		case "title":
			if _, ok := fieldSeen[anime.FieldTitle]; !ok {
				selectedFields = append(selectedFields, anime.FieldTitle)
				fieldSeen[anime.FieldTitle] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		a.Select(selectedFields...)
	}
	return nil
}

type animePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AnimePaginateOption
}

func newAnimePaginateArgs(rv map[string]any) *animePaginateArgs {
	args := &animePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (e *EpisodeQuery) CollectFields(ctx context.Context, satisfies ...string) (*EpisodeQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return e, nil
	}
	if err := e.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return e, nil
}

func (e *EpisodeQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(episode.Columns))
		selectedFields = []string{episode.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "item":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ItemClient{config: e.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, itemImplementors)...); err != nil {
				return err
			}
			e.withItem = query
		case "title":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AnimeClient{config: e.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, animeImplementors)...); err != nil {
				return err
			}
			e.withTitle = query
		case "releaseGroup":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ReleaseGroupClient{config: e.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, releasegroupImplementors)...); err != nil {
				return err
			}
			e.withReleaseGroup = query
		case "episodeNumber":
			if _, ok := fieldSeen[episode.FieldEpisodeNumber]; !ok {
				selectedFields = append(selectedFields, episode.FieldEpisodeNumber)
				fieldSeen[episode.FieldEpisodeNumber] = struct{}{}
			}
		case "animeSeason":
			if _, ok := fieldSeen[episode.FieldAnimeSeason]; !ok {
				selectedFields = append(selectedFields, episode.FieldAnimeSeason)
				fieldSeen[episode.FieldAnimeSeason] = struct{}{}
			}
		case "resolution":
			if _, ok := fieldSeen[episode.FieldResolution]; !ok {
				selectedFields = append(selectedFields, episode.FieldResolution)
				fieldSeen[episode.FieldResolution] = struct{}{}
			}
		case "videoCodec":
			if _, ok := fieldSeen[episode.FieldVideoCodec]; !ok {
				selectedFields = append(selectedFields, episode.FieldVideoCodec)
				fieldSeen[episode.FieldVideoCodec] = struct{}{}
			}
		case "audioCodec":
			if _, ok := fieldSeen[episode.FieldAudioCodec]; !ok {
				selectedFields = append(selectedFields, episode.FieldAudioCodec)
				fieldSeen[episode.FieldAudioCodec] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		e.Select(selectedFields...)
	}
	return nil
}

type episodePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []EpisodePaginateOption
}

func newEpisodePaginateArgs(rv map[string]any) *episodePaginateArgs {
	args := &episodePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case []*EpisodeOrder:
			args.opts = append(args.opts, WithEpisodeOrder(v))
		case []any:
			var orders []*EpisodeOrder
			for i := range v {
				mv, ok := v[i].(map[string]any)
				if !ok {
					continue
				}
				var (
					err1, err2 error
					order      = &EpisodeOrder{Field: &EpisodeOrderField{}, Direction: entgql.OrderDirectionAsc}
				)
				if d, ok := mv[directionField]; ok {
					err1 = order.Direction.UnmarshalGQL(d)
				}
				if f, ok := mv[fieldField]; ok {
					err2 = order.Field.UnmarshalGQL(f)
				}
				if err1 == nil && err2 == nil {
					orders = append(orders, order)
				}
			}
			args.opts = append(args.opts, WithEpisodeOrder(orders))
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (i *ItemQuery) CollectFields(ctx context.Context, satisfies ...string) (*ItemQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return i, nil
	}
	if err := i.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return i, nil
}

func (i *ItemQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(item.Columns))
		selectedFields = []string{item.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "episode":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&EpisodeClient{config: i.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, episodeImplementors)...); err != nil {
				return err
			}
			i.withEpisode = query
		case "viewURL":
			if _, ok := fieldSeen[item.FieldViewURL]; !ok {
				selectedFields = append(selectedFields, item.FieldViewURL)
				fieldSeen[item.FieldViewURL] = struct{}{}
			}
		case "downloadURL":
			if _, ok := fieldSeen[item.FieldDownloadURL]; !ok {
				selectedFields = append(selectedFields, item.FieldDownloadURL)
				fieldSeen[item.FieldDownloadURL] = struct{}{}
			}
		case "fileName":
			if _, ok := fieldSeen[item.FieldFileName]; !ok {
				selectedFields = append(selectedFields, item.FieldFileName)
				fieldSeen[item.FieldFileName] = struct{}{}
			}
		case "fileSize":
			if _, ok := fieldSeen[item.FieldFileSize]; !ok {
				selectedFields = append(selectedFields, item.FieldFileSize)
				fieldSeen[item.FieldFileSize] = struct{}{}
			}
		case "publishDate":
			if _, ok := fieldSeen[item.FieldPublishDate]; !ok {
				selectedFields = append(selectedFields, item.FieldPublishDate)
				fieldSeen[item.FieldPublishDate] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		i.Select(selectedFields...)
	}
	return nil
}

type itemPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ItemPaginateOption
}

func newItemPaginateArgs(rv map[string]any) *itemPaginateArgs {
	args := &itemPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case []*ItemOrder:
			args.opts = append(args.opts, WithItemOrder(v))
		case []any:
			var orders []*ItemOrder
			for i := range v {
				mv, ok := v[i].(map[string]any)
				if !ok {
					continue
				}
				var (
					err1, err2 error
					order      = &ItemOrder{Field: &ItemOrderField{}, Direction: entgql.OrderDirectionAsc}
				)
				if d, ok := mv[directionField]; ok {
					err1 = order.Direction.UnmarshalGQL(d)
				}
				if f, ok := mv[fieldField]; ok {
					err2 = order.Field.UnmarshalGQL(f)
				}
				if err1 == nil && err2 == nil {
					orders = append(orders, order)
				}
			}
			args.opts = append(args.opts, WithItemOrder(orders))
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (rg *ReleaseGroupQuery) CollectFields(ctx context.Context, satisfies ...string) (*ReleaseGroupQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return rg, nil
	}
	if err := rg.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return rg, nil
}

func (rg *ReleaseGroupQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(releasegroup.Columns))
		selectedFields = []string{releasegroup.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "episodes":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&EpisodeClient{config: rg.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, episodeImplementors)...); err != nil {
				return err
			}
			rg.WithNamedEpisodes(alias, func(wq *EpisodeQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[releasegroup.FieldName]; !ok {
				selectedFields = append(selectedFields, releasegroup.FieldName)
				fieldSeen[releasegroup.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		rg.Select(selectedFields...)
	}
	return nil
}

type releasegroupPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ReleaseGroupPaginateOption
}

func newReleaseGroupPaginateArgs(rv map[string]any) *releasegroupPaginateArgs {
	args := &releasegroupPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
