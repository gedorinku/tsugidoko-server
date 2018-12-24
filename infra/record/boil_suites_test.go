// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package record

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Sessions", testSessions)
	t.Run("Tags", testTags)
	t.Run("UserTags", testUserTags)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Sessions", testSessionsDelete)
	t.Run("Tags", testTagsDelete)
	t.Run("UserTags", testUserTagsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Sessions", testSessionsQueryDeleteAll)
	t.Run("Tags", testTagsQueryDeleteAll)
	t.Run("UserTags", testUserTagsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Sessions", testSessionsSliceDeleteAll)
	t.Run("Tags", testTagsSliceDeleteAll)
	t.Run("UserTags", testUserTagsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Sessions", testSessionsExists)
	t.Run("Tags", testTagsExists)
	t.Run("UserTags", testUserTagsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Sessions", testSessionsFind)
	t.Run("Tags", testTagsFind)
	t.Run("UserTags", testUserTagsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Sessions", testSessionsBind)
	t.Run("Tags", testTagsBind)
	t.Run("UserTags", testUserTagsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Sessions", testSessionsOne)
	t.Run("Tags", testTagsOne)
	t.Run("UserTags", testUserTagsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Sessions", testSessionsAll)
	t.Run("Tags", testTagsAll)
	t.Run("UserTags", testUserTagsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Sessions", testSessionsCount)
	t.Run("Tags", testTagsCount)
	t.Run("UserTags", testUserTagsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Sessions", testSessionsHooks)
	t.Run("Tags", testTagsHooks)
	t.Run("UserTags", testUserTagsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Sessions", testSessionsInsert)
	t.Run("Sessions", testSessionsInsertWhitelist)
	t.Run("Tags", testTagsInsert)
	t.Run("Tags", testTagsInsertWhitelist)
	t.Run("UserTags", testUserTagsInsert)
	t.Run("UserTags", testUserTagsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Sessions", testSessionsReload)
	t.Run("Tags", testTagsReload)
	t.Run("UserTags", testUserTagsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Sessions", testSessionsReloadAll)
	t.Run("Tags", testTagsReloadAll)
	t.Run("UserTags", testUserTagsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Sessions", testSessionsSelect)
	t.Run("Tags", testTagsSelect)
	t.Run("UserTags", testUserTagsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Sessions", testSessionsUpdate)
	t.Run("Tags", testTagsUpdate)
	t.Run("UserTags", testUserTagsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Sessions", testSessionsSliceUpdateAll)
	t.Run("Tags", testTagsSliceUpdateAll)
	t.Run("UserTags", testUserTagsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
