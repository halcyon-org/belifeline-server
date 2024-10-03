// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/halcyon-org/kizuna/ent/koyodata"
	"github.com/halcyon-org/kizuna/ent/predicate"
)

// KoyoDataDelete is the builder for deleting a KoyoData entity.
type KoyoDataDelete struct {
	config
	hooks    []Hook
	mutation *KoyoDataMutation
}

// Where appends a list predicates to the KoyoDataDelete builder.
func (kdd *KoyoDataDelete) Where(ps ...predicate.KoyoData) *KoyoDataDelete {
	kdd.mutation.Where(ps...)
	return kdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (kdd *KoyoDataDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, kdd.sqlExec, kdd.mutation, kdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (kdd *KoyoDataDelete) ExecX(ctx context.Context) int {
	n, err := kdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (kdd *KoyoDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(koyodata.Table, sqlgraph.NewFieldSpec(koyodata.FieldID, field.TypeString))
	if ps := kdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, kdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	kdd.mutation.done = true
	return affected, err
}

// KoyoDataDeleteOne is the builder for deleting a single KoyoData entity.
type KoyoDataDeleteOne struct {
	kdd *KoyoDataDelete
}

// Where appends a list predicates to the KoyoDataDelete builder.
func (kddo *KoyoDataDeleteOne) Where(ps ...predicate.KoyoData) *KoyoDataDeleteOne {
	kddo.kdd.mutation.Where(ps...)
	return kddo
}

// Exec executes the deletion query.
func (kddo *KoyoDataDeleteOne) Exec(ctx context.Context) error {
	n, err := kddo.kdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{koyodata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (kddo *KoyoDataDeleteOne) ExecX(ctx context.Context) {
	if err := kddo.Exec(ctx); err != nil {
		panic(err)
	}
}
