// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AnimesColumns holds the columns for the "animes" table.
	AnimesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Unique: true},
	}
	// AnimesTable holds the schema information for the "animes" table.
	AnimesTable = &schema.Table{
		Name:       "animes",
		Columns:    AnimesColumns,
		PrimaryKey: []*schema.Column{AnimesColumns[0]},
	}
	// EpisodesColumns holds the columns for the "episodes" table.
	EpisodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "episode_number", Type: field.TypeInt, Default: 0},
		{Name: "anime_season", Type: field.TypeInt, Default: 1},
		{Name: "resolution", Type: field.TypeString, Nullable: true},
		{Name: "video_codec", Type: field.TypeString, Nullable: true},
		{Name: "audio_codec", Type: field.TypeString, Nullable: true},
		{Name: "anime_id", Type: field.TypeInt},
		{Name: "item_id", Type: field.TypeInt, Unique: true},
		{Name: "release_group_id", Type: field.TypeInt, Nullable: true},
	}
	// EpisodesTable holds the schema information for the "episodes" table.
	EpisodesTable = &schema.Table{
		Name:       "episodes",
		Columns:    EpisodesColumns,
		PrimaryKey: []*schema.Column{EpisodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "episodes_animes_episodes",
				Columns:    []*schema.Column{EpisodesColumns[6]},
				RefColumns: []*schema.Column{AnimesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "episodes_items_episode",
				Columns:    []*schema.Column{EpisodesColumns[7]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "episodes_release_groups_episodes",
				Columns:    []*schema.Column{EpisodesColumns[8]},
				RefColumns: []*schema.Column{ReleaseGroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "episode_resolution",
				Unique:  false,
				Columns: []*schema.Column{EpisodesColumns[3]},
			},
		},
	}
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "view_url", Type: field.TypeString, Unique: true},
		{Name: "download_url", Type: field.TypeString, Unique: true},
		{Name: "file_name", Type: field.TypeString},
		{Name: "file_size", Type: field.TypeInt},
		{Name: "publish_date", Type: field.TypeTime},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "item_view_url",
				Unique:  false,
				Columns: []*schema.Column{ItemsColumns[1]},
			},
		},
	}
	// ReleaseGroupsColumns holds the columns for the "release_groups" table.
	ReleaseGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// ReleaseGroupsTable holds the schema information for the "release_groups" table.
	ReleaseGroupsTable = &schema.Table{
		Name:       "release_groups",
		Columns:    ReleaseGroupsColumns,
		PrimaryKey: []*schema.Column{ReleaseGroupsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AnimesTable,
		EpisodesTable,
		ItemsTable,
		ReleaseGroupsTable,
	}
)

func init() {
	EpisodesTable.ForeignKeys[0].RefTable = AnimesTable
	EpisodesTable.ForeignKeys[1].RefTable = ItemsTable
	EpisodesTable.ForeignKeys[2].RefTable = ReleaseGroupsTable
}
