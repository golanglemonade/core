// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/customtypes"
	"github.com/theopenlane/core/internal/ent/generated/oauthprovider"
	"github.com/theopenlane/core/internal/ent/generated/organization"
	"github.com/theopenlane/core/internal/ent/generated/predicate"

	"github.com/theopenlane/core/internal/ent/generated/internal"
)

// OauthProviderUpdate is the builder for updating OauthProvider entities.
type OauthProviderUpdate struct {
	config
	hooks    []Hook
	mutation *OauthProviderMutation
}

// Where appends a list predicates to the OauthProviderUpdate builder.
func (opu *OauthProviderUpdate) Where(ps ...predicate.OauthProvider) *OauthProviderUpdate {
	opu.mutation.Where(ps...)
	return opu
}

// SetUpdatedAt sets the "updated_at" field.
func (opu *OauthProviderUpdate) SetUpdatedAt(t time.Time) *OauthProviderUpdate {
	opu.mutation.SetUpdatedAt(t)
	return opu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (opu *OauthProviderUpdate) ClearUpdatedAt() *OauthProviderUpdate {
	opu.mutation.ClearUpdatedAt()
	return opu
}

// SetUpdatedBy sets the "updated_by" field.
func (opu *OauthProviderUpdate) SetUpdatedBy(s string) *OauthProviderUpdate {
	opu.mutation.SetUpdatedBy(s)
	return opu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableUpdatedBy(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetUpdatedBy(*s)
	}
	return opu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (opu *OauthProviderUpdate) ClearUpdatedBy() *OauthProviderUpdate {
	opu.mutation.ClearUpdatedBy()
	return opu
}

// SetTags sets the "tags" field.
func (opu *OauthProviderUpdate) SetTags(s []string) *OauthProviderUpdate {
	opu.mutation.SetTags(s)
	return opu
}

// AppendTags appends s to the "tags" field.
func (opu *OauthProviderUpdate) AppendTags(s []string) *OauthProviderUpdate {
	opu.mutation.AppendTags(s)
	return opu
}

// ClearTags clears the value of the "tags" field.
func (opu *OauthProviderUpdate) ClearTags() *OauthProviderUpdate {
	opu.mutation.ClearTags()
	return opu
}

// SetDeletedAt sets the "deleted_at" field.
func (opu *OauthProviderUpdate) SetDeletedAt(t time.Time) *OauthProviderUpdate {
	opu.mutation.SetDeletedAt(t)
	return opu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableDeletedAt(t *time.Time) *OauthProviderUpdate {
	if t != nil {
		opu.SetDeletedAt(*t)
	}
	return opu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (opu *OauthProviderUpdate) ClearDeletedAt() *OauthProviderUpdate {
	opu.mutation.ClearDeletedAt()
	return opu
}

// SetDeletedBy sets the "deleted_by" field.
func (opu *OauthProviderUpdate) SetDeletedBy(s string) *OauthProviderUpdate {
	opu.mutation.SetDeletedBy(s)
	return opu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableDeletedBy(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetDeletedBy(*s)
	}
	return opu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (opu *OauthProviderUpdate) ClearDeletedBy() *OauthProviderUpdate {
	opu.mutation.ClearDeletedBy()
	return opu
}

// SetOwnerID sets the "owner_id" field.
func (opu *OauthProviderUpdate) SetOwnerID(s string) *OauthProviderUpdate {
	opu.mutation.SetOwnerID(s)
	return opu
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableOwnerID(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetOwnerID(*s)
	}
	return opu
}

// ClearOwnerID clears the value of the "owner_id" field.
func (opu *OauthProviderUpdate) ClearOwnerID() *OauthProviderUpdate {
	opu.mutation.ClearOwnerID()
	return opu
}

// SetName sets the "name" field.
func (opu *OauthProviderUpdate) SetName(s string) *OauthProviderUpdate {
	opu.mutation.SetName(s)
	return opu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableName(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetName(*s)
	}
	return opu
}

// SetClientID sets the "client_id" field.
func (opu *OauthProviderUpdate) SetClientID(s string) *OauthProviderUpdate {
	opu.mutation.SetClientID(s)
	return opu
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableClientID(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetClientID(*s)
	}
	return opu
}

// SetClientSecret sets the "client_secret" field.
func (opu *OauthProviderUpdate) SetClientSecret(s string) *OauthProviderUpdate {
	opu.mutation.SetClientSecret(s)
	return opu
}

// SetNillableClientSecret sets the "client_secret" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableClientSecret(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetClientSecret(*s)
	}
	return opu
}

// SetRedirectURL sets the "redirect_url" field.
func (opu *OauthProviderUpdate) SetRedirectURL(s string) *OauthProviderUpdate {
	opu.mutation.SetRedirectURL(s)
	return opu
}

// SetNillableRedirectURL sets the "redirect_url" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableRedirectURL(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetRedirectURL(*s)
	}
	return opu
}

// SetScopes sets the "scopes" field.
func (opu *OauthProviderUpdate) SetScopes(s string) *OauthProviderUpdate {
	opu.mutation.SetScopes(s)
	return opu
}

// SetNillableScopes sets the "scopes" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableScopes(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetScopes(*s)
	}
	return opu
}

