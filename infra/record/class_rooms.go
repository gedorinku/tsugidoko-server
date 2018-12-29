// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package record

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// ClassRoom is an object representing the database table.
type ClassRoom struct {
	ID         int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name       string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Latitude   float64   `boil:"latitude" json:"latitude" toml:"latitude" yaml:"latitude"`
	Longitude  float64   `boil:"longitude" json:"longitude" toml:"longitude" yaml:"longitude"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	BuildingID int       `boil:"building_id" json:"building_id" toml:"building_id" yaml:"building_id"`

	R *classRoomR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L classRoomL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ClassRoomColumns = struct {
	ID         string
	Name       string
	Latitude   string
	Longitude  string
	CreatedAt  string
	UpdatedAt  string
	BuildingID string
}{
	ID:         "id",
	Name:       "name",
	Latitude:   "latitude",
	Longitude:  "longitude",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	BuildingID: "building_id",
}

// ClassRoomRels is where relationship names are stored.
var ClassRoomRels = struct {
	Beacons       string
	UserPositions string
}{
	Beacons:       "Beacons",
	UserPositions: "UserPositions",
}

// classRoomR is where relationships are stored.
type classRoomR struct {
	Beacons       BeaconSlice
	UserPositions UserPositionSlice
}

// NewStruct creates a new relationship struct
func (*classRoomR) NewStruct() *classRoomR {
	return &classRoomR{}
}

// classRoomL is where Load methods for each relationship are stored.
type classRoomL struct{}

var (
	classRoomColumns               = []string{"id", "name", "latitude", "longitude", "created_at", "updated_at", "building_id"}
	classRoomColumnsWithoutDefault = []string{"name", "latitude", "longitude", "created_at", "updated_at", "building_id"}
	classRoomColumnsWithDefault    = []string{"id"}
	classRoomPrimaryKeyColumns     = []string{"id"}
)

type (
	// ClassRoomSlice is an alias for a slice of pointers to ClassRoom.
	// This should generally be used opposed to []ClassRoom.
	ClassRoomSlice []*ClassRoom
	// ClassRoomHook is the signature for custom ClassRoom hook methods
	ClassRoomHook func(context.Context, boil.ContextExecutor, *ClassRoom) error

	classRoomQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	classRoomType                 = reflect.TypeOf(&ClassRoom{})
	classRoomMapping              = queries.MakeStructMapping(classRoomType)
	classRoomPrimaryKeyMapping, _ = queries.BindMapping(classRoomType, classRoomMapping, classRoomPrimaryKeyColumns)
	classRoomInsertCacheMut       sync.RWMutex
	classRoomInsertCache          = make(map[string]insertCache)
	classRoomUpdateCacheMut       sync.RWMutex
	classRoomUpdateCache          = make(map[string]updateCache)
	classRoomUpsertCacheMut       sync.RWMutex
	classRoomUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var classRoomBeforeInsertHooks []ClassRoomHook
var classRoomBeforeUpdateHooks []ClassRoomHook
var classRoomBeforeDeleteHooks []ClassRoomHook
var classRoomBeforeUpsertHooks []ClassRoomHook

var classRoomAfterInsertHooks []ClassRoomHook
var classRoomAfterSelectHooks []ClassRoomHook
var classRoomAfterUpdateHooks []ClassRoomHook
var classRoomAfterDeleteHooks []ClassRoomHook
var classRoomAfterUpsertHooks []ClassRoomHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ClassRoom) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range classRoomBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ClassRoom) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range classRoomBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ClassRoom) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range classRoomBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ClassRoom) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range classRoomBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ClassRoom) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range classRoomAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ClassRoom) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range classRoomAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ClassRoom) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range classRoomAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ClassRoom) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range classRoomAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ClassRoom) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range classRoomAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddClassRoomHook registers your hook function for all future operations.
func AddClassRoomHook(hookPoint boil.HookPoint, classRoomHook ClassRoomHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		classRoomBeforeInsertHooks = append(classRoomBeforeInsertHooks, classRoomHook)
	case boil.BeforeUpdateHook:
		classRoomBeforeUpdateHooks = append(classRoomBeforeUpdateHooks, classRoomHook)
	case boil.BeforeDeleteHook:
		classRoomBeforeDeleteHooks = append(classRoomBeforeDeleteHooks, classRoomHook)
	case boil.BeforeUpsertHook:
		classRoomBeforeUpsertHooks = append(classRoomBeforeUpsertHooks, classRoomHook)
	case boil.AfterInsertHook:
		classRoomAfterInsertHooks = append(classRoomAfterInsertHooks, classRoomHook)
	case boil.AfterSelectHook:
		classRoomAfterSelectHooks = append(classRoomAfterSelectHooks, classRoomHook)
	case boil.AfterUpdateHook:
		classRoomAfterUpdateHooks = append(classRoomAfterUpdateHooks, classRoomHook)
	case boil.AfterDeleteHook:
		classRoomAfterDeleteHooks = append(classRoomAfterDeleteHooks, classRoomHook)
	case boil.AfterUpsertHook:
		classRoomAfterUpsertHooks = append(classRoomAfterUpsertHooks, classRoomHook)
	}
}

// One returns a single classRoom record from the query.
func (q classRoomQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ClassRoom, error) {
	o := &ClassRoom{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "record: failed to execute a one query for class_rooms")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ClassRoom records from the query.
func (q classRoomQuery) All(ctx context.Context, exec boil.ContextExecutor) (ClassRoomSlice, error) {
	var o []*ClassRoom

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "record: failed to assign all query results to ClassRoom slice")
	}

	if len(classRoomAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ClassRoom records in the query.
func (q classRoomQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to count class_rooms rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q classRoomQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "record: failed to check if class_rooms exists")
	}

	return count > 0, nil
}

// Beacons retrieves all the beacon's Beacons with an executor.
func (o *ClassRoom) Beacons(mods ...qm.QueryMod) beaconQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"beacons\".\"class_room_id\"=?", o.ID),
	)

	query := Beacons(queryMods...)
	queries.SetFrom(query.Query, "\"beacons\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"beacons\".*"})
	}

	return query
}

// UserPositions retrieves all the user_position's UserPositions with an executor.
func (o *ClassRoom) UserPositions(mods ...qm.QueryMod) userPositionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_positions\".\"class_room_id\"=?", o.ID),
	)

	query := UserPositions(queryMods...)
	queries.SetFrom(query.Query, "\"user_positions\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"user_positions\".*"})
	}

	return query
}

// LoadBeacons allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (classRoomL) LoadBeacons(ctx context.Context, e boil.ContextExecutor, singular bool, maybeClassRoom interface{}, mods queries.Applicator) error {
	var slice []*ClassRoom
	var object *ClassRoom

	if singular {
		object = maybeClassRoom.(*ClassRoom)
	} else {
		slice = *maybeClassRoom.(*[]*ClassRoom)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &classRoomR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &classRoomR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`beacons`), qm.WhereIn(`class_room_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load beacons")
	}

	var resultSlice []*Beacon
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice beacons")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on beacons")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for beacons")
	}

	if len(beaconAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Beacons = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &beaconR{}
			}
			foreign.R.ClassRoom = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ClassRoomID {
				local.R.Beacons = append(local.R.Beacons, foreign)
				if foreign.R == nil {
					foreign.R = &beaconR{}
				}
				foreign.R.ClassRoom = local
				break
			}
		}
	}

	return nil
}

