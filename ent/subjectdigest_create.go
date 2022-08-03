// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/testifysec/archivist/ent/subject"
	"github.com/testifysec/archivist/ent/subjectdigest"
)

// SubjectDigestCreate is the builder for creating a SubjectDigest entity.
type SubjectDigestCreate struct {
	config
	mutation *SubjectDigestMutation
	hooks    []Hook
}

// SetAlgorithm sets the "algorithm" field.
func (sdc *SubjectDigestCreate) SetAlgorithm(s string) *SubjectDigestCreate {
	sdc.mutation.SetAlgorithm(s)
	return sdc
}

// SetValue sets the "value" field.
func (sdc *SubjectDigestCreate) SetValue(s string) *SubjectDigestCreate {
	sdc.mutation.SetValue(s)
	return sdc
}

// SetSubjectID sets the "subject" edge to the Subject entity by ID.
func (sdc *SubjectDigestCreate) SetSubjectID(id int) *SubjectDigestCreate {
	sdc.mutation.SetSubjectID(id)
	return sdc
}

// SetNillableSubjectID sets the "subject" edge to the Subject entity by ID if the given value is not nil.
func (sdc *SubjectDigestCreate) SetNillableSubjectID(id *int) *SubjectDigestCreate {
	if id != nil {
		sdc = sdc.SetSubjectID(*id)
	}
	return sdc
}

// SetSubject sets the "subject" edge to the Subject entity.
func (sdc *SubjectDigestCreate) SetSubject(s *Subject) *SubjectDigestCreate {
	return sdc.SetSubjectID(s.ID)
}

// Mutation returns the SubjectDigestMutation object of the builder.
func (sdc *SubjectDigestCreate) Mutation() *SubjectDigestMutation {
	return sdc.mutation
}

// Save creates the SubjectDigest in the database.
func (sdc *SubjectDigestCreate) Save(ctx context.Context) (*SubjectDigest, error) {
	var (
		err  error
		node *SubjectDigest
	)
	if len(sdc.hooks) == 0 {
		if err = sdc.check(); err != nil {
			return nil, err
		}
		node, err = sdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubjectDigestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sdc.check(); err != nil {
				return nil, err
			}
			sdc.mutation = mutation
			if node, err = sdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sdc.hooks) - 1; i >= 0; i-- {
			if sdc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sdc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SubjectDigest)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SubjectDigestMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sdc *SubjectDigestCreate) SaveX(ctx context.Context) *SubjectDigest {
	v, err := sdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdc *SubjectDigestCreate) Exec(ctx context.Context) error {
	_, err := sdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdc *SubjectDigestCreate) ExecX(ctx context.Context) {
	if err := sdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdc *SubjectDigestCreate) check() error {
	if _, ok := sdc.mutation.Algorithm(); !ok {
		return &ValidationError{Name: "algorithm", err: errors.New(`ent: missing required field "SubjectDigest.algorithm"`)}
	}
	if v, ok := sdc.mutation.Algorithm(); ok {
		if err := subjectdigest.AlgorithmValidator(v); err != nil {
			return &ValidationError{Name: "algorithm", err: fmt.Errorf(`ent: validator failed for field "SubjectDigest.algorithm": %w`, err)}
		}
	}
	if _, ok := sdc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "SubjectDigest.value"`)}
	}
	if v, ok := sdc.mutation.Value(); ok {
		if err := subjectdigest.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "SubjectDigest.value": %w`, err)}
		}
	}
	return nil
}

func (sdc *SubjectDigestCreate) sqlSave(ctx context.Context) (*SubjectDigest, error) {
	_node, _spec := sdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sdc *SubjectDigestCreate) createSpec() (*SubjectDigest, *sqlgraph.CreateSpec) {
	var (
		_node = &SubjectDigest{config: sdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: subjectdigest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subjectdigest.FieldID,
			},
		}
	)
	if value, ok := sdc.mutation.Algorithm(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subjectdigest.FieldAlgorithm,
		})
		_node.Algorithm = value
	}
	if value, ok := sdc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subjectdigest.FieldValue,
		})
		_node.Value = value
	}
	if nodes := sdc.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subjectdigest.SubjectTable,
			Columns: []string{subjectdigest.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.subject_subject_digests = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubjectDigestCreateBulk is the builder for creating many SubjectDigest entities in bulk.
type SubjectDigestCreateBulk struct {
	config
	builders []*SubjectDigestCreate
}

// Save creates the SubjectDigest entities in the database.
func (sdcb *SubjectDigestCreateBulk) Save(ctx context.Context) ([]*SubjectDigest, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sdcb.builders))
	nodes := make([]*SubjectDigest, len(sdcb.builders))
	mutators := make([]Mutator, len(sdcb.builders))
	for i := range sdcb.builders {
		func(i int, root context.Context) {
			builder := sdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubjectDigestMutation)
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
					_, err = mutators[i+1].Mutate(root, sdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sdcb *SubjectDigestCreateBulk) SaveX(ctx context.Context) []*SubjectDigest {
	v, err := sdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdcb *SubjectDigestCreateBulk) Exec(ctx context.Context) error {
	_, err := sdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdcb *SubjectDigestCreateBulk) ExecX(ctx context.Context) {
	if err := sdcb.Exec(ctx); err != nil {
		panic(err)
	}
}