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

func testUserPositions(t *testing.T) {
	t.Parallel()

	query := UserPositions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserPositionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPosition{}
	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
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

	count, err := UserPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserPositionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPosition{}
	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserPositions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserPositionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPosition{}
	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserPositionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserPositionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPosition{}
	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserPositionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if UserPosition exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserPositionExists to return true, but got false.")
	}
}

func testUserPositionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPosition{}
	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userPositionFound, err := FindUserPosition(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if userPositionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserPositionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPosition{}
	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserPositions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserPositionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPosition{}
	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserPositions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserPositionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userPositionOne := &UserPosition{}
	userPositionTwo := &UserPosition{}
	if err = randomize.Struct(seed, userPositionOne, userPositionDBTypes, false, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}
	if err = randomize.Struct(seed, userPositionTwo, userPositionDBTypes, false, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userPositionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userPositionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserPositions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserPositionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userPositionOne := &UserPosition{}
	userPositionTwo := &UserPosition{}
	if err = randomize.Struct(seed, userPositionOne, userPositionDBTypes, false, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}
	if err = randomize.Struct(seed, userPositionTwo, userPositionDBTypes, false, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userPositionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userPositionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userPositionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserPosition) error {
	*o = UserPosition{}
	return nil
}

func userPositionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserPosition) error {
	*o = UserPosition{}
	return nil
}

func userPositionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserPosition) error {
	*o = UserPosition{}
	return nil
}

func userPositionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserPosition) error {
	*o = UserPosition{}
	return nil
}

func userPositionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserPosition) error {
	*o = UserPosition{}
	return nil
}

func userPositionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserPosition) error {
	*o = UserPosition{}
	return nil
}

func userPositionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserPosition) error {
	*o = UserPosition{}
	return nil
}

func userPositionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserPosition) error {
	*o = UserPosition{}
	return nil
}

func userPositionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserPosition) error {
	*o = UserPosition{}
	return nil
}

func testUserPositionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserPosition{}
	o := &UserPosition{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userPositionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserPosition object: %s", err)
	}

	AddUserPositionHook(boil.BeforeInsertHook, userPositionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userPositionBeforeInsertHooks = []UserPositionHook{}

	AddUserPositionHook(boil.AfterInsertHook, userPositionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userPositionAfterInsertHooks = []UserPositionHook{}

	AddUserPositionHook(boil.AfterSelectHook, userPositionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userPositionAfterSelectHooks = []UserPositionHook{}

	AddUserPositionHook(boil.BeforeUpdateHook, userPositionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userPositionBeforeUpdateHooks = []UserPositionHook{}

	AddUserPositionHook(boil.AfterUpdateHook, userPositionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userPositionAfterUpdateHooks = []UserPositionHook{}

	AddUserPositionHook(boil.BeforeDeleteHook, userPositionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userPositionBeforeDeleteHooks = []UserPositionHook{}

	AddUserPositionHook(boil.AfterDeleteHook, userPositionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userPositionAfterDeleteHooks = []UserPositionHook{}

	AddUserPositionHook(boil.BeforeUpsertHook, userPositionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userPositionBeforeUpsertHooks = []UserPositionHook{}

	AddUserPositionHook(boil.AfterUpsertHook, userPositionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userPositionAfterUpsertHooks = []UserPositionHook{}
}

func testUserPositionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPosition{}
	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserPositionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPosition{}
	if err = randomize.Struct(seed, o, userPositionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userPositionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserPositionToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserPosition
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userPositionDBTypes, false, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UserPositionSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*UserPosition)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUserPositionToOneClassRoomUsingClassRoom(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserPosition
	var foreign ClassRoom

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userPositionDBTypes, false, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, classRoomDBTypes, false, classRoomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClassRoom struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ClassRoomID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ClassRoom().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UserPositionSlice{&local}
	if err = local.L.LoadClassRoom(ctx, tx, false, (*[]*UserPosition)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ClassRoom == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ClassRoom = nil
	if err = local.L.LoadClassRoom(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ClassRoom == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUserPositionToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserPosition
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userPositionDBTypes, false, strmangle.SetComplement(userPositionPrimaryKeyColumns, userPositionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserPositions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}
func testUserPositionToOneSetOpClassRoomUsingClassRoom(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserPosition
	var b, c ClassRoom

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userPositionDBTypes, false, strmangle.SetComplement(userPositionPrimaryKeyColumns, userPositionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, classRoomDBTypes, false, strmangle.SetComplement(classRoomPrimaryKeyColumns, classRoomColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, classRoomDBTypes, false, strmangle.SetComplement(classRoomPrimaryKeyColumns, classRoomColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ClassRoom{&b, &c} {
		err = a.SetClassRoom(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ClassRoom != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserPositions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ClassRoomID != x.ID {
			t.Error("foreign key was wrong value", a.ClassRoomID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ClassRoomID))
		reflect.Indirect(reflect.ValueOf(&a.ClassRoomID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ClassRoomID != x.ID {
			t.Error("foreign key was wrong value", a.ClassRoomID, x.ID)
		}
	}
}

func testUserPositionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPosition{}
	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
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

func testUserPositionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPosition{}
	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserPositionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserPositionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserPosition{}
	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserPositions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userPositionDBTypes = map[string]string{`ClassRoomID`: `bigint`, `CreatedAt`: `timestamp without time zone`, `ID`: `bigint`, `UpdatedAt`: `timestamp without time zone`, `UserID`: `bigint`}
	_                   = bytes.MinRead
)

func testUserPositionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userPositionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userPositionColumns) == len(userPositionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserPosition{}
	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserPositionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userPositionColumns) == len(userPositionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserPosition{}
	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userPositionDBTypes, true, userPositionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userPositionColumns, userPositionPrimaryKeyColumns) {
		fields = userPositionColumns
	} else {
		fields = strmangle.SetComplement(
			userPositionColumns,
			userPositionPrimaryKeyColumns,
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

	slice := UserPositionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserPositionsUpsert(t *testing.T) {
	t.Parallel()

	if len(userPositionColumns) == len(userPositionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserPosition{}
	if err = randomize.Struct(seed, &o, userPositionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserPosition: %s", err)
	}

	count, err := UserPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userPositionDBTypes, false, userPositionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserPosition struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserPosition: %s", err)
	}

	count, err = UserPositions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}