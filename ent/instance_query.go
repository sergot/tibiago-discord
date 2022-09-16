// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/sergot/tibiago/ent/instance"
	"github.com/sergot/tibiago/ent/instanceconfig"
	"github.com/sergot/tibiago/ent/predicate"
)

// InstanceQuery is the builder for querying Instance entities.
type InstanceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Instance
	withConfig *InstanceConfigQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InstanceQuery builder.
func (iq *InstanceQuery) Where(ps ...predicate.Instance) *InstanceQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit adds a limit step to the query.
func (iq *InstanceQuery) Limit(limit int) *InstanceQuery {
	iq.limit = &limit
	return iq
}

// Offset adds an offset step to the query.
func (iq *InstanceQuery) Offset(offset int) *InstanceQuery {
	iq.offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *InstanceQuery) Unique(unique bool) *InstanceQuery {
	iq.unique = &unique
	return iq
}

// Order adds an order step to the query.
func (iq *InstanceQuery) Order(o ...OrderFunc) *InstanceQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryConfig chains the current query on the "config" edge.
func (iq *InstanceQuery) QueryConfig() *InstanceConfigQuery {
	query := &InstanceConfigQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instance.Table, instance.FieldID, selector),
			sqlgraph.To(instanceconfig.Table, instanceconfig.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, instance.ConfigTable, instance.ConfigColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Instance entity from the query.
// Returns a *NotFoundError when no Instance was found.
func (iq *InstanceQuery) First(ctx context.Context) (*Instance, error) {
	nodes, err := iq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{instance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InstanceQuery) FirstX(ctx context.Context) *Instance {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Instance ID from the query.
// Returns a *NotFoundError when no Instance ID was found.
func (iq *InstanceQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{instance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *InstanceQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Instance entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Instance entity is found.
// Returns a *NotFoundError when no Instance entities are found.
func (iq *InstanceQuery) Only(ctx context.Context) (*Instance, error) {
	nodes, err := iq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{instance.Label}
	default:
		return nil, &NotSingularError{instance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InstanceQuery) OnlyX(ctx context.Context) *Instance {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Instance ID in the query.
// Returns a *NotSingularError when more than one Instance ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *InstanceQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{instance.Label}
	default:
		err = &NotSingularError{instance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InstanceQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Instances.
func (iq *InstanceQuery) All(ctx context.Context) ([]*Instance, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iq *InstanceQuery) AllX(ctx context.Context) []*Instance {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Instance IDs.
func (iq *InstanceQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := iq.Select(instance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InstanceQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InstanceQuery) Count(ctx context.Context) (int, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InstanceQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InstanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InstanceQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InstanceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InstanceQuery) Clone() *InstanceQuery {
	if iq == nil {
		return nil
	}
	return &InstanceQuery{
		config:     iq.config,
		limit:      iq.limit,
		offset:     iq.offset,
		order:      append([]OrderFunc{}, iq.order...),
		predicates: append([]predicate.Instance{}, iq.predicates...),
		withConfig: iq.withConfig.Clone(),
		// clone intermediate query.
		sql:    iq.sql.Clone(),
		path:   iq.path,
		unique: iq.unique,
	}
}

// WithConfig tells the query-builder to eager-load the nodes that are connected to
// the "config" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InstanceQuery) WithConfig(opts ...func(*InstanceConfigQuery)) *InstanceQuery {
	query := &InstanceConfigQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withConfig = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SessionID string `json:"session_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Instance.Query().
//		GroupBy(instance.FieldSessionID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *InstanceQuery) GroupBy(field string, fields ...string) *InstanceGroupBy {
	grbuild := &InstanceGroupBy{config: iq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iq.sqlQuery(ctx), nil
	}
	grbuild.label = instance.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SessionID string `json:"session_id,omitempty"`
//	}
//
//	client.Instance.Query().
//		Select(instance.FieldSessionID).
//		Scan(ctx, &v)
func (iq *InstanceQuery) Select(fields ...string) *InstanceSelect {
	iq.fields = append(iq.fields, fields...)
	selbuild := &InstanceSelect{InstanceQuery: iq}
	selbuild.label = instance.Label
	selbuild.flds, selbuild.scan = &iq.fields, selbuild.Scan
	return selbuild
}

func (iq *InstanceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range iq.fields {
		if !instance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InstanceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Instance, error) {
	var (
		nodes       = []*Instance{}
		_spec       = iq.querySpec()
		loadedTypes = [1]bool{
			iq.withConfig != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Instance).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Instance{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withConfig; query != nil {
		if err := iq.loadConfig(ctx, query, nodes,
			func(n *Instance) { n.Edges.Config = []*InstanceConfig{} },
			func(n *Instance, e *InstanceConfig) { n.Edges.Config = append(n.Edges.Config, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *InstanceQuery) loadConfig(ctx context.Context, query *InstanceConfigQuery, nodes []*Instance, init func(*Instance), assign func(*Instance, *InstanceConfig)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Instance)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.InstanceConfig(func(s *sql.Selector) {
		s.Where(sql.InValues(instance.ConfigColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.instance_config
		if fk == nil {
			return fmt.Errorf(`foreign-key "instance_config" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "instance_config" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (iq *InstanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.fields
	if len(iq.fields) > 0 {
		_spec.Unique = iq.unique != nil && *iq.unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InstanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := iq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (iq *InstanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instance.Table,
			Columns: instance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: instance.FieldID,
			},
		},
		From:   iq.sql,
		Unique: true,
	}
	if unique := iq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := iq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, instance.FieldID)
		for i := range fields {
			if fields[i] != instance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InstanceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(instance.Table)
	columns := iq.fields
	if len(columns) == 0 {
		columns = instance.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.unique != nil && *iq.unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InstanceGroupBy is the group-by builder for Instance entities.
type InstanceGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InstanceGroupBy) Aggregate(fns ...AggregateFunc) *InstanceGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the group-by query and scans the result into the given value.
func (igb *InstanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := igb.path(ctx)
	if err != nil {
		return err
	}
	igb.sql = query
	return igb.sqlScan(ctx, v)
}

func (igb *InstanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range igb.fields {
		if !instance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := igb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (igb *InstanceGroupBy) sqlQuery() *sql.Selector {
	selector := igb.sql.Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(igb.fields)+len(igb.fns))
		for _, f := range igb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(igb.fields...)...)
}

// InstanceSelect is the builder for selecting fields of Instance entities.
type InstanceSelect struct {
	*InstanceQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (is *InstanceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	is.sql = is.InstanceQuery.sqlQuery(ctx)
	return is.sqlScan(ctx, v)
}

func (is *InstanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := is.sql.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
