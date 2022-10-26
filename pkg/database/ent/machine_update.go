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
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/alert"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/machine"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/predicate"
)

// MachineUpdate is the builder for updating Machine entities.
type MachineUpdate struct {
	config
	hooks    []Hook
	mutation *MachineMutation
}

// Where appends a list predicates to the MachineUpdate builder.
func (mu *MachineUpdate) Where(ps ...predicate.Machine) *MachineUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetCreatedAt sets the "created_at" field.
func (mu *MachineUpdate) SetCreatedAt(t time.Time) *MachineUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (mu *MachineUpdate) ClearCreatedAt() *MachineUpdate {
	mu.mutation.ClearCreatedAt()
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MachineUpdate) SetUpdatedAt(t time.Time) *MachineUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (mu *MachineUpdate) ClearUpdatedAt() *MachineUpdate {
	mu.mutation.ClearUpdatedAt()
	return mu
}

// SetLastPush sets the "last_push" field.
func (mu *MachineUpdate) SetLastPush(t time.Time) *MachineUpdate {
	mu.mutation.SetLastPush(t)
	return mu
}

// ClearLastPush clears the value of the "last_push" field.
func (mu *MachineUpdate) ClearLastPush() *MachineUpdate {
	mu.mutation.ClearLastPush()
	return mu
}

// SetLastHeartbeat sets the "last_heartbeat" field.
func (mu *MachineUpdate) SetLastHeartbeat(t time.Time) *MachineUpdate {
	mu.mutation.SetLastHeartbeat(t)
	return mu
}

// ClearLastHeartbeat clears the value of the "last_heartbeat" field.
func (mu *MachineUpdate) ClearLastHeartbeat() *MachineUpdate {
	mu.mutation.ClearLastHeartbeat()
	return mu
}

// SetMachineId sets the "machineId" field.
func (mu *MachineUpdate) SetMachineId(s string) *MachineUpdate {
	mu.mutation.SetMachineId(s)
	return mu
}

// SetPassword sets the "password" field.
func (mu *MachineUpdate) SetPassword(s string) *MachineUpdate {
	mu.mutation.SetPassword(s)
	return mu
}

// SetIpAddress sets the "ipAddress" field.
func (mu *MachineUpdate) SetIpAddress(s string) *MachineUpdate {
	mu.mutation.SetIpAddress(s)
	return mu
}

// SetScenarios sets the "scenarios" field.
func (mu *MachineUpdate) SetScenarios(s string) *MachineUpdate {
	mu.mutation.SetScenarios(s)
	return mu
}

// SetNillableScenarios sets the "scenarios" field if the given value is not nil.
func (mu *MachineUpdate) SetNillableScenarios(s *string) *MachineUpdate {
	if s != nil {
		mu.SetScenarios(*s)
	}
	return mu
}

// ClearScenarios clears the value of the "scenarios" field.
func (mu *MachineUpdate) ClearScenarios() *MachineUpdate {
	mu.mutation.ClearScenarios()
	return mu
}

// SetVersion sets the "version" field.
func (mu *MachineUpdate) SetVersion(s string) *MachineUpdate {
	mu.mutation.SetVersion(s)
	return mu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (mu *MachineUpdate) SetNillableVersion(s *string) *MachineUpdate {
	if s != nil {
		mu.SetVersion(*s)
	}
	return mu
}

// ClearVersion clears the value of the "version" field.
func (mu *MachineUpdate) ClearVersion() *MachineUpdate {
	mu.mutation.ClearVersion()
	return mu
}

// SetIsValidated sets the "isValidated" field.
func (mu *MachineUpdate) SetIsValidated(b bool) *MachineUpdate {
	mu.mutation.SetIsValidated(b)
	return mu
}

// SetNillableIsValidated sets the "isValidated" field if the given value is not nil.
func (mu *MachineUpdate) SetNillableIsValidated(b *bool) *MachineUpdate {
	if b != nil {
		mu.SetIsValidated(*b)
	}
	return mu
}

// SetStatus sets the "status" field.
func (mu *MachineUpdate) SetStatus(s string) *MachineUpdate {
	mu.mutation.SetStatus(s)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *MachineUpdate) SetNillableStatus(s *string) *MachineUpdate {
	if s != nil {
		mu.SetStatus(*s)
	}
	return mu
}

// ClearStatus clears the value of the "status" field.
func (mu *MachineUpdate) ClearStatus() *MachineUpdate {
	mu.mutation.ClearStatus()
	return mu
}

// SetAuthType sets the "auth_type" field.
func (mu *MachineUpdate) SetAuthType(s string) *MachineUpdate {
	mu.mutation.SetAuthType(s)
	return mu
}

// SetNillableAuthType sets the "auth_type" field if the given value is not nil.
func (mu *MachineUpdate) SetNillableAuthType(s *string) *MachineUpdate {
	if s != nil {
		mu.SetAuthType(*s)
	}
	return mu
}

// AddAlertIDs adds the "alerts" edge to the Alert entity by IDs.
func (mu *MachineUpdate) AddAlertIDs(ids ...int) *MachineUpdate {
	mu.mutation.AddAlertIDs(ids...)
	return mu
}

// AddAlerts adds the "alerts" edges to the Alert entity.
func (mu *MachineUpdate) AddAlerts(a ...*Alert) *MachineUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return mu.AddAlertIDs(ids...)
}

