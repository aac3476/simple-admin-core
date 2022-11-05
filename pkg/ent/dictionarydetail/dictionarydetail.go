// Code generated by ent, DO NOT EDIT.

package dictionarydetail

import (
	"time"

	"github.com/suyuan32/simple-admin-core/pkg/gotype"
)

const (
	// Label holds the string label denoting the dictionarydetail type in the database.
	Label = "dictionary_detail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// EdgeDictionary holds the string denoting the dictionary edge name in mutations.
	EdgeDictionary = "dictionary"
	// Table holds the table name of the dictionarydetail in the database.
	Table = "dictionary_details"
	// DictionaryTable is the table that holds the dictionary relation/edge.
	DictionaryTable = "dictionary_details"
	// DictionaryInverseTable is the table name for the Dictionary entity.
	// It exists in this package in order to avoid circular dependency with the "dictionary" package.
	DictionaryInverseTable = "dictionaries"
	// DictionaryColumn is the table column denoting the dictionary relation/edge.
	DictionaryColumn = "dictionary_details"
)

// Columns holds all SQL columns for dictionarydetail fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldTitle,
	FieldKey,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "dictionary_details"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"dictionary_details",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus gotype.Status
)
