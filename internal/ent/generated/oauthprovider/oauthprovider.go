// Code generated by ent, DO NOT EDIT.

package oauthprovider

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
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
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldMappingID holds the string denoting the mapping_id field in the database.
	FieldMappingID = "mapping_id"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
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
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the oauthprovider in the database.
	Table = "oauth_providers"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "oauth_providers"
	// OwnerInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OwnerInverseTable = "organizations"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"
)

// Columns holds all SQL columns for oauthprovider fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldMappingID,
	FieldTags,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldOwnerID,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/theopenlane/core/internal/ent/generated/runtime"
var (
	Hooks        [4]ent.Hook
	Interceptors [2]ent.Interceptor
	Policy       ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultMappingID holds the default value on creation for the "mapping_id" field.
	DefaultMappingID func() string
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// OwnerIDValidator is a validator for the "owner_id" field. It is called by the builders before save.
	OwnerIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the OauthProvider queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByMappingID orders the results by the mapping_id field.
func ByMappingID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMappingID, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByDeletedBy orders the results by the deleted_by field.
func ByDeletedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedBy, opts...).ToFunc()
}

// ByOwnerID orders the results by the owner_id field.
func ByOwnerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByClientID orders the results by the client_id field.
func ByClientID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientID, opts...).ToFunc()
}

// ByClientSecret orders the results by the client_secret field.
func ByClientSecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientSecret, opts...).ToFunc()
}

// ByRedirectURL orders the results by the redirect_url field.
func ByRedirectURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRedirectURL, opts...).ToFunc()
}

// ByScopes orders the results by the scopes field.
func ByScopes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScopes, opts...).ToFunc()
}

// ByAuthURL orders the results by the auth_url field.
func ByAuthURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthURL, opts...).ToFunc()
}

// ByTokenURL orders the results by the token_url field.
func ByTokenURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTokenURL, opts...).ToFunc()
}

// ByAuthStyle orders the results by the auth_style field.
func ByAuthStyle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthStyle, opts...).ToFunc()
}

// ByInfoURL orders the results by the info_url field.
func ByInfoURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInfoURL, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}