// Mutation returns the MachineMutation object of the builder.
func (mu *MachineUpdate) Mutation() *MachineMutation {
	return mu.mutation
}

// ClearAlerts clears all "alerts" edges to the Alert entity.
func (mu *MachineUpdate) ClearAlerts() *MachineUpdate {
	mu.mutation.ClearAlerts()
	return mu
}

// RemoveAlertIDs removes the "alerts" edge to Alert entities by IDs.
func (mu *MachineUpdate) RemoveAlertIDs(ids ...int) *MachineUpdate {
	mu.mutation.RemoveAlertIDs(ids...)
	return mu
}

// RemoveAlerts removes "alerts" edges to Alert entities.
func (mu *MachineUpdate) RemoveAlerts(a ...*Alert) *MachineUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return mu.RemoveAlertIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MachineUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MachineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MachineUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MachineUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MachineUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MachineUpdate) defaults() {
	if _, ok := mu.mutation.CreatedAt(); !ok && !mu.mutation.CreatedAtCleared() {
		v := machine.UpdateDefaultCreatedAt()
		mu.mutation.SetCreatedAt(v)
	}
	if _, ok := mu.mutation.UpdatedAt(); !ok && !mu.mutation.UpdatedAtCleared() {
		v := machine.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
	if _, ok := mu.mutation.LastPush(); !ok && !mu.mutation.LastPushCleared() {
		v := machine.UpdateDefaultLastPush()
		mu.mutation.SetLastPush(v)
	}
	if _, ok := mu.mutation.LastHeartbeat(); !ok && !mu.mutation.LastHeartbeatCleared() {
		v := machine.UpdateDefaultLastHeartbeat()
		mu.mutation.SetLastHeartbeat(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MachineUpdate) check() error {
	if v, ok := mu.mutation.Scenarios(); ok {
		if err := machine.ScenariosValidator(v); err != nil {
			return &ValidationError{Name: "scenarios", err: fmt.Errorf(`ent: validator failed for field "Machine.scenarios": %w`, err)}
		}
	}
	return nil
}

func (mu *MachineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   machine.Table,
			Columns: machine.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: machine.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: machine.FieldCreatedAt,
		})
	}
	if mu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: machine.FieldCreatedAt,
		})
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: machine.FieldUpdatedAt,
		})
	}
	if mu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: machine.FieldUpdatedAt,
		})
	}
	if value, ok := mu.mutation.LastPush(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: machine.FieldLastPush,
		})
	}
	if mu.mutation.LastPushCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: machine.FieldLastPush,
		})
	}
	if value, ok := mu.mutation.LastHeartbeat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: machine.FieldLastHeartbeat,
		})
	}
	if mu.mutation.LastHeartbeatCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: machine.FieldLastHeartbeat,
		})
	}
	if value, ok := mu.mutation.MachineId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldMachineId,
		})
	}
	if value, ok := mu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldPassword,
		})
	}
	if value, ok := mu.mutation.IpAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldIpAddress,
		})
	}
	if value, ok := mu.mutation.Scenarios(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldScenarios,
		})
	}
	if mu.mutation.ScenariosCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: machine.FieldScenarios,
		})
	}
	if value, ok := mu.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldVersion,
		})
	}
	if mu.mutation.VersionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: machine.FieldVersion,
		})
	}
	if value, ok := mu.mutation.IsValidated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: machine.FieldIsValidated,
		})
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldStatus,
		})
	}
	if mu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: machine.FieldStatus,
		})
	}
	if value, ok := mu.mutation.AuthType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldAuthType,
		})
	}
	if mu.mutation.AlertsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.AlertsTable,
			Columns: []string{machine.AlertsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alert.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedAlertsIDs(); len(nodes) > 0 && !mu.mutation.AlertsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.AlertsTable,
			Columns: []string{machine.AlertsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alert.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.AlertsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.AlertsTable,
			Columns: []string{machine.AlertsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alert.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{machine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MachineUpdateOne is the builder for updating a single Machine entity.
type MachineUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MachineMutation
}

// SetCreatedAt sets the "created_at" field.
func (muo *MachineUpdateOne) SetCreatedAt(t time.Time) *MachineUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (muo *MachineUpdateOne) ClearCreatedAt() *MachineUpdateOne {
	muo.mutation.ClearCreatedAt()
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MachineUpdateOne) SetUpdatedAt(t time.Time) *MachineUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (muo *MachineUpdateOne) ClearUpdatedAt() *MachineUpdateOne {
	muo.mutation.ClearUpdatedAt()
	return muo
}

// SetLastPush sets the "last_push" field.
func (muo *MachineUpdateOne) SetLastPush(t time.Time) *MachineUpdateOne {
	muo.mutation.SetLastPush(t)
	return muo
}

// ClearLastPush clears the value of the "last_push" field.
func (muo *MachineUpdateOne) ClearLastPush() *MachineUpdateOne {
	muo.mutation.ClearLastPush()
	return muo
}

// SetLastHeartbeat sets the "last_heartbeat" field.
func (muo *MachineUpdateOne) SetLastHeartbeat(t time.Time) *MachineUpdateOne {
	muo.mutation.SetLastHeartbeat(t)
	return muo
}

// ClearLastHeartbeat clears the value of the "last_heartbeat" field.
func (muo *MachineUpdateOne) ClearLastHeartbeat() *MachineUpdateOne {
	muo.mutation.ClearLastHeartbeat()
	return muo
}

// SetMachineId sets the "machineId" field.
func (muo *MachineUpdateOne) SetMachineId(s string) *MachineUpdateOne {
	muo.mutation.SetMachineId(s)
	return muo
}

// SetPassword sets the "password" field.
func (muo *MachineUpdateOne) SetPassword(s string) *MachineUpdateOne {
	muo.mutation.SetPassword(s)
	return muo
}

// SetIpAddress sets the "ipAddress" field.
func (muo *MachineUpdateOne) SetIpAddress(s string) *MachineUpdateOne {
	muo.mutation.SetIpAddress(s)
	return muo
}

// SetScenarios sets the "scenarios" field.
func (muo *MachineUpdateOne) SetScenarios(s string) *MachineUpdateOne {
	muo.mutation.SetScenarios(s)
	return muo
}

// SetNillableScenarios sets the "scenarios" field if the given value is not nil.
func (muo *MachineUpdateOne) SetNillableScenarios(s *string) *MachineUpdateOne {
	if s != nil {
		muo.SetScenarios(*s)
	}
	return muo
}

// ClearScenarios clears the value of the "scenarios" field.
func (muo *MachineUpdateOne) ClearScenarios() *MachineUpdateOne {
	muo.mutation.ClearScenarios()
	return muo
}

// SetVersion sets the "version" field.
func (muo *MachineUpdateOne) SetVersion(s string) *MachineUpdateOne {
	muo.mutation.SetVersion(s)
	return muo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (muo *MachineUpdateOne) SetNillableVersion(s *string) *MachineUpdateOne {
	if s != nil {
		muo.SetVersion(*s)
	}
	return muo
}

// ClearVersion clears the value of the "version" field.
func (muo *MachineUpdateOne) ClearVersion() *MachineUpdateOne {
	muo.mutation.ClearVersion()
	return muo
}

// SetIsValidated sets the "isValidated" field.
func (muo *MachineUpdateOne) SetIsValidated(b bool) *MachineUpdateOne {
	muo.mutation.SetIsValidated(b)
	return muo
}

// SetNillableIsValidated sets the "isValidated" field if the given value is not nil.
func (muo *MachineUpdateOne) SetNillableIsValidated(b *bool) *MachineUpdateOne {
	if b != nil {
		muo.SetIsValidated(*b)
	}
	return muo
}

// SetStatus sets the "status" field.
func (muo *MachineUpdateOne) SetStatus(s string) *MachineUpdateOne {
	muo.mutation.SetStatus(s)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *MachineUpdateOne) SetNillableStatus(s *string) *MachineUpdateOne {
	if s != nil {
		muo.SetStatus(*s)
	}
	return muo
}

// ClearStatus clears the value of the "status" field.
func (muo *MachineUpdateOne) ClearStatus() *MachineUpdateOne {
	muo.mutation.ClearStatus()
	return muo
}

// SetAuthType sets the "auth_type" field.
func (muo *MachineUpdateOne) SetAuthType(s string) *MachineUpdateOne {
	muo.mutation.SetAuthType(s)
	return muo
}

// SetNillableAuthType sets the "auth_type" field if the given value is not nil.
func (muo *MachineUpdateOne) SetNillableAuthType(s *string) *MachineUpdateOne {
	if s != nil {
		muo.SetAuthType(*s)
	}
	return muo
}

// AddAlertIDs adds the "alerts" edge to the Alert entity by IDs.
func (muo *MachineUpdateOne) AddAlertIDs(ids ...int) *MachineUpdateOne {
	muo.mutation.AddAlertIDs(ids...)
	return muo
}

// AddAlerts adds the "alerts" edges to the Alert entity.
func (muo *MachineUpdateOne) AddAlerts(a ...*Alert) *MachineUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return muo.AddAlertIDs(ids...)
}

// Mutation returns the MachineMutation object of the builder.
func (muo *MachineUpdateOne) Mutation() *MachineMutation {
	return muo.mutation
}

// ClearAlerts clears all "alerts" edges to the Alert entity.
func (muo *MachineUpdateOne) ClearAlerts() *MachineUpdateOne {
	muo.mutation.ClearAlerts()
	return muo
}

// RemoveAlertIDs removes the "alerts" edge to Alert entities by IDs.
func (muo *MachineUpdateOne) RemoveAlertIDs(ids ...int) *MachineUpdateOne {
	muo.mutation.RemoveAlertIDs(ids...)
	return muo
}

// RemoveAlerts removes "alerts" edges to Alert entities.
func (muo *MachineUpdateOne) RemoveAlerts(a ...*Alert) *MachineUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return muo.RemoveAlertIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MachineUpdateOne) Select(field string, fields ...string) *MachineUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Machine entity.
func (muo *MachineUpdateOne) Save(ctx context.Context) (*Machine, error) {
	var (
		err  error
		node *Machine
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MachineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Machine)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MachineMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MachineUpdateOne) SaveX(ctx context.Context) *Machine {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MachineUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MachineUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MachineUpdateOne) defaults() {
	if _, ok := muo.mutation.CreatedAt(); !ok && !muo.mutation.CreatedAtCleared() {
		v := machine.UpdateDefaultCreatedAt()
		muo.mutation.SetCreatedAt(v)
	}
	if _, ok := muo.mutation.UpdatedAt(); !ok && !muo.mutation.UpdatedAtCleared() {
		v := machine.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
	if _, ok := muo.mutation.LastPush(); !ok && !muo.mutation.LastPushCleared() {
		v := machine.UpdateDefaultLastPush()
		muo.mutation.SetLastPush(v)
	}
	if _, ok := muo.mutation.LastHeartbeat(); !ok && !muo.mutation.LastHeartbeatCleared() {
		v := machine.UpdateDefaultLastHeartbeat()
		muo.mutation.SetLastHeartbeat(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MachineUpdateOne) check() error {
	if v, ok := muo.mutation.Scenarios(); ok {
		if err := machine.ScenariosValidator(v); err != nil {
			return &ValidationError{Name: "scenarios", err: fmt.Errorf(`ent: validator failed for field "Machine.scenarios": %w`, err)}
		}
	}
	return nil
}

func (muo *MachineUpdateOne) sqlSave(ctx context.Context) (_node *Machine, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   machine.Table,
			Columns: machine.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: machine.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Machine.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, machine.FieldID)
		for _, f := range fields {
			if !machine.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != machine.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: machine.FieldCreatedAt,
		})
	}
	if muo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: machine.FieldCreatedAt,
		})
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: machine.FieldUpdatedAt,
		})
	}
	if muo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: machine.FieldUpdatedAt,
		})
	}
	if value, ok := muo.mutation.LastPush(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: machine.FieldLastPush,
		})
	}
	if muo.mutation.LastPushCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: machine.FieldLastPush,
		})
	}
	if value, ok := muo.mutation.LastHeartbeat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: machine.FieldLastHeartbeat,
		})
	}
	if muo.mutation.LastHeartbeatCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: machine.FieldLastHeartbeat,
		})
	}
	if value, ok := muo.mutation.MachineId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldMachineId,
		})
	}
	if value, ok := muo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldPassword,
		})
	}
	if value, ok := muo.mutation.IpAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldIpAddress,
		})
	}
	if value, ok := muo.mutation.Scenarios(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldScenarios,
		})
	}
	if muo.mutation.ScenariosCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: machine.FieldScenarios,
		})
	}
	if value, ok := muo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldVersion,
		})
	}
	if muo.mutation.VersionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: machine.FieldVersion,
		})
	}
	if value, ok := muo.mutation.IsValidated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: machine.FieldIsValidated,
		})
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldStatus,
		})
	}
	if muo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: machine.FieldStatus,
		})
	}
	if value, ok := muo.mutation.AuthType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldAuthType,
		})
	}
	if muo.mutation.AlertsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.AlertsTable,
			Columns: []string{machine.AlertsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alert.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedAlertsIDs(); len(nodes) > 0 && !muo.mutation.AlertsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.AlertsTable,
			Columns: []string{machine.AlertsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alert.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.AlertsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   machine.AlertsTable,
			Columns: []string{machine.AlertsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alert.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Machine{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{machine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
