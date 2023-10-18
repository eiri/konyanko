// Code generated by ent, DO NOT EDIT.

package item

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/eiri/konyanko/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Item {
	return predicate.Item(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Item {
	return predicate.Item(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Item {
	return predicate.Item(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Item {
	return predicate.Item(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Item {
	return predicate.Item(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Item {
	return predicate.Item(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Item {
	return predicate.Item(sql.FieldLTE(FieldID, id))
}

// ViewURL applies equality check predicate on the "view_url" field. It's identical to ViewURLEQ.
func ViewURL(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldViewURL, v))
}

// DownloadURL applies equality check predicate on the "download_url" field. It's identical to DownloadURLEQ.
func DownloadURL(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldDownloadURL, v))
}

// FileName applies equality check predicate on the "file_name" field. It's identical to FileNameEQ.
func FileName(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldFileName, v))
}

// FileSize applies equality check predicate on the "file_size" field. It's identical to FileSizeEQ.
func FileSize(v int) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldFileSize, v))
}

// PublishDate applies equality check predicate on the "publish_date" field. It's identical to PublishDateEQ.
func PublishDate(v time.Time) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldPublishDate, v))
}

// ViewURLEQ applies the EQ predicate on the "view_url" field.
func ViewURLEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldViewURL, v))
}

// ViewURLNEQ applies the NEQ predicate on the "view_url" field.
func ViewURLNEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldNEQ(FieldViewURL, v))
}

// ViewURLIn applies the In predicate on the "view_url" field.
func ViewURLIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldIn(FieldViewURL, vs...))
}

// ViewURLNotIn applies the NotIn predicate on the "view_url" field.
func ViewURLNotIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldNotIn(FieldViewURL, vs...))
}

// ViewURLGT applies the GT predicate on the "view_url" field.
func ViewURLGT(v string) predicate.Item {
	return predicate.Item(sql.FieldGT(FieldViewURL, v))
}

// ViewURLGTE applies the GTE predicate on the "view_url" field.
func ViewURLGTE(v string) predicate.Item {
	return predicate.Item(sql.FieldGTE(FieldViewURL, v))
}

// ViewURLLT applies the LT predicate on the "view_url" field.
func ViewURLLT(v string) predicate.Item {
	return predicate.Item(sql.FieldLT(FieldViewURL, v))
}

// ViewURLLTE applies the LTE predicate on the "view_url" field.
func ViewURLLTE(v string) predicate.Item {
	return predicate.Item(sql.FieldLTE(FieldViewURL, v))
}

// ViewURLContains applies the Contains predicate on the "view_url" field.
func ViewURLContains(v string) predicate.Item {
	return predicate.Item(sql.FieldContains(FieldViewURL, v))
}

// ViewURLHasPrefix applies the HasPrefix predicate on the "view_url" field.
func ViewURLHasPrefix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasPrefix(FieldViewURL, v))
}

// ViewURLHasSuffix applies the HasSuffix predicate on the "view_url" field.
func ViewURLHasSuffix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasSuffix(FieldViewURL, v))
}

// ViewURLEqualFold applies the EqualFold predicate on the "view_url" field.
func ViewURLEqualFold(v string) predicate.Item {
	return predicate.Item(sql.FieldEqualFold(FieldViewURL, v))
}

// ViewURLContainsFold applies the ContainsFold predicate on the "view_url" field.
func ViewURLContainsFold(v string) predicate.Item {
	return predicate.Item(sql.FieldContainsFold(FieldViewURL, v))
}

// DownloadURLEQ applies the EQ predicate on the "download_url" field.
func DownloadURLEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldDownloadURL, v))
}

// DownloadURLNEQ applies the NEQ predicate on the "download_url" field.
func DownloadURLNEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldNEQ(FieldDownloadURL, v))
}

// DownloadURLIn applies the In predicate on the "download_url" field.
func DownloadURLIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldIn(FieldDownloadURL, vs...))
}

// DownloadURLNotIn applies the NotIn predicate on the "download_url" field.
func DownloadURLNotIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldNotIn(FieldDownloadURL, vs...))
}

