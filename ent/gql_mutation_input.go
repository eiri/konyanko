// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateItemInput represents a mutation input for creating items.
type CreateItemInput struct {
	ViewURL     string
	DownloadURL string
	FileName    string
	FileSize    int
	PublishDate *time.Time
	EpisodeID   *int
}

// Mutate applies the CreateItemInput on the ItemMutation builder.
func (i *CreateItemInput) Mutate(m *ItemMutation) {
	m.SetViewURL(i.ViewURL)
	m.SetDownloadURL(i.DownloadURL)
	m.SetFileName(i.FileName)
	m.SetFileSize(i.FileSize)
	if v := i.PublishDate; v != nil {
		m.SetPublishDate(*v)
	}
	if v := i.EpisodeID; v != nil {
		m.SetEpisodeID(*v)
	}
}

// SetInput applies the change-set in the CreateItemInput on the ItemCreate builder.
func (c *ItemCreate) SetInput(i CreateItemInput) *ItemCreate {
	i.Mutate(c.Mutation())
	return c
}