// SetAuthURL sets the "auth_url" field.
func (opu *OauthProviderUpdate) SetAuthURL(s string) *OauthProviderUpdate {
	opu.mutation.SetAuthURL(s)
	return opu
}

// SetNillableAuthURL sets the "auth_url" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableAuthURL(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetAuthURL(*s)
	}
	return opu
}

// SetTokenURL sets the "token_url" field.
func (opu *OauthProviderUpdate) SetTokenURL(s string) *OauthProviderUpdate {
	opu.mutation.SetTokenURL(s)
	return opu
}

// SetNillableTokenURL sets the "token_url" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableTokenURL(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetTokenURL(*s)
	}
	return opu
}

// SetAuthStyle sets the "auth_style" field.
func (opu *OauthProviderUpdate) SetAuthStyle(c customtypes.Uint8) *OauthProviderUpdate {
	opu.mutation.ResetAuthStyle()
	opu.mutation.SetAuthStyle(c)
	return opu
}

// SetNillableAuthStyle sets the "auth_style" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableAuthStyle(c *customtypes.Uint8) *OauthProviderUpdate {
	if c != nil {
		opu.SetAuthStyle(*c)
	}
	return opu
}

// AddAuthStyle adds c to the "auth_style" field.
func (opu *OauthProviderUpdate) AddAuthStyle(c customtypes.Uint8) *OauthProviderUpdate {
	opu.mutation.AddAuthStyle(c)
	return opu
}

// SetInfoURL sets the "info_url" field.
func (opu *OauthProviderUpdate) SetInfoURL(s string) *OauthProviderUpdate {
	opu.mutation.SetInfoURL(s)
	return opu
}

// SetNillableInfoURL sets the "info_url" field if the given value is not nil.
func (opu *OauthProviderUpdate) SetNillableInfoURL(s *string) *OauthProviderUpdate {
	if s != nil {
		opu.SetInfoURL(*s)
	}
	return opu
}

// SetOwner sets the "owner" edge to the Organization entity.
func (opu *OauthProviderUpdate) SetOwner(o *Organization) *OauthProviderUpdate {
	return opu.SetOwnerID(o.ID)
}