// LoadUserPositions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (classRoomL) LoadUserPositions(ctx context.Context, e boil.ContextExecutor, singular bool, maybeClassRoom interface{}, mods queries.Applicator) error {
	var slice []*ClassRoom
	var object *ClassRoom

	if singular {
		object = maybeClassRoom.(*ClassRoom)
	} else {
		slice = *maybeClassRoom.(*[]*ClassRoom)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &classRoomR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &classRoomR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`user_positions`), qm.WhereIn(`class_room_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_positions")
	}

	var resultSlice []*UserPosition
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_positions")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_positions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_positions")
	}

	if len(userPositionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserPositions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userPositionR{}
			}
			foreign.R.ClassRoom = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ClassRoomID {
				local.R.UserPositions = append(local.R.UserPositions, foreign)
				if foreign.R == nil {
					foreign.R = &userPositionR{}
				}
				foreign.R.ClassRoom = local
				break
			}
		}
	}

	return nil
}

// AddBeacons adds the given related objects to the existing relationships
// of the class_room, optionally inserting them as new records.
// Appends related to o.R.Beacons.
// Sets related.R.ClassRoom appropriately.
func (o *ClassRoom) AddBeacons(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Beacon) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ClassRoomID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"beacons\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"class_room_id"}),
				strmangle.WhereClause("\"", "\"", 2, beaconPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ClassRoomID = o.ID
		}
	}

	if o.R == nil {
		o.R = &classRoomR{
			Beacons: related,
		}
	} else {
		o.R.Beacons = append(o.R.Beacons, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &beaconR{
				ClassRoom: o,
			}
		} else {
			rel.R.ClassRoom = o
		}
	}
	return nil
}

// AddUserPositions adds the given related objects to the existing relationships
// of the class_room, optionally inserting them as new records.
// Appends related to o.R.UserPositions.
// Sets related.R.ClassRoom appropriately.
func (o *ClassRoom) AddUserPositions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserPosition) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ClassRoomID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_positions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"class_room_id"}),
				strmangle.WhereClause("\"", "\"", 2, userPositionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ClassRoomID = o.ID
		}
	}

	if o.R == nil {
		o.R = &classRoomR{
			UserPositions: related,
		}
	} else {
		o.R.UserPositions = append(o.R.UserPositions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userPositionR{
				ClassRoom: o,
			}
		} else {
			rel.R.ClassRoom = o
		}
	}
	return nil
}

// ClassRooms retrieves all the records using an executor.
func ClassRooms(mods ...qm.QueryMod) classRoomQuery {
	mods = append(mods, qm.From("\"class_rooms\""))
	return classRoomQuery{NewQuery(mods...)}
}

// FindClassRoom retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindClassRoom(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*ClassRoom, error) {
	classRoomObj := &ClassRoom{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"class_rooms\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, classRoomObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "record: unable to select from class_rooms")
	}

	return classRoomObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ClassRoom) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("record: no class_rooms provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(classRoomColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	classRoomInsertCacheMut.RLock()
	cache, cached := classRoomInsertCache[key]
	classRoomInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			classRoomColumns,
			classRoomColumnsWithDefault,
			classRoomColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(classRoomType, classRoomMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(classRoomType, classRoomMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"class_rooms\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"class_rooms\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "record: unable to insert into class_rooms")
	}

	if !cached {
		classRoomInsertCacheMut.Lock()
		classRoomInsertCache[key] = cache
		classRoomInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ClassRoom.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ClassRoom) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	classRoomUpdateCacheMut.RLock()
	cache, cached := classRoomUpdateCache[key]
	classRoomUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			classRoomColumns,
			classRoomPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("record: unable to update class_rooms, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"class_rooms\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, classRoomPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(classRoomType, classRoomMapping, append(wl, classRoomPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update class_rooms row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by update for class_rooms")
	}

	if !cached {
		classRoomUpdateCacheMut.Lock()
		classRoomUpdateCache[key] = cache
		classRoomUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q classRoomQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update all for class_rooms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to retrieve rows affected for class_rooms")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ClassRoomSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("record: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), classRoomPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"class_rooms\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, classRoomPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update all in classRoom slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to retrieve rows affected all in update all classRoom")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ClassRoom) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("record: no class_rooms provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(classRoomColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	classRoomUpsertCacheMut.RLock()
	cache, cached := classRoomUpsertCache[key]
	classRoomUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			classRoomColumns,
			classRoomColumnsWithDefault,
			classRoomColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			classRoomColumns,
			classRoomPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("record: unable to upsert class_rooms, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(classRoomPrimaryKeyColumns))
			copy(conflict, classRoomPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"class_rooms\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(classRoomType, classRoomMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(classRoomType, classRoomMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "record: unable to upsert class_rooms")
	}

	if !cached {
		classRoomUpsertCacheMut.Lock()
		classRoomUpsertCache[key] = cache
		classRoomUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ClassRoom record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ClassRoom) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("record: no ClassRoom provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), classRoomPrimaryKeyMapping)
	sql := "DELETE FROM \"class_rooms\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete from class_rooms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by delete for class_rooms")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q classRoomQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("record: no classRoomQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete all from class_rooms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by deleteall for class_rooms")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ClassRoomSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("record: no ClassRoom slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(classRoomBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), classRoomPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"class_rooms\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, classRoomPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete all from classRoom slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by deleteall for class_rooms")
	}

	if len(classRoomAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ClassRoom) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindClassRoom(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ClassRoomSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ClassRoomSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), classRoomPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"class_rooms\".* FROM \"class_rooms\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, classRoomPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "record: unable to reload all in ClassRoomSlice")
	}

	*o = slice

	return nil
}

// ClassRoomExists checks if the ClassRoom row exists.
func ClassRoomExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"class_rooms\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "record: unable to check if class_rooms exists")
	}

	return exists, nil
}