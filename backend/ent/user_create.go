// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pattapong1/app/ent/club"
	"github.com/pattapong1/app/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetUSERNAME sets the USER_NAME field.
func (uc *UserCreate) SetUSERNAME(s string) *UserCreate {
	uc.mutation.SetUSERNAME(s)
	return uc
}

// SetUSEREMAIL sets the USER_EMAIL field.
func (uc *UserCreate) SetUSEREMAIL(s string) *UserCreate {
	uc.mutation.SetUSEREMAIL(s)
	return uc
}

// AddClubIDs adds the club edge to Club by ids.
func (uc *UserCreate) AddClubIDs(ids ...int) *UserCreate {
	uc.mutation.AddClubIDs(ids...)
	return uc
}

// AddClub adds the club edges to Club.
func (uc *UserCreate) AddClub(c ...*Club) *UserCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddClubIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if _, ok := uc.mutation.USERNAME(); !ok {
		return nil, &ValidationError{Name: "USER_NAME", err: errors.New("ent: missing required field \"USER_NAME\"")}
	}
	if v, ok := uc.mutation.USERNAME(); ok {
		if err := user.USERNAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "USER_NAME", err: fmt.Errorf("ent: validator failed for field \"USER_NAME\": %w", err)}
		}
	}
	if _, ok := uc.mutation.USEREMAIL(); !ok {
		return nil, &ValidationError{Name: "USER_EMAIL", err: errors.New("ent: missing required field \"USER_EMAIL\"")}
	}
	if v, ok := uc.mutation.USEREMAIL(); ok {
		if err := user.USEREMAILValidator(v); err != nil {
			return nil, &ValidationError{Name: "USER_EMAIL", err: fmt.Errorf("ent: validator failed for field \"USER_EMAIL\": %w", err)}
		}
	}
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	u, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.USERNAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUSERNAME,
		})
		u.USERNAME = value
	}
	if value, ok := uc.mutation.USEREMAIL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUSEREMAIL,
		})
		u.USEREMAIL = value
	}
	if nodes := uc.mutation.ClubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ClubTable,
			Columns: []string{user.ClubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return u, _spec
}