// DownloadURLGT applies the GT predicate on the "download_url" field.
func DownloadURLGT(v string) predicate.Item {
	return predicate.Item(sql.FieldGT(FieldDownloadURL, v))
}

// DownloadURLGTE applies the GTE predicate on the "download_url" field.
func DownloadURLGTE(v string) predicate.Item {
	return predicate.Item(sql.FieldGTE(FieldDownloadURL, v))
}

// DownloadURLLT applies the LT predicate on the "download_url" field.
func DownloadURLLT(v string) predicate.Item {
	return predicate.Item(sql.FieldLT(FieldDownloadURL, v))
}

// DownloadURLLTE applies the LTE predicate on the "download_url" field.
func DownloadURLLTE(v string) predicate.Item {
	return predicate.Item(sql.FieldLTE(FieldDownloadURL, v))
}

// DownloadURLContains applies the Contains predicate on the "download_url" field.
func DownloadURLContains(v string) predicate.Item {
	return predicate.Item(sql.FieldContains(FieldDownloadURL, v))
}

// DownloadURLHasPrefix applies the HasPrefix predicate on the "download_url" field.
func DownloadURLHasPrefix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasPrefix(FieldDownloadURL, v))
}

// DownloadURLHasSuffix applies the HasSuffix predicate on the "download_url" field.
func DownloadURLHasSuffix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasSuffix(FieldDownloadURL, v))
}

// DownloadURLEqualFold applies the EqualFold predicate on the "download_url" field.
func DownloadURLEqualFold(v string) predicate.Item {
	return predicate.Item(sql.FieldEqualFold(FieldDownloadURL, v))
}

// DownloadURLContainsFold applies the ContainsFold predicate on the "download_url" field.
func DownloadURLContainsFold(v string) predicate.Item {
	return predicate.Item(sql.FieldContainsFold(FieldDownloadURL, v))
}

// FileNameEQ applies the EQ predicate on the "file_name" field.
func FileNameEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldFileName, v))
}

// FileNameNEQ applies the NEQ predicate on the "file_name" field.
func FileNameNEQ(v string) predicate.Item {
	return predicate.Item(sql.FieldNEQ(FieldFileName, v))
}

// FileNameIn applies the In predicate on the "file_name" field.
func FileNameIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldIn(FieldFileName, vs...))
}

// FileNameNotIn applies the NotIn predicate on the "file_name" field.
func FileNameNotIn(vs ...string) predicate.Item {
	return predicate.Item(sql.FieldNotIn(FieldFileName, vs...))
}

// FileNameGT applies the GT predicate on the "file_name" field.
func FileNameGT(v string) predicate.Item {
	return predicate.Item(sql.FieldGT(FieldFileName, v))
}

// FileNameGTE applies the GTE predicate on the "file_name" field.
func FileNameGTE(v string) predicate.Item {
	return predicate.Item(sql.FieldGTE(FieldFileName, v))
}

// FileNameLT applies the LT predicate on the "file_name" field.
func FileNameLT(v string) predicate.Item {
	return predicate.Item(sql.FieldLT(FieldFileName, v))
}

// FileNameLTE applies the LTE predicate on the "file_name" field.
func FileNameLTE(v string) predicate.Item {
	return predicate.Item(sql.FieldLTE(FieldFileName, v))
}

// FileNameContains applies the Contains predicate on the "file_name" field.
func FileNameContains(v string) predicate.Item {
	return predicate.Item(sql.FieldContains(FieldFileName, v))
}

// FileNameHasPrefix applies the HasPrefix predicate on the "file_name" field.
func FileNameHasPrefix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasPrefix(FieldFileName, v))
}

// FileNameHasSuffix applies the HasSuffix predicate on the "file_name" field.
func FileNameHasSuffix(v string) predicate.Item {
	return predicate.Item(sql.FieldHasSuffix(FieldFileName, v))
}

// FileNameEqualFold applies the EqualFold predicate on the "file_name" field.
func FileNameEqualFold(v string) predicate.Item {
	return predicate.Item(sql.FieldEqualFold(FieldFileName, v))
}

