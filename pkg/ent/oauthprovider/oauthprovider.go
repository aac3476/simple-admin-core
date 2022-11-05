// Code generated by ent, DO NOT EDIT.

package oauthprovider

import (
	"time"
)

const (
	// Label holds the string label denoting the oauthprovider type in the database.
	Label = "oauth_provider"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldClientID holds the string denoting the client_id field in the database.
	FieldClientID = "client_id"
	// FieldClientSecret holds the string denoting the client_secret field in the database.
	FieldClientSecret = "client_secret"
	// FieldRedirectURL holds the string denoting the redirect_url field in the database.
	FieldRedirectURL = "redirect_url"
	// FieldScopes holds the string denoting the scopes field in the database.
	FieldScopes = "scopes"
	// FieldAuthURL holds the string denoting the auth_url field in the database.
	FieldAuthURL = "auth_url"
	// FieldTokenURL holds the string denoting the token_url field in the database.
	FieldTokenURL = "token_url"
	// FieldAuthStyle holds the string denoting the auth_style field in the database.
	FieldAuthStyle = "auth_style"
	// FieldInfoURL holds the string denoting the info_url field in the database.
	FieldInfoURL = "info_url"
	// Table holds the table name of the oauthprovider in the database.
	Table = "oauth_providers"
)

// Columns holds all SQL columns for oauthprovider fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldClientID,
	FieldClientSecret,
	FieldRedirectURL,
	FieldScopes,
	FieldAuthURL,
	FieldTokenURL,
	FieldAuthStyle,
	FieldInfoURL,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
