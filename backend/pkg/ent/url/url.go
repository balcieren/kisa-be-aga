// Code generated by ent, DO NOT EDIT.

package url

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the url type in the database.
	Label = "url"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldShortID holds the string denoting the short_id field in the database.
	FieldShortID = "short_id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the url in the database.
	Table = "urls"
)

// Columns holds all SQL columns for url fields.
var Columns = []string{
	FieldID,
	FieldShortID,
	FieldURL,
	FieldCreatedAt,
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
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
