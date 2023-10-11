// Code generated by ent, DO NOT EDIT.

package irregular

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the irregular type in the database.
	Label = "irregular"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldViewURL holds the string denoting the view_url field in the database.
	FieldViewURL = "view_url"
	// FieldDownloadURL holds the string denoting the download_url field in the database.
	FieldDownloadURL = "download_url"
	// FieldFileName holds the string denoting the file_name field in the database.
	FieldFileName = "file_name"
	// FieldFileSize holds the string denoting the file_size field in the database.
	FieldFileSize = "file_size"
	// FieldPublishDate holds the string denoting the publish_date field in the database.
	FieldPublishDate = "publish_date"
	// Table holds the table name of the irregular in the database.
	Table = "irregulars"
)

// Columns holds all SQL columns for irregular fields.
var Columns = []string{
	FieldID,
	FieldViewURL,
	FieldDownloadURL,
	FieldFileName,
	FieldFileSize,
	FieldPublishDate,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// FileNameValidator is a validator for the "file_name" field. It is called by the builders before save.
	FileNameValidator func(string) error
	// FileSizeValidator is a validator for the "file_size" field. It is called by the builders before save.
	FileSizeValidator func(int) error
	// DefaultPublishDate holds the default value on creation for the "publish_date" field.
	DefaultPublishDate func() time.Time
)

// OrderOption defines the ordering options for the Irregular queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByViewURL orders the results by the view_url field.
func ByViewURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldViewURL, opts...).ToFunc()
}

// ByDownloadURL orders the results by the download_url field.
func ByDownloadURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDownloadURL, opts...).ToFunc()
}

// ByFileName orders the results by the file_name field.
func ByFileName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileName, opts...).ToFunc()
}

// ByFileSize orders the results by the file_size field.
func ByFileSize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileSize, opts...).ToFunc()
}

// ByPublishDate orders the results by the publish_date field.
func ByPublishDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublishDate, opts...).ToFunc()
}
