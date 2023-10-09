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
		{Name: "number", Type: field.TypeInt},
		{Name: "view_url", Type: field.TypeString, Unique: true},
		{Name: "download_url", Type: field.TypeString, Unique: true},
		{Name: "file_name", Type: field.TypeString},
		{Name: "file_size", Type: field.TypeInt},
		{Name: "resolution", Type: field.TypeString, Nullable: true},
		{Name: "video_codec", Type: field.TypeString, Nullable: true},
		{Name: "audio_codec", Type: field.TypeString, Nullable: true},
		{Name: "anime_id", Type: field.TypeInt},
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
				Columns:    []*schema.Column{EpisodesColumns[9]},
				RefColumns: []*schema.Column{AnimesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "episodes_release_groups_episodes",
				Columns:    []*schema.Column{EpisodesColumns[10]},
				RefColumns: []*schema.Column{ReleaseGroupsColumns[0]},
				OnDelete:   schema.SetNull,
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
		ReleaseGroupsTable,
	}
)

func init() {
	EpisodesTable.ForeignKeys[0].RefTable = AnimesTable
	EpisodesTable.ForeignKeys[1].RefTable = ReleaseGroupsTable
}
