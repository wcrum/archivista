// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/testifysec/archivist/ent/dsse"
	"github.com/testifysec/archivist/ent/payloaddigest"
)

// PayloadDigestCreate is the builder for creating a PayloadDigest entity.
type PayloadDigestCreate struct {
	config
	mutation *PayloadDigestMutation
	hooks    []Hook
}

// SetAlgorithm sets the "algorithm" field.
func (pdc *PayloadDigestCreate) SetAlgorithm(s string) *PayloadDigestCreate {
	pdc.mutation.SetAlgorithm(s)
	return pdc
}

// SetValue sets the "value" field.
func (pdc *PayloadDigestCreate) SetValue(s string) *PayloadDigestCreate {
	pdc.mutation.SetValue(s)
	return pdc
}

// SetDsseID sets the "dsse" edge to the Dsse entity by ID.
func (pdc *PayloadDigestCreate) SetDsseID(id int) *PayloadDigestCreate {
	pdc.mutation.SetDsseID(id)
	return pdc
}

// SetNillableDsseID sets the "dsse" edge to the Dsse entity by ID if the given value is not nil.
func (pdc *PayloadDigestCreate) SetNillableDsseID(id *int) *PayloadDigestCreate {
	if id != nil {
		pdc = pdc.SetDsseID(*id)
	}
	return pdc
}

// SetDsse sets the "dsse" edge to the Dsse entity.
func (pdc *PayloadDigestCreate) SetDsse(d *Dsse) *PayloadDigestCreate {
	return pdc.SetDsseID(d.ID)
}

// Mutation returns the PayloadDigestMutation object of the builder.
func (pdc *PayloadDigestCreate) Mutation() *PayloadDigestMutation {
	return pdc.mutation
}

// Save creates the PayloadDigest in the database.
func (pdc *PayloadDigestCreate) Save(ctx context.Context) (*PayloadDigest, error) {
	var (
		err  error
		node *PayloadDigest
	)
	if len(pdc.hooks) == 0 {
		if err = pdc.check(); err != nil {
			return nil, err
		}
		node, err = pdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PayloadDigestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pdc.check(); err != nil {
				return nil, err
			}
			pdc.mutation = mutation
			if node, err = pdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pdc.hooks) - 1; i >= 0; i-- {
			if pdc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pdc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pdc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PayloadDigest)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PayloadDigestMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pdc *PayloadDigestCreate) SaveX(ctx context.Context) *PayloadDigest {
	v, err := pdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pdc *PayloadDigestCreate) Exec(ctx context.Context) error {
	_, err := pdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdc *PayloadDigestCreate) ExecX(ctx context.Context) {
	if err := pdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pdc *PayloadDigestCreate) check() error {
	if _, ok := pdc.mutation.Algorithm(); !ok {
		return &ValidationError{Name: "algorithm", err: errors.New(`ent: missing required field "PayloadDigest.algorithm"`)}
	}
	if v, ok := pdc.mutation.Algorithm(); ok {
		if err := payloaddigest.AlgorithmValidator(v); err != nil {
			return &ValidationError{Name: "algorithm", err: fmt.Errorf(`ent: validator failed for field "PayloadDigest.algorithm": %w`, err)}
		}
	}
	if _, ok := pdc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "PayloadDigest.value"`)}
	}
	if v, ok := pdc.mutation.Value(); ok {
		if err := payloaddigest.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "PayloadDigest.value": %w`, err)}
		}
	}
	return nil
}

func (pdc *PayloadDigestCreate) sqlSave(ctx context.Context) (*PayloadDigest, error) {
	_node, _spec := pdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pdc *PayloadDigestCreate) createSpec() (*PayloadDigest, *sqlgraph.CreateSpec) {
	var (
		_node = &PayloadDigest{config: pdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: payloaddigest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payloaddigest.FieldID,
			},
		}
	)
	if value, ok := pdc.mutation.Algorithm(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payloaddigest.FieldAlgorithm,
		})
		_node.Algorithm = value
	}
	if value, ok := pdc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payloaddigest.FieldValue,
		})
		_node.Value = value
	}
	if nodes := pdc.mutation.DsseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payloaddigest.DsseTable,
			Columns: []string{payloaddigest.DsseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dsse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.dsse_payload_digests = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PayloadDigestCreateBulk is the builder for creating many PayloadDigest entities in bulk.
type PayloadDigestCreateBulk struct {
	config
	builders []*PayloadDigestCreate
}

// Save creates the PayloadDigest entities in the database.
func (pdcb *PayloadDigestCreateBulk) Save(ctx context.Context) ([]*PayloadDigest, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pdcb.builders))
	nodes := make([]*PayloadDigest, len(pdcb.builders))
	mutators := make([]Mutator, len(pdcb.builders))
	for i := range pdcb.builders {
		func(i int, root context.Context) {
			builder := pdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PayloadDigestMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pdcb *PayloadDigestCreateBulk) SaveX(ctx context.Context) []*PayloadDigest {
	v, err := pdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pdcb *PayloadDigestCreateBulk) Exec(ctx context.Context) error {
	_, err := pdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdcb *PayloadDigestCreateBulk) ExecX(ctx context.Context) {
	if err := pdcb.Exec(ctx); err != nil {
		panic(err)
	}
}