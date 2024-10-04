// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/halcyon-org/kizuna/ent/clientdata"
	"github.com/halcyon-org/kizuna/ent/predicate"
)

// ClientDataUpdate is the builder for updating ClientData entities.
type ClientDataUpdate struct {
	config
	hooks    []Hook
	mutation *ClientDataMutation
}

// Where appends a list predicates to the ClientDataUpdate builder.
func (cdu *ClientDataUpdate) Where(ps ...predicate.ClientData) *ClientDataUpdate {
	cdu.mutation.Where(ps...)
	return cdu
}

// SetUsername sets the "username" field.
func (cdu *ClientDataUpdate) SetUsername(s string) *ClientDataUpdate {
	cdu.mutation.SetUsername(s)
	return cdu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (cdu *ClientDataUpdate) SetNillableUsername(s *string) *ClientDataUpdate {
	if s != nil {
		cdu.SetUsername(*s)
	}
	return cdu
}

// SetAPIKey sets the "api_key" field.
func (cdu *ClientDataUpdate) SetAPIKey(s string) *ClientDataUpdate {
	cdu.mutation.SetAPIKey(s)
	return cdu
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (cdu *ClientDataUpdate) SetNillableAPIKey(s *string) *ClientDataUpdate {
	if s != nil {
		cdu.SetAPIKey(*s)
	}
	return cdu
}

// SetLastUsedAt sets the "last_used_at" field.
func (cdu *ClientDataUpdate) SetLastUsedAt(t time.Time) *ClientDataUpdate {
	cdu.mutation.SetLastUsedAt(t)
	return cdu
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (cdu *ClientDataUpdate) SetNillableLastUsedAt(t *time.Time) *ClientDataUpdate {
	if t != nil {
		cdu.SetLastUsedAt(*t)
	}
	return cdu
}

// SetLastUpdatedAt sets the "last_updated_at" field.
func (cdu *ClientDataUpdate) SetLastUpdatedAt(t time.Time) *ClientDataUpdate {
	cdu.mutation.SetLastUpdatedAt(t)
	return cdu
}

// Mutation returns the ClientDataMutation object of the builder.
func (cdu *ClientDataUpdate) Mutation() *ClientDataMutation {
	return cdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cdu *ClientDataUpdate) Save(ctx context.Context) (int, error) {
	cdu.defaults()
	return withHooks(ctx, cdu.sqlSave, cdu.mutation, cdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cdu *ClientDataUpdate) SaveX(ctx context.Context) int {
	affected, err := cdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cdu *ClientDataUpdate) Exec(ctx context.Context) error {
	_, err := cdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cdu *ClientDataUpdate) ExecX(ctx context.Context) {
	if err := cdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cdu *ClientDataUpdate) defaults() {
	if _, ok := cdu.mutation.LastUpdatedAt(); !ok {
		v := clientdata.UpdateDefaultLastUpdatedAt()
		cdu.mutation.SetLastUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cdu *ClientDataUpdate) check() error {
	if v, ok := cdu.mutation.Username(); ok {
		if err := clientdata.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "ClientData.username": %w`, err)}
		}
	}
	return nil
}

func (cdu *ClientDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cdu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(clientdata.Table, clientdata.Columns, sqlgraph.NewFieldSpec(clientdata.FieldID, field.TypeString))
	if ps := cdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cdu.mutation.Username(); ok {
		_spec.SetField(clientdata.FieldUsername, field.TypeString, value)
	}
	if value, ok := cdu.mutation.APIKey(); ok {
		_spec.SetField(clientdata.FieldAPIKey, field.TypeString, value)
	}
	if value, ok := cdu.mutation.LastUsedAt(); ok {
		_spec.SetField(clientdata.FieldLastUsedAt, field.TypeTime, value)
	}
	if value, ok := cdu.mutation.LastUpdatedAt(); ok {
		_spec.SetField(clientdata.FieldLastUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clientdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cdu.mutation.done = true
	return n, nil
}

// ClientDataUpdateOne is the builder for updating a single ClientData entity.
type ClientDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClientDataMutation
}

// SetUsername sets the "username" field.
func (cduo *ClientDataUpdateOne) SetUsername(s string) *ClientDataUpdateOne {
	cduo.mutation.SetUsername(s)
	return cduo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (cduo *ClientDataUpdateOne) SetNillableUsername(s *string) *ClientDataUpdateOne {
	if s != nil {
		cduo.SetUsername(*s)
	}
	return cduo
}

// SetAPIKey sets the "api_key" field.
func (cduo *ClientDataUpdateOne) SetAPIKey(s string) *ClientDataUpdateOne {
	cduo.mutation.SetAPIKey(s)
	return cduo
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (cduo *ClientDataUpdateOne) SetNillableAPIKey(s *string) *ClientDataUpdateOne {
	if s != nil {
		cduo.SetAPIKey(*s)
	}
	return cduo
}

// SetLastUsedAt sets the "last_used_at" field.
func (cduo *ClientDataUpdateOne) SetLastUsedAt(t time.Time) *ClientDataUpdateOne {
	cduo.mutation.SetLastUsedAt(t)
	return cduo
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (cduo *ClientDataUpdateOne) SetNillableLastUsedAt(t *time.Time) *ClientDataUpdateOne {
	if t != nil {
		cduo.SetLastUsedAt(*t)
	}
	return cduo
}

// SetLastUpdatedAt sets the "last_updated_at" field.
func (cduo *ClientDataUpdateOne) SetLastUpdatedAt(t time.Time) *ClientDataUpdateOne {
	cduo.mutation.SetLastUpdatedAt(t)
	return cduo
}

// Mutation returns the ClientDataMutation object of the builder.
func (cduo *ClientDataUpdateOne) Mutation() *ClientDataMutation {
	return cduo.mutation
}

// Where appends a list predicates to the ClientDataUpdate builder.
func (cduo *ClientDataUpdateOne) Where(ps ...predicate.ClientData) *ClientDataUpdateOne {
	cduo.mutation.Where(ps...)
	return cduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cduo *ClientDataUpdateOne) Select(field string, fields ...string) *ClientDataUpdateOne {
	cduo.fields = append([]string{field}, fields...)
	return cduo
}

// Save executes the query and returns the updated ClientData entity.
func (cduo *ClientDataUpdateOne) Save(ctx context.Context) (*ClientData, error) {
	cduo.defaults()
	return withHooks(ctx, cduo.sqlSave, cduo.mutation, cduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cduo *ClientDataUpdateOne) SaveX(ctx context.Context) *ClientData {
	node, err := cduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cduo *ClientDataUpdateOne) Exec(ctx context.Context) error {
	_, err := cduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cduo *ClientDataUpdateOne) ExecX(ctx context.Context) {
	if err := cduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cduo *ClientDataUpdateOne) defaults() {
	if _, ok := cduo.mutation.LastUpdatedAt(); !ok {
		v := clientdata.UpdateDefaultLastUpdatedAt()
		cduo.mutation.SetLastUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cduo *ClientDataUpdateOne) check() error {
	if v, ok := cduo.mutation.Username(); ok {
		if err := clientdata.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "ClientData.username": %w`, err)}
		}
	}
	return nil
}

func (cduo *ClientDataUpdateOne) sqlSave(ctx context.Context) (_node *ClientData, err error) {
	if err := cduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(clientdata.Table, clientdata.Columns, sqlgraph.NewFieldSpec(clientdata.FieldID, field.TypeString))
	id, ok := cduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ClientData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, clientdata.FieldID)
		for _, f := range fields {
			if !clientdata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != clientdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cduo.mutation.Username(); ok {
		_spec.SetField(clientdata.FieldUsername, field.TypeString, value)
	}
	if value, ok := cduo.mutation.APIKey(); ok {
		_spec.SetField(clientdata.FieldAPIKey, field.TypeString, value)
	}
	if value, ok := cduo.mutation.LastUsedAt(); ok {
		_spec.SetField(clientdata.FieldLastUsedAt, field.TypeTime, value)
	}
	if value, ok := cduo.mutation.LastUpdatedAt(); ok {
		_spec.SetField(clientdata.FieldLastUpdatedAt, field.TypeTime, value)
	}
	_node = &ClientData{config: cduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clientdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cduo.mutation.done = true
	return _node, nil
}
