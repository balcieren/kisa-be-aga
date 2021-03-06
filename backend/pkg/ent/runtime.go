// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/pkg/ent/schema"
	"backend/pkg/ent/url"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	urlFields := schema.Url{}.Fields()
	_ = urlFields
	// urlDescURL is the schema descriptor for url field.
	urlDescURL := urlFields[2].Descriptor()
	// url.URLValidator is a validator for the "url" field. It is called by the builders before save.
	url.URLValidator = urlDescURL.Validators[0].(func(string) error)
	// urlDescCreatedAt is the schema descriptor for created_at field.
	urlDescCreatedAt := urlFields[3].Descriptor()
	// url.DefaultCreatedAt holds the default value on creation for the created_at field.
	url.DefaultCreatedAt = urlDescCreatedAt.Default.(func() time.Time)
	// urlDescID is the schema descriptor for id field.
	urlDescID := urlFields[0].Descriptor()
	// url.DefaultID holds the default value on creation for the id field.
	url.DefaultID = urlDescID.Default.(func() uuid.UUID)
}
