// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/in-toto/archivista/ent/predicate"
	"github.com/in-toto/archivista/ent/statement"
	"github.com/in-toto/archivista/ent/subject"
	"github.com/in-toto/archivista/ent/subjectdigest"
)

// SubjectUpdate is the builder for updating Subject entities.
type SubjectUpdate struct {
	config
	hooks    []Hook
	mutation *SubjectMutation
}

// Where appends a list predicates to the SubjectUpdate builder.
func (su *SubjectUpdate) Where(ps ...predicate.Subject) *SubjectUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *SubjectUpdate) SetName(s string) *SubjectUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *SubjectUpdate) SetNillableName(s *string) *SubjectUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// AddSubjectDigestIDs adds the "subject_digests" edge to the SubjectDigest entity by IDs.
func (su *SubjectUpdate) AddSubjectDigestIDs(ids ...uuid.UUID) *SubjectUpdate {
	su.mutation.AddSubjectDigestIDs(ids...)
	return su
}

// AddSubjectDigests adds the "subject_digests" edges to the SubjectDigest entity.
func (su *SubjectUpdate) AddSubjectDigests(s ...*SubjectDigest) *SubjectUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddSubjectDigestIDs(ids...)
}

// SetStatementID sets the "statement" edge to the Statement entity by ID.
func (su *SubjectUpdate) SetStatementID(id uuid.UUID) *SubjectUpdate {
	su.mutation.SetStatementID(id)
	return su
}

// SetNillableStatementID sets the "statement" edge to the Statement entity by ID if the given value is not nil.
func (su *SubjectUpdate) SetNillableStatementID(id *uuid.UUID) *SubjectUpdate {
	if id != nil {
		su = su.SetStatementID(*id)
	}
	return su
}

// SetStatement sets the "statement" edge to the Statement entity.
func (su *SubjectUpdate) SetStatement(s *Statement) *SubjectUpdate {
	return su.SetStatementID(s.ID)
}

// Mutation returns the SubjectMutation object of the builder.
func (su *SubjectUpdate) Mutation() *SubjectMutation {
	return su.mutation
}

// ClearSubjectDigests clears all "subject_digests" edges to the SubjectDigest entity.
func (su *SubjectUpdate) ClearSubjectDigests() *SubjectUpdate {
	su.mutation.ClearSubjectDigests()
	return su
}

// RemoveSubjectDigestIDs removes the "subject_digests" edge to SubjectDigest entities by IDs.
func (su *SubjectUpdate) RemoveSubjectDigestIDs(ids ...uuid.UUID) *SubjectUpdate {
	su.mutation.RemoveSubjectDigestIDs(ids...)
	return su
}

// RemoveSubjectDigests removes "subject_digests" edges to SubjectDigest entities.
func (su *SubjectUpdate) RemoveSubjectDigests(s ...*SubjectDigest) *SubjectUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveSubjectDigestIDs(ids...)
}

// ClearStatement clears the "statement" edge to the Statement entity.
func (su *SubjectUpdate) ClearStatement() *SubjectUpdate {
	su.mutation.ClearStatement()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubjectUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubjectUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubjectUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubjectUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SubjectUpdate) check() error {
	if v, ok := su.mutation.Name(); ok {
		if err := subject.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Subject.name": %w`, err)}
		}
	}
	return nil
}

func (su *SubjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(subject.Table, subject.Columns, sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(subject.FieldName, field.TypeString, value)
	}
	if su.mutation.SubjectDigestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.SubjectDigestsTable,
			Columns: []string{subject.SubjectDigestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectdigest.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedSubjectDigestsIDs(); len(nodes) > 0 && !su.mutation.SubjectDigestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.SubjectDigestsTable,
			Columns: []string{subject.SubjectDigestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectdigest.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SubjectDigestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.SubjectDigestsTable,
			Columns: []string{subject.SubjectDigestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectdigest.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.StatementCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.StatementTable,
			Columns: []string{subject.StatementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statement.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.StatementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.StatementTable,
			Columns: []string{subject.StatementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statement.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SubjectUpdateOne is the builder for updating a single Subject entity.
type SubjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubjectMutation
}

// SetName sets the "name" field.
func (suo *SubjectUpdateOne) SetName(s string) *SubjectUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *SubjectUpdateOne) SetNillableName(s *string) *SubjectUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// AddSubjectDigestIDs adds the "subject_digests" edge to the SubjectDigest entity by IDs.
func (suo *SubjectUpdateOne) AddSubjectDigestIDs(ids ...uuid.UUID) *SubjectUpdateOne {
	suo.mutation.AddSubjectDigestIDs(ids...)
	return suo
}

// AddSubjectDigests adds the "subject_digests" edges to the SubjectDigest entity.
func (suo *SubjectUpdateOne) AddSubjectDigests(s ...*SubjectDigest) *SubjectUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddSubjectDigestIDs(ids...)
}

// SetStatementID sets the "statement" edge to the Statement entity by ID.
func (suo *SubjectUpdateOne) SetStatementID(id uuid.UUID) *SubjectUpdateOne {
	suo.mutation.SetStatementID(id)
	return suo
}

// SetNillableStatementID sets the "statement" edge to the Statement entity by ID if the given value is not nil.
func (suo *SubjectUpdateOne) SetNillableStatementID(id *uuid.UUID) *SubjectUpdateOne {
	if id != nil {
		suo = suo.SetStatementID(*id)
	}
	return suo
}

// SetStatement sets the "statement" edge to the Statement entity.
func (suo *SubjectUpdateOne) SetStatement(s *Statement) *SubjectUpdateOne {
	return suo.SetStatementID(s.ID)
}

// Mutation returns the SubjectMutation object of the builder.
func (suo *SubjectUpdateOne) Mutation() *SubjectMutation {
	return suo.mutation
}

// ClearSubjectDigests clears all "subject_digests" edges to the SubjectDigest entity.
func (suo *SubjectUpdateOne) ClearSubjectDigests() *SubjectUpdateOne {
	suo.mutation.ClearSubjectDigests()
	return suo
}

// RemoveSubjectDigestIDs removes the "subject_digests" edge to SubjectDigest entities by IDs.
func (suo *SubjectUpdateOne) RemoveSubjectDigestIDs(ids ...uuid.UUID) *SubjectUpdateOne {
	suo.mutation.RemoveSubjectDigestIDs(ids...)
	return suo
}

// RemoveSubjectDigests removes "subject_digests" edges to SubjectDigest entities.
func (suo *SubjectUpdateOne) RemoveSubjectDigests(s ...*SubjectDigest) *SubjectUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveSubjectDigestIDs(ids...)
}

// ClearStatement clears the "statement" edge to the Statement entity.
func (suo *SubjectUpdateOne) ClearStatement() *SubjectUpdateOne {
	suo.mutation.ClearStatement()
	return suo
}

// Where appends a list predicates to the SubjectUpdate builder.
func (suo *SubjectUpdateOne) Where(ps ...predicate.Subject) *SubjectUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubjectUpdateOne) Select(field string, fields ...string) *SubjectUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Subject entity.
func (suo *SubjectUpdateOne) Save(ctx context.Context) (*Subject, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubjectUpdateOne) SaveX(ctx context.Context) *Subject {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SubjectUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubjectUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SubjectUpdateOne) check() error {
	if v, ok := suo.mutation.Name(); ok {
		if err := subject.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Subject.name": %w`, err)}
		}
	}
	return nil
}

func (suo *SubjectUpdateOne) sqlSave(ctx context.Context) (_node *Subject, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(subject.Table, subject.Columns, sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Subject.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subject.FieldID)
		for _, f := range fields {
			if !subject.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(subject.FieldName, field.TypeString, value)
	}
	if suo.mutation.SubjectDigestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.SubjectDigestsTable,
			Columns: []string{subject.SubjectDigestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectdigest.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedSubjectDigestsIDs(); len(nodes) > 0 && !suo.mutation.SubjectDigestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.SubjectDigestsTable,
			Columns: []string{subject.SubjectDigestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectdigest.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SubjectDigestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.SubjectDigestsTable,
			Columns: []string{subject.SubjectDigestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectdigest.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.StatementCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.StatementTable,
			Columns: []string{subject.StatementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statement.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.StatementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.StatementTable,
			Columns: []string{subject.StatementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statement.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Subject{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