// Mutation returns the OauthProviderMutation object of the builder.
func (opu *OauthProviderUpdate) Mutation() *OauthProviderMutation {
	return opu.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (opu *OauthProviderUpdate) ClearOwner() *OauthProviderUpdate {
	opu.mutation.ClearOwner()
	return opu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (opu *OauthProviderUpdate) Save(ctx context.Context) (int, error) {
	if err := opu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, opu.sqlSave, opu.mutation, opu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (opu *OauthProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := opu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (opu *OauthProviderUpdate) Exec(ctx context.Context) error {
	_, err := opu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opu *OauthProviderUpdate) ExecX(ctx context.Context) {
	if err := opu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opu *OauthProviderUpdate) defaults() error {
	if _, ok := opu.mutation.UpdatedAt(); !ok && !opu.mutation.UpdatedAtCleared() {
		if oauthprovider.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized oauthprovider.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := oauthprovider.UpdateDefaultUpdatedAt()
		opu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (opu *OauthProviderUpdate) check() error {
	if v, ok := opu.mutation.OwnerID(); ok {
		if err := oauthprovider.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`generated: validator failed for field "OauthProvider.owner_id": %w`, err)}
		}
	}
	return nil
}

func (opu *OauthProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := opu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(oauthprovider.Table, oauthprovider.Columns, sqlgraph.NewFieldSpec(oauthprovider.FieldID, field.TypeString))
	if ps := opu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if opu.mutation.CreatedAtCleared() {
		_spec.ClearField(oauthprovider.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := opu.mutation.UpdatedAt(); ok {
		_spec.SetField(oauthprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if opu.mutation.UpdatedAtCleared() {
		_spec.ClearField(oauthprovider.FieldUpdatedAt, field.TypeTime)
	}
	if opu.mutation.CreatedByCleared() {
		_spec.ClearField(oauthprovider.FieldCreatedBy, field.TypeString)
	}
	if value, ok := opu.mutation.UpdatedBy(); ok {
		_spec.SetField(oauthprovider.FieldUpdatedBy, field.TypeString, value)
	}
	if opu.mutation.UpdatedByCleared() {
		_spec.ClearField(oauthprovider.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := opu.mutation.Tags(); ok {
		_spec.SetField(oauthprovider.FieldTags, field.TypeJSON, value)
	}
	if value, ok := opu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oauthprovider.FieldTags, value)
		})
	}
	if opu.mutation.TagsCleared() {
		_spec.ClearField(oauthprovider.FieldTags, field.TypeJSON)
	}
	if value, ok := opu.mutation.DeletedAt(); ok {
		_spec.SetField(oauthprovider.FieldDeletedAt, field.TypeTime, value)
	}
	if opu.mutation.DeletedAtCleared() {
		_spec.ClearField(oauthprovider.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := opu.mutation.DeletedBy(); ok {
		_spec.SetField(oauthprovider.FieldDeletedBy, field.TypeString, value)
	}
	if opu.mutation.DeletedByCleared() {
		_spec.ClearField(oauthprovider.FieldDeletedBy, field.TypeString)
	}
	if value, ok := opu.mutation.Name(); ok {
		_spec.SetField(oauthprovider.FieldName, field.TypeString, value)
	}
	if value, ok := opu.mutation.ClientID(); ok {
		_spec.SetField(oauthprovider.FieldClientID, field.TypeString, value)
	}
	if value, ok := opu.mutation.ClientSecret(); ok {
		_spec.SetField(oauthprovider.FieldClientSecret, field.TypeString, value)
	}
	if value, ok := opu.mutation.RedirectURL(); ok {
		_spec.SetField(oauthprovider.FieldRedirectURL, field.TypeString, value)
	}
	if value, ok := opu.mutation.Scopes(); ok {
		_spec.SetField(oauthprovider.FieldScopes, field.TypeString, value)
	}
	if value, ok := opu.mutation.AuthURL(); ok {
		_spec.SetField(oauthprovider.FieldAuthURL, field.TypeString, value)
	}
	if value, ok := opu.mutation.TokenURL(); ok {
		_spec.SetField(oauthprovider.FieldTokenURL, field.TypeString, value)
	}
	if value, ok := opu.mutation.AuthStyle(); ok {
		_spec.SetField(oauthprovider.FieldAuthStyle, field.TypeUint8, value)
	}
	if value, ok := opu.mutation.AddedAuthStyle(); ok {
		_spec.AddField(oauthprovider.FieldAuthStyle, field.TypeUint8, value)
	}
	if value, ok := opu.mutation.InfoURL(); ok {
		_spec.SetField(oauthprovider.FieldInfoURL, field.TypeString, value)
	}
	if opu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauthprovider.OwnerTable,
			Columns: []string{oauthprovider.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = opu.schemaConfig.OauthProvider
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauthprovider.OwnerTable,
			Columns: []string{oauthprovider.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = opu.schemaConfig.OauthProvider
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = opu.schemaConfig.OauthProvider
	ctx = internal.NewSchemaConfigContext(ctx, opu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, opu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauthprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	opu.mutation.done = true
	return n, nil
}

// OauthProviderUpdateOne is the builder for updating a single OauthProvider entity.
type OauthProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OauthProviderMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (opuo *OauthProviderUpdateOne) SetUpdatedAt(t time.Time) *OauthProviderUpdateOne {
	opuo.mutation.SetUpdatedAt(t)
	return opuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (opuo *OauthProviderUpdateOne) ClearUpdatedAt() *OauthProviderUpdateOne {
	opuo.mutation.ClearUpdatedAt()
	return opuo
}

// SetUpdatedBy sets the "updated_by" field.
func (opuo *OauthProviderUpdateOne) SetUpdatedBy(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetUpdatedBy(s)
	return opuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableUpdatedBy(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetUpdatedBy(*s)
	}
	return opuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (opuo *OauthProviderUpdateOne) ClearUpdatedBy() *OauthProviderUpdateOne {
	opuo.mutation.ClearUpdatedBy()
	return opuo
}

// SetTags sets the "tags" field.
func (opuo *OauthProviderUpdateOne) SetTags(s []string) *OauthProviderUpdateOne {
	opuo.mutation.SetTags(s)
	return opuo
}

// AppendTags appends s to the "tags" field.
func (opuo *OauthProviderUpdateOne) AppendTags(s []string) *OauthProviderUpdateOne {
	opuo.mutation.AppendTags(s)
	return opuo
}

// ClearTags clears the value of the "tags" field.
func (opuo *OauthProviderUpdateOne) ClearTags() *OauthProviderUpdateOne {
	opuo.mutation.ClearTags()
	return opuo
}

// SetDeletedAt sets the "deleted_at" field.
func (opuo *OauthProviderUpdateOne) SetDeletedAt(t time.Time) *OauthProviderUpdateOne {
	opuo.mutation.SetDeletedAt(t)
	return opuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableDeletedAt(t *time.Time) *OauthProviderUpdateOne {
	if t != nil {
		opuo.SetDeletedAt(*t)
	}
	return opuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (opuo *OauthProviderUpdateOne) ClearDeletedAt() *OauthProviderUpdateOne {
	opuo.mutation.ClearDeletedAt()
	return opuo
}

// SetDeletedBy sets the "deleted_by" field.
func (opuo *OauthProviderUpdateOne) SetDeletedBy(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetDeletedBy(s)
	return opuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableDeletedBy(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetDeletedBy(*s)
	}
	return opuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (opuo *OauthProviderUpdateOne) ClearDeletedBy() *OauthProviderUpdateOne {
	opuo.mutation.ClearDeletedBy()
	return opuo
}

// SetOwnerID sets the "owner_id" field.
func (opuo *OauthProviderUpdateOne) SetOwnerID(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetOwnerID(s)
	return opuo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableOwnerID(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetOwnerID(*s)
	}
	return opuo
}

// ClearOwnerID clears the value of the "owner_id" field.
func (opuo *OauthProviderUpdateOne) ClearOwnerID() *OauthProviderUpdateOne {
	opuo.mutation.ClearOwnerID()
	return opuo
}

// SetName sets the "name" field.
func (opuo *OauthProviderUpdateOne) SetName(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetName(s)
	return opuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableName(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetName(*s)
	}
	return opuo
}

// SetClientID sets the "client_id" field.
func (opuo *OauthProviderUpdateOne) SetClientID(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetClientID(s)
	return opuo
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableClientID(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetClientID(*s)
	}
	return opuo
}

// SetClientSecret sets the "client_secret" field.
func (opuo *OauthProviderUpdateOne) SetClientSecret(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetClientSecret(s)
	return opuo
}

// SetNillableClientSecret sets the "client_secret" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableClientSecret(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetClientSecret(*s)
	}
	return opuo
}

// SetRedirectURL sets the "redirect_url" field.
func (opuo *OauthProviderUpdateOne) SetRedirectURL(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetRedirectURL(s)
	return opuo
}

// SetNillableRedirectURL sets the "redirect_url" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableRedirectURL(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetRedirectURL(*s)
	}
	return opuo
}

// SetScopes sets the "scopes" field.
func (opuo *OauthProviderUpdateOne) SetScopes(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetScopes(s)
	return opuo
}

// SetNillableScopes sets the "scopes" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableScopes(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetScopes(*s)
	}
	return opuo
}

// SetAuthURL sets the "auth_url" field.
func (opuo *OauthProviderUpdateOne) SetAuthURL(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetAuthURL(s)
	return opuo
}

// SetNillableAuthURL sets the "auth_url" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableAuthURL(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetAuthURL(*s)
	}
	return opuo
}

// SetTokenURL sets the "token_url" field.
func (opuo *OauthProviderUpdateOne) SetTokenURL(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetTokenURL(s)
	return opuo
}

// SetNillableTokenURL sets the "token_url" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableTokenURL(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetTokenURL(*s)
	}
	return opuo
}

// SetAuthStyle sets the "auth_style" field.
func (opuo *OauthProviderUpdateOne) SetAuthStyle(c customtypes.Uint8) *OauthProviderUpdateOne {
	opuo.mutation.ResetAuthStyle()
	opuo.mutation.SetAuthStyle(c)
	return opuo
}

// SetNillableAuthStyle sets the "auth_style" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableAuthStyle(c *customtypes.Uint8) *OauthProviderUpdateOne {
	if c != nil {
		opuo.SetAuthStyle(*c)
	}
	return opuo
}

// AddAuthStyle adds c to the "auth_style" field.
func (opuo *OauthProviderUpdateOne) AddAuthStyle(c customtypes.Uint8) *OauthProviderUpdateOne {
	opuo.mutation.AddAuthStyle(c)
	return opuo
}

// SetInfoURL sets the "info_url" field.
func (opuo *OauthProviderUpdateOne) SetInfoURL(s string) *OauthProviderUpdateOne {
	opuo.mutation.SetInfoURL(s)
	return opuo
}

// SetNillableInfoURL sets the "info_url" field if the given value is not nil.
func (opuo *OauthProviderUpdateOne) SetNillableInfoURL(s *string) *OauthProviderUpdateOne {
	if s != nil {
		opuo.SetInfoURL(*s)
	}
	return opuo
}

// SetOwner sets the "owner" edge to the Organization entity.
func (opuo *OauthProviderUpdateOne) SetOwner(o *Organization) *OauthProviderUpdateOne {
	return opuo.SetOwnerID(o.ID)
}

// Mutation returns the OauthProviderMutation object of the builder.
func (opuo *OauthProviderUpdateOne) Mutation() *OauthProviderMutation {
	return opuo.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (opuo *OauthProviderUpdateOne) ClearOwner() *OauthProviderUpdateOne {
	opuo.mutation.ClearOwner()
	return opuo
}

// Where appends a list predicates to the OauthProviderUpdate builder.
func (opuo *OauthProviderUpdateOne) Where(ps ...predicate.OauthProvider) *OauthProviderUpdateOne {
	opuo.mutation.Where(ps...)
	return opuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (opuo *OauthProviderUpdateOne) Select(field string, fields ...string) *OauthProviderUpdateOne {
	opuo.fields = append([]string{field}, fields...)
	return opuo
}

// Save executes the query and returns the updated OauthProvider entity.
func (opuo *OauthProviderUpdateOne) Save(ctx context.Context) (*OauthProvider, error) {
	if err := opuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, opuo.sqlSave, opuo.mutation, opuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (opuo *OauthProviderUpdateOne) SaveX(ctx context.Context) *OauthProvider {
	node, err := opuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (opuo *OauthProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := opuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opuo *OauthProviderUpdateOne) ExecX(ctx context.Context) {
	if err := opuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opuo *OauthProviderUpdateOne) defaults() error {
	if _, ok := opuo.mutation.UpdatedAt(); !ok && !opuo.mutation.UpdatedAtCleared() {
		if oauthprovider.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized oauthprovider.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := oauthprovider.UpdateDefaultUpdatedAt()
		opuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (opuo *OauthProviderUpdateOne) check() error {
	if v, ok := opuo.mutation.OwnerID(); ok {
		if err := oauthprovider.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`generated: validator failed for field "OauthProvider.owner_id": %w`, err)}
		}
	}
	return nil
}

func (opuo *OauthProviderUpdateOne) sqlSave(ctx context.Context) (_node *OauthProvider, err error) {
	if err := opuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(oauthprovider.Table, oauthprovider.Columns, sqlgraph.NewFieldSpec(oauthprovider.FieldID, field.TypeString))
	id, ok := opuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "OauthProvider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := opuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauthprovider.FieldID)
		for _, f := range fields {
			if !oauthprovider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != oauthprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := opuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if opuo.mutation.CreatedAtCleared() {
		_spec.ClearField(oauthprovider.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := opuo.mutation.UpdatedAt(); ok {
		_spec.SetField(oauthprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if opuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(oauthprovider.FieldUpdatedAt, field.TypeTime)
	}
	if opuo.mutation.CreatedByCleared() {
		_spec.ClearField(oauthprovider.FieldCreatedBy, field.TypeString)
	}
	if value, ok := opuo.mutation.UpdatedBy(); ok {
		_spec.SetField(oauthprovider.FieldUpdatedBy, field.TypeString, value)
	}
	if opuo.mutation.UpdatedByCleared() {
		_spec.ClearField(oauthprovider.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := opuo.mutation.Tags(); ok {
		_spec.SetField(oauthprovider.FieldTags, field.TypeJSON, value)
	}
	if value, ok := opuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oauthprovider.FieldTags, value)
		})
	}
	if opuo.mutation.TagsCleared() {
		_spec.ClearField(oauthprovider.FieldTags, field.TypeJSON)
	}
	if value, ok := opuo.mutation.DeletedAt(); ok {
		_spec.SetField(oauthprovider.FieldDeletedAt, field.TypeTime, value)
	}
	if opuo.mutation.DeletedAtCleared() {
		_spec.ClearField(oauthprovider.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := opuo.mutation.DeletedBy(); ok {
		_spec.SetField(oauthprovider.FieldDeletedBy, field.TypeString, value)
	}
	if opuo.mutation.DeletedByCleared() {
		_spec.ClearField(oauthprovider.FieldDeletedBy, field.TypeString)
	}
	if value, ok := opuo.mutation.Name(); ok {
		_spec.SetField(oauthprovider.FieldName, field.TypeString, value)
	}
	if value, ok := opuo.mutation.ClientID(); ok {
		_spec.SetField(oauthprovider.FieldClientID, field.TypeString, value)
	}
	if value, ok := opuo.mutation.ClientSecret(); ok {
		_spec.SetField(oauthprovider.FieldClientSecret, field.TypeString, value)
	}
	if value, ok := opuo.mutation.RedirectURL(); ok {
		_spec.SetField(oauthprovider.FieldRedirectURL, field.TypeString, value)
	}
	if value, ok := opuo.mutation.Scopes(); ok {
		_spec.SetField(oauthprovider.FieldScopes, field.TypeString, value)
	}
	if value, ok := opuo.mutation.AuthURL(); ok {
		_spec.SetField(oauthprovider.FieldAuthURL, field.TypeString, value)
	}
	if value, ok := opuo.mutation.TokenURL(); ok {
		_spec.SetField(oauthprovider.FieldTokenURL, field.TypeString, value)
	}
	if value, ok := opuo.mutation.AuthStyle(); ok {
		_spec.SetField(oauthprovider.FieldAuthStyle, field.TypeUint8, value)
	}
	if value, ok := opuo.mutation.AddedAuthStyle(); ok {
		_spec.AddField(oauthprovider.FieldAuthStyle, field.TypeUint8, value)
	}
	if value, ok := opuo.mutation.InfoURL(); ok {
		_spec.SetField(oauthprovider.FieldInfoURL, field.TypeString, value)
	}
	if opuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauthprovider.OwnerTable,
			Columns: []string{oauthprovider.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = opuo.schemaConfig.OauthProvider
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauthprovider.OwnerTable,
			Columns: []string{oauthprovider.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = opuo.schemaConfig.OauthProvider
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = opuo.schemaConfig.OauthProvider
	ctx = internal.NewSchemaConfigContext(ctx, opuo.schemaConfig)
	_node = &OauthProvider{config: opuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, opuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauthprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	opuo.mutation.done = true
	return _node, nil
}