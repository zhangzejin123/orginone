// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"fmt"
	"orginone/common/schema/asmarketappcomponent"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsMarketAppComponentDelete is the builder for deleting a AsMarketAppComponent entity.
type AsMarketAppComponentDelete struct {
	config
	hooks    []Hook
	mutation *AsMarketAppComponentMutation
}

// Where appends a list predicates to the AsMarketAppComponentDelete builder.
func (amacd *AsMarketAppComponentDelete) Where(ps ...predicate.AsMarketAppComponent) *AsMarketAppComponentDelete {
	amacd.mutation.Where(ps...)
	return amacd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (amacd *AsMarketAppComponentDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(amacd.hooks) == 0 {
		affected, err = amacd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsMarketAppComponentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			amacd.mutation = mutation
			affected, err = amacd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(amacd.hooks) - 1; i >= 0; i-- {
			if amacd.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = amacd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, amacd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (amacd *AsMarketAppComponentDelete) ExecX(ctx context.Context) int {
	n, err := amacd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (amacd *AsMarketAppComponentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: asmarketappcomponent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asmarketappcomponent.FieldID,
			},
		},
	}
	if ps := amacd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, amacd.driver, _spec)
}

// AsMarketAppComponentDeleteOne is the builder for deleting a single AsMarketAppComponent entity.
type AsMarketAppComponentDeleteOne struct {
	amacd *AsMarketAppComponentDelete
}

// Exec executes the deletion query.
func (amacdo *AsMarketAppComponentDeleteOne) Exec(ctx context.Context) error {
	n, err := amacdo.amacd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{asmarketappcomponent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (amacdo *AsMarketAppComponentDeleteOne) ExecX(ctx context.Context) {
	amacdo.amacd.ExecX(ctx)
}