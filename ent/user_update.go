// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/kcarretto/paragon/ent/event"
	"github.com/kcarretto/paragon/ent/job"
	"github.com/kcarretto/paragon/ent/predicate"
	"github.com/kcarretto/paragon/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	Name *string

	PhotoURL          *string
	SessionToken      *string
	clearSessionToken bool
	IsActivated       *bool
	IsAdmin           *bool
	jobs              map[int]struct{}
	events            map[int]struct{}
	removedJobs       map[int]struct{}
	removedEvents     map[int]struct{}
	predicates        []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetName sets the Name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.Name = &s
	return uu
}

// SetPhotoURL sets the PhotoURL field.
func (uu *UserUpdate) SetPhotoURL(s string) *UserUpdate {
	uu.PhotoURL = &s
	return uu
}

// SetSessionToken sets the SessionToken field.
func (uu *UserUpdate) SetSessionToken(s string) *UserUpdate {
	uu.SessionToken = &s
	return uu
}

// SetNillableSessionToken sets the SessionToken field if the given value is not nil.
func (uu *UserUpdate) SetNillableSessionToken(s *string) *UserUpdate {
	if s != nil {
		uu.SetSessionToken(*s)
	}
	return uu
}

// ClearSessionToken clears the value of SessionToken.
func (uu *UserUpdate) ClearSessionToken() *UserUpdate {
	uu.SessionToken = nil
	uu.clearSessionToken = true
	return uu
}

// SetIsActivated sets the IsActivated field.
func (uu *UserUpdate) SetIsActivated(b bool) *UserUpdate {
	uu.IsActivated = &b
	return uu
}

// SetNillableIsActivated sets the IsActivated field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsActivated(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsActivated(*b)
	}
	return uu
}

// SetIsAdmin sets the IsAdmin field.
func (uu *UserUpdate) SetIsAdmin(b bool) *UserUpdate {
	uu.IsAdmin = &b
	return uu
}

// SetNillableIsAdmin sets the IsAdmin field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsAdmin(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsAdmin(*b)
	}
	return uu
}

// AddJobIDs adds the jobs edge to Job by ids.
func (uu *UserUpdate) AddJobIDs(ids ...int) *UserUpdate {
	if uu.jobs == nil {
		uu.jobs = make(map[int]struct{})
	}
	for i := range ids {
		uu.jobs[ids[i]] = struct{}{}
	}
	return uu
}

// AddJobs adds the jobs edges to Job.
func (uu *UserUpdate) AddJobs(j ...*Job) *UserUpdate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return uu.AddJobIDs(ids...)
}

// AddEventIDs adds the events edge to Event by ids.
func (uu *UserUpdate) AddEventIDs(ids ...int) *UserUpdate {
	if uu.events == nil {
		uu.events = make(map[int]struct{})
	}
	for i := range ids {
		uu.events[ids[i]] = struct{}{}
	}
	return uu
}

// AddEvents adds the events edges to Event.
func (uu *UserUpdate) AddEvents(e ...*Event) *UserUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.AddEventIDs(ids...)
}

// RemoveJobIDs removes the jobs edge to Job by ids.
func (uu *UserUpdate) RemoveJobIDs(ids ...int) *UserUpdate {
	if uu.removedJobs == nil {
		uu.removedJobs = make(map[int]struct{})
	}
	for i := range ids {
		uu.removedJobs[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveJobs removes jobs edges to Job.
func (uu *UserUpdate) RemoveJobs(j ...*Job) *UserUpdate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return uu.RemoveJobIDs(ids...)
}

// RemoveEventIDs removes the events edge to Event by ids.
func (uu *UserUpdate) RemoveEventIDs(ids ...int) *UserUpdate {
	if uu.removedEvents == nil {
		uu.removedEvents = make(map[int]struct{})
	}
	for i := range ids {
		uu.removedEvents[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveEvents removes events edges to Event.
func (uu *UserUpdate) RemoveEvents(e ...*Event) *UserUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.RemoveEventIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if uu.Name != nil {
		if err := user.NameValidator(*uu.Name); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"Name\": %v", err)
		}
	}
	return uu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := uu.Name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
	}
	if value := uu.PhotoURL; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldPhotoURL,
		})
	}
	if value := uu.SessionToken; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldSessionToken,
		})
	}
	if uu.clearSessionToken {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldSessionToken,
		})
	}
	if value := uu.IsActivated; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: user.FieldIsActivated,
		})
	}
	if value := uu.IsAdmin; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: user.FieldIsAdmin,
		})
	}
	if nodes := uu.removedJobs; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.JobsTable,
			Columns: []string{user.JobsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.jobs; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.JobsTable,
			Columns: []string{user.JobsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uu.removedEvents; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EventsTable,
			Columns: []string{user.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.events; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EventsTable,
			Columns: []string{user.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	id   int
	Name *string

	PhotoURL          *string
	SessionToken      *string
	clearSessionToken bool
	IsActivated       *bool
	IsAdmin           *bool
	jobs              map[int]struct{}
	events            map[int]struct{}
	removedJobs       map[int]struct{}
	removedEvents     map[int]struct{}
}

// SetName sets the Name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.Name = &s
	return uuo
}

// SetPhotoURL sets the PhotoURL field.
func (uuo *UserUpdateOne) SetPhotoURL(s string) *UserUpdateOne {
	uuo.PhotoURL = &s
	return uuo
}

// SetSessionToken sets the SessionToken field.
func (uuo *UserUpdateOne) SetSessionToken(s string) *UserUpdateOne {
	uuo.SessionToken = &s
	return uuo
}

// SetNillableSessionToken sets the SessionToken field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSessionToken(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSessionToken(*s)
	}
	return uuo
}