// FileNameContainsFold applies the ContainsFold predicate on the "file_name" field.
func FileNameContainsFold(v string) predicate.Item {
	return predicate.Item(sql.FieldContainsFold(FieldFileName, v))
}

// FileSizeEQ applies the EQ predicate on the "file_size" field.
func FileSizeEQ(v int) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldFileSize, v))
}

// FileSizeNEQ applies the NEQ predicate on the "file_size" field.
func FileSizeNEQ(v int) predicate.Item {
	return predicate.Item(sql.FieldNEQ(FieldFileSize, v))
}

// FileSizeIn applies the In predicate on the "file_size" field.
func FileSizeIn(vs ...int) predicate.Item {
	return predicate.Item(sql.FieldIn(FieldFileSize, vs...))
}

// FileSizeNotIn applies the NotIn predicate on the "file_size" field.
func FileSizeNotIn(vs ...int) predicate.Item {
	return predicate.Item(sql.FieldNotIn(FieldFileSize, vs...))
}

// FileSizeGT applies the GT predicate on the "file_size" field.
func FileSizeGT(v int) predicate.Item {
	return predicate.Item(sql.FieldGT(FieldFileSize, v))
}

// FileSizeGTE applies the GTE predicate on the "file_size" field.
func FileSizeGTE(v int) predicate.Item {
	return predicate.Item(sql.FieldGTE(FieldFileSize, v))
}

// FileSizeLT applies the LT predicate on the "file_size" field.
func FileSizeLT(v int) predicate.Item {
	return predicate.Item(sql.FieldLT(FieldFileSize, v))
}

// FileSizeLTE applies the LTE predicate on the "file_size" field.
func FileSizeLTE(v int) predicate.Item {
	return predicate.Item(sql.FieldLTE(FieldFileSize, v))
}

// PublishDateEQ applies the EQ predicate on the "publish_date" field.
func PublishDateEQ(v time.Time) predicate.Item {
	return predicate.Item(sql.FieldEQ(FieldPublishDate, v))
}

// PublishDateNEQ applies the NEQ predicate on the "publish_date" field.
func PublishDateNEQ(v time.Time) predicate.Item {
	return predicate.Item(sql.FieldNEQ(FieldPublishDate, v))
}

// PublishDateIn applies the In predicate on the "publish_date" field.
func PublishDateIn(vs ...time.Time) predicate.Item {
	return predicate.Item(sql.FieldIn(FieldPublishDate, vs...))
}

// PublishDateNotIn applies the NotIn predicate on the "publish_date" field.
func PublishDateNotIn(vs ...time.Time) predicate.Item {
	return predicate.Item(sql.FieldNotIn(FieldPublishDate, vs...))
}

// PublishDateGT applies the GT predicate on the "publish_date" field.
func PublishDateGT(v time.Time) predicate.Item {
	return predicate.Item(sql.FieldGT(FieldPublishDate, v))
}

// PublishDateGTE applies the GTE predicate on the "publish_date" field.
func PublishDateGTE(v time.Time) predicate.Item {
	return predicate.Item(sql.FieldGTE(FieldPublishDate, v))
}

// PublishDateLT applies the LT predicate on the "publish_date" field.
func PublishDateLT(v time.Time) predicate.Item {
	return predicate.Item(sql.FieldLT(FieldPublishDate, v))
}

// PublishDateLTE applies the LTE predicate on the "publish_date" field.
func PublishDateLTE(v time.Time) predicate.Item {
	return predicate.Item(sql.FieldLTE(FieldPublishDate, v))
}

// HasEpisode applies the HasEdge predicate on the "episode" edge.
func HasEpisode() predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, EpisodeTable, EpisodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEpisodeWith applies the HasEdge predicate on the "episode" edge with a given conditions (other predicates).
func HasEpisodeWith(preds ...predicate.Episode) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		step := newEpisodeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Item) predicate.Item {
	return predicate.Item(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Item) predicate.Item {
	return predicate.Item(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Item) predicate.Item {
	return predicate.Item(sql.NotPredicates(p))
}
