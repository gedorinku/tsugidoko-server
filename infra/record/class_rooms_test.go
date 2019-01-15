// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package record

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testClassRooms(t *testing.T) {
	t.Parallel()

	query := ClassRooms()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testClassRoomsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClassRoom{}
	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClassRooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClassRoomsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClassRoom{}
	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ClassRooms().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClassRooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClassRoomsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClassRoom{}
	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClassRoomSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClassRooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClassRoomsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClassRoom{}
	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ClassRoomExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ClassRoom exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ClassRoomExists to return true, but got false.")
	}
}

func testClassRoomsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClassRoom{}
	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	classRoomFound, err := FindClassRoom(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if classRoomFound == nil {
		t.Error("want a record, got nil")
	}
}

func testClassRoomsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClassRoom{}
	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ClassRooms().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testClassRoomsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClassRoom{}
	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ClassRooms().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testClassRoomsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	classRoomOne := &ClassRoom{}
	classRoomTwo := &ClassRoom{}
	if err = randomize.Struct(seed, classRoomOne, classRoomDBTypes, false, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}
	if err = randomize.Struct(seed, classRoomTwo, classRoomDBTypes, false, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = classRoomOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = classRoomTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClassRooms().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testClassRoomsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	classRoomOne := &ClassRoom{}
	classRoomTwo := &ClassRoom{}
	if err = randomize.Struct(seed, classRoomOne, classRoomDBTypes, false, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}
	if err = randomize.Struct(seed, classRoomTwo, classRoomDBTypes, false, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = classRoomOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = classRoomTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClassRooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func classRoomBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ClassRoom) error {
	*o = ClassRoom{}
	return nil
}

func classRoomAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ClassRoom) error {
	*o = ClassRoom{}
	return nil
}

func classRoomAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ClassRoom) error {
	*o = ClassRoom{}
	return nil
}

func classRoomBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ClassRoom) error {
	*o = ClassRoom{}
	return nil
}

func classRoomAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ClassRoom) error {
	*o = ClassRoom{}
	return nil
}

func classRoomBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ClassRoom) error {
	*o = ClassRoom{}
	return nil
}

func classRoomAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ClassRoom) error {
	*o = ClassRoom{}
	return nil
}

func classRoomBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ClassRoom) error {
	*o = ClassRoom{}
	return nil
}

func classRoomAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ClassRoom) error {
	*o = ClassRoom{}
	return nil
}

func testClassRoomsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ClassRoom{}
	o := &ClassRoom{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, classRoomDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ClassRoom object: %s", err)
	}

	AddClassRoomHook(boil.BeforeInsertHook, classRoomBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	classRoomBeforeInsertHooks = []ClassRoomHook{}

	AddClassRoomHook(boil.AfterInsertHook, classRoomAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	classRoomAfterInsertHooks = []ClassRoomHook{}

	AddClassRoomHook(boil.AfterSelectHook, classRoomAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	classRoomAfterSelectHooks = []ClassRoomHook{}

	AddClassRoomHook(boil.BeforeUpdateHook, classRoomBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	classRoomBeforeUpdateHooks = []ClassRoomHook{}

	AddClassRoomHook(boil.AfterUpdateHook, classRoomAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	classRoomAfterUpdateHooks = []ClassRoomHook{}

	AddClassRoomHook(boil.BeforeDeleteHook, classRoomBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	classRoomBeforeDeleteHooks = []ClassRoomHook{}

	AddClassRoomHook(boil.AfterDeleteHook, classRoomAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	classRoomAfterDeleteHooks = []ClassRoomHook{}

	AddClassRoomHook(boil.BeforeUpsertHook, classRoomBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	classRoomBeforeUpsertHooks = []ClassRoomHook{}

	AddClassRoomHook(boil.AfterUpsertHook, classRoomAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	classRoomAfterUpsertHooks = []ClassRoomHook{}
}

func testClassRoomsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClassRoom{}
	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClassRooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClassRoomsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClassRoom{}
	if err = randomize.Struct(seed, o, classRoomDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(classRoomColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ClassRooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClassRoomToManyBeacons(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ClassRoom
	var b, c Beacon

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, beaconDBTypes, false, beaconColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, beaconDBTypes, false, beaconColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ClassRoomID = a.ID
	c.ClassRoomID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	beacon, err := a.Beacons().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range beacon {
		if v.ClassRoomID == b.ClassRoomID {
			bFound = true
		}
		if v.ClassRoomID == c.ClassRoomID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ClassRoomSlice{&a}
	if err = a.L.LoadBeacons(ctx, tx, false, (*[]*ClassRoom)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Beacons); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Beacons = nil
	if err = a.L.LoadBeacons(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Beacons); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", beacon)
	}
}

func testClassRoomToManyClassRoomTags(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ClassRoom
	var b, c ClassRoomTag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, classRoomTagDBTypes, false, classRoomTagColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, classRoomTagDBTypes, false, classRoomTagColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ClassRoomID = a.ID
	c.ClassRoomID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	classRoomTag, err := a.ClassRoomTags().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range classRoomTag {
		if v.ClassRoomID == b.ClassRoomID {
			bFound = true
		}
		if v.ClassRoomID == c.ClassRoomID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ClassRoomSlice{&a}
	if err = a.L.LoadClassRoomTags(ctx, tx, false, (*[]*ClassRoom)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ClassRoomTags); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ClassRoomTags = nil
	if err = a.L.LoadClassRoomTags(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ClassRoomTags); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", classRoomTag)
	}
}

func testClassRoomToManyUserPositions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ClassRoom
	var b, c UserPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userPositionDBTypes, false, userPositionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userPositionDBTypes, false, userPositionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ClassRoomID = a.ID
	c.ClassRoomID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	userPosition, err := a.UserPositions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range userPosition {
		if v.ClassRoomID == b.ClassRoomID {
			bFound = true
		}
		if v.ClassRoomID == c.ClassRoomID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ClassRoomSlice{&a}
	if err = a.L.LoadUserPositions(ctx, tx, false, (*[]*ClassRoom)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserPositions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserPositions = nil
	if err = a.L.LoadUserPositions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserPositions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", userPosition)
	}
}

func testClassRoomToManyAddOpBeacons(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ClassRoom
	var b, c, d, e Beacon

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classRoomDBTypes, false, strmangle.SetComplement(classRoomPrimaryKeyColumns, classRoomColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Beacon{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, beaconDBTypes, false, strmangle.SetComplement(beaconPrimaryKeyColumns, beaconColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Beacon{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBeacons(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ClassRoomID {
			t.Error("foreign key was wrong value", a.ID, first.ClassRoomID)
		}
		if a.ID != second.ClassRoomID {
			t.Error("foreign key was wrong value", a.ID, second.ClassRoomID)
		}

		if first.R.ClassRoom != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ClassRoom != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Beacons[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Beacons[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Beacons().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testClassRoomToManyAddOpClassRoomTags(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ClassRoom
	var b, c, d, e ClassRoomTag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classRoomDBTypes, false, strmangle.SetComplement(classRoomPrimaryKeyColumns, classRoomColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ClassRoomTag{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, classRoomTagDBTypes, false, strmangle.SetComplement(classRoomTagPrimaryKeyColumns, classRoomTagColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ClassRoomTag{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddClassRoomTags(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ClassRoomID {
			t.Error("foreign key was wrong value", a.ID, first.ClassRoomID)
		}
		if a.ID != second.ClassRoomID {
			t.Error("foreign key was wrong value", a.ID, second.ClassRoomID)
		}

		if first.R.ClassRoom != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ClassRoom != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ClassRoomTags[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ClassRoomTags[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ClassRoomTags().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testClassRoomToManyAddOpUserPositions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ClassRoom
	var b, c, d, e UserPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classRoomDBTypes, false, strmangle.SetComplement(classRoomPrimaryKeyColumns, classRoomColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserPosition{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userPositionDBTypes, false, strmangle.SetComplement(userPositionPrimaryKeyColumns, userPositionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*UserPosition{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserPositions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ClassRoomID {
			t.Error("foreign key was wrong value", a.ID, first.ClassRoomID)
		}
		if a.ID != second.ClassRoomID {
			t.Error("foreign key was wrong value", a.ID, second.ClassRoomID)
		}

		if first.R.ClassRoom != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ClassRoom != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserPositions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserPositions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserPositions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testClassRoomToOneBuildingUsingBuilding(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ClassRoom
	var foreign Building

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, buildingDBTypes, false, buildingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Building struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.BuildingID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Building().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ClassRoomSlice{&local}
	if err = local.L.LoadBuilding(ctx, tx, false, (*[]*ClassRoom)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Building == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Building = nil
	if err = local.L.LoadBuilding(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Building == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testClassRoomToOneSetOpBuildingUsingBuilding(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ClassRoom
	var b, c Building

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classRoomDBTypes, false, strmangle.SetComplement(classRoomPrimaryKeyColumns, classRoomColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, buildingDBTypes, false, strmangle.SetComplement(buildingPrimaryKeyColumns, buildingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, buildingDBTypes, false, strmangle.SetComplement(buildingPrimaryKeyColumns, buildingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Building{&b, &c} {
		err = a.SetBuilding(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Building != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ClassRooms[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.BuildingID, x.ID) {
			t.Error("foreign key was wrong value", a.BuildingID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BuildingID))
		reflect.Indirect(reflect.ValueOf(&a.BuildingID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.BuildingID, x.ID) {
			t.Error("foreign key was wrong value", a.BuildingID, x.ID)
		}
	}
}

func testClassRoomToOneRemoveOpBuildingUsingBuilding(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ClassRoom
	var b Building

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, classRoomDBTypes, false, strmangle.SetComplement(classRoomPrimaryKeyColumns, classRoomColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, buildingDBTypes, false, strmangle.SetComplement(buildingPrimaryKeyColumns, buildingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetBuilding(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveBuilding(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Building().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Building != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.BuildingID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.ClassRooms) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testClassRoomsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClassRoom{}
	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testClassRoomsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClassRoom{}
	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClassRoomSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testClassRoomsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClassRoom{}
	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClassRooms().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	classRoomDBTypes = map[string]string{`BuildingID`: `bigint`, `CreatedAt`: `timestamp without time zone`, `Floor`: `integer`, `ID`: `bigint`, `Latitude`: `double precision`, `LocalX`: `double precision`, `LocalY`: `double precision`, `Longitude`: `double precision`, `Name`: `character varying`, `UpdatedAt`: `timestamp without time zone`}
	_                = bytes.MinRead
)

func testClassRoomsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(classRoomPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(classRoomColumns) == len(classRoomPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClassRoom{}
	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClassRooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testClassRoomsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(classRoomColumns) == len(classRoomPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClassRoom{}
	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClassRooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, classRoomDBTypes, true, classRoomPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(classRoomColumns, classRoomPrimaryKeyColumns) {
		fields = classRoomColumns
	} else {
		fields = strmangle.SetComplement(
			classRoomColumns,
			classRoomPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ClassRoomSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testClassRoomsUpsert(t *testing.T) {
	t.Parallel()

	if len(classRoomColumns) == len(classRoomPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ClassRoom{}
	if err = randomize.Struct(seed, &o, classRoomDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClassRoom: %s", err)
	}

	count, err := ClassRooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, classRoomDBTypes, false, classRoomPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClassRoom: %s", err)
	}

	count, err = ClassRooms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