// ClearSessionToken clears the value of SessionToken.
func (uuo *UserUpdateOne) ClearSessionToken() *UserUpdateOne {
	uuo.SessionToken = nil
	uuo.clearSessionToken = true
	return uuo
}

// SetIsActivated sets the IsActivated field.
func (uuo *UserUpdateOne) SetIsActivated(b bool) *UserUpdateOne {
	uuo.IsActivated = &b
	return uuo
}

// SetNillableIsActivated sets the IsActivated field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsActivated(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsActivated(*b)
	}
	return uuo
}

// SetIsAdmin sets the IsAdmin field.
func (uuo *UserUpdateOne) SetIsAdmin(b bool) *UserUpdateOne {
	uuo.IsAdmin = &b
	return uuo
}

// SetNillableIsAdmin sets the IsAdmin field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsAdmin(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsAdmin(*b)
	}
	return uuo
}

// AddJobIDs adds the jobs edge to Job by ids.
func (uuo *UserUpdateOne) AddJobIDs(ids ...int) *UserUpdateOne {
	if uuo.jobs == nil {
		uuo.jobs = make(map[int]struct{})
	}
	for i := range ids {
		uuo.jobs[ids[i]] = struct{}{}
	}
	return uuo
}

// AddJobs adds the jobs edges to Job.
func (uuo *UserUpdateOne) AddJobs(j ...*Job) *UserUpdateOne {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return uuo.AddJobIDs(ids...)
}

// AddEventIDs adds the events edge to Event by ids.
func (uuo *UserUpdateOne) AddEventIDs(ids ...int) *UserUpdateOne {
	if uuo.events == nil {
		uuo.events = make(map[int]struct{})
	}
	for i := range ids {
		uuo.events[ids[i]] = struct{}{}
	}
	return uuo
}

// AddEvents adds the events edges to Event.
func (uuo *UserUpdateOne) AddEvents(e ...*Event) *UserUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.AddEventIDs(ids...)
}

// RemoveJobIDs removes the jobs edge to Job by ids.
func (uuo *UserUpdateOne) RemoveJobIDs(ids ...int) *UserUpdateOne {
	if uuo.removedJobs == nil {
		uuo.removedJobs = make(map[int]struct{})
	}
	for i := range ids {
		uuo.removedJobs[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveJobs removes jobs edges to Job.
func (uuo *UserUpdateOne) RemoveJobs(j ...*Job) *UserUpdateOne {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return uuo.RemoveJobIDs(ids...)
}

// RemoveEventIDs removes the events edge to Event by ids.
func (uuo *UserUpdateOne) RemoveEventIDs(ids ...int) *UserUpdateOne {
	if uuo.removedEvents == nil {
		uuo.removedEvents = make(map[int]struct{})
	}
	for i := range ids {
		uuo.removedEvents[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveEvents removes events edges to Event.
func (uuo *UserUpdateOne) RemoveEvents(e ...*Event) *UserUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.RemoveEventIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if uuo.Name != nil {
		if err := user.NameValidator(*uuo.Name); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"Name\": %v", err)
		}
	}
	return uuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  uuo.id,
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if value := uuo.Name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
	}
	if value := uuo.PhotoURL; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldPhotoURL,
		})
	}
	if value := uuo.SessionToken; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldSessionToken,
		})
	}
	if uuo.clearSessionToken {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldSessionToken,
		})
	}
	if value := uuo.IsActivated; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: user.FieldIsActivated,
		})
	}
	if value := uuo.IsAdmin; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: user.FieldIsAdmin,
		})
	}
	if nodes := uuo.removedJobs; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.JobsTable,
			Columns: []string{user.JobsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.jobs; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.JobsTable,
			Columns: []string{user.JobsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uuo.removedEvents; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EventsTable,
			Columns: []string{user.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.events; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EventsTable,
			Columns: []string{user.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
