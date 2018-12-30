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
	t.Run("Beacons", testBeacons)
	t.Run("ClassRoomTags", testClassRoomTags)
	t.Run("ClassRooms", testClassRooms)
	t.Run("Sessions", testSessions)
	t.Run("Tags", testTags)
	t.Run("UserPositions", testUserPositions)
	t.Run("UserTags", testUserTags)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Beacons", testBeaconsDelete)
	t.Run("ClassRoomTags", testClassRoomTagsDelete)
	t.Run("ClassRooms", testClassRoomsDelete)
	t.Run("Sessions", testSessionsDelete)
	t.Run("Tags", testTagsDelete)
	t.Run("UserPositions", testUserPositionsDelete)
	t.Run("UserTags", testUserTagsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Beacons", testBeaconsQueryDeleteAll)
	t.Run("ClassRoomTags", testClassRoomTagsQueryDeleteAll)
	t.Run("ClassRooms", testClassRoomsQueryDeleteAll)
	t.Run("Sessions", testSessionsQueryDeleteAll)
	t.Run("Tags", testTagsQueryDeleteAll)
	t.Run("UserPositions", testUserPositionsQueryDeleteAll)
	t.Run("UserTags", testUserTagsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Beacons", testBeaconsSliceDeleteAll)
	t.Run("ClassRoomTags", testClassRoomTagsSliceDeleteAll)
	t.Run("ClassRooms", testClassRoomsSliceDeleteAll)
	t.Run("Sessions", testSessionsSliceDeleteAll)
	t.Run("Tags", testTagsSliceDeleteAll)
	t.Run("UserPositions", testUserPositionsSliceDeleteAll)
	t.Run("UserTags", testUserTagsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Beacons", testBeaconsExists)
	t.Run("ClassRoomTags", testClassRoomTagsExists)
	t.Run("ClassRooms", testClassRoomsExists)
	t.Run("Sessions", testSessionsExists)
	t.Run("Tags", testTagsExists)
	t.Run("UserPositions", testUserPositionsExists)
	t.Run("UserTags", testUserTagsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Beacons", testBeaconsFind)
	t.Run("ClassRoomTags", testClassRoomTagsFind)
	t.Run("ClassRooms", testClassRoomsFind)
	t.Run("Sessions", testSessionsFind)
	t.Run("Tags", testTagsFind)
	t.Run("UserPositions", testUserPositionsFind)
	t.Run("UserTags", testUserTagsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Beacons", testBeaconsBind)
	t.Run("ClassRoomTags", testClassRoomTagsBind)
	t.Run("ClassRooms", testClassRoomsBind)
	t.Run("Sessions", testSessionsBind)
	t.Run("Tags", testTagsBind)
	t.Run("UserPositions", testUserPositionsBind)
	t.Run("UserTags", testUserTagsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Beacons", testBeaconsOne)
	t.Run("ClassRoomTags", testClassRoomTagsOne)
	t.Run("ClassRooms", testClassRoomsOne)
	t.Run("Sessions", testSessionsOne)
	t.Run("Tags", testTagsOne)
	t.Run("UserPositions", testUserPositionsOne)
	t.Run("UserTags", testUserTagsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Beacons", testBeaconsAll)
	t.Run("ClassRoomTags", testClassRoomTagsAll)
	t.Run("ClassRooms", testClassRoomsAll)
	t.Run("Sessions", testSessionsAll)
	t.Run("Tags", testTagsAll)
	t.Run("UserPositions", testUserPositionsAll)
	t.Run("UserTags", testUserTagsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Beacons", testBeaconsCount)
	t.Run("ClassRoomTags", testClassRoomTagsCount)
	t.Run("ClassRooms", testClassRoomsCount)
	t.Run("Sessions", testSessionsCount)
	t.Run("Tags", testTagsCount)
	t.Run("UserPositions", testUserPositionsCount)
	t.Run("UserTags", testUserTagsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Beacons", testBeaconsHooks)
	t.Run("ClassRoomTags", testClassRoomTagsHooks)
	t.Run("ClassRooms", testClassRoomsHooks)
	t.Run("Sessions", testSessionsHooks)
	t.Run("Tags", testTagsHooks)
	t.Run("UserPositions", testUserPositionsHooks)
	t.Run("UserTags", testUserTagsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Beacons", testBeaconsInsert)
	t.Run("Beacons", testBeaconsInsertWhitelist)
	t.Run("ClassRoomTags", testClassRoomTagsInsert)
	t.Run("ClassRoomTags", testClassRoomTagsInsertWhitelist)
	t.Run("ClassRooms", testClassRoomsInsert)
	t.Run("ClassRooms", testClassRoomsInsertWhitelist)
	t.Run("Sessions", testSessionsInsert)
	t.Run("Sessions", testSessionsInsertWhitelist)
	t.Run("Tags", testTagsInsert)
	t.Run("Tags", testTagsInsertWhitelist)
	t.Run("UserPositions", testUserPositionsInsert)
	t.Run("UserPositions", testUserPositionsInsertWhitelist)
	t.Run("UserTags", testUserTagsInsert)
	t.Run("UserTags", testUserTagsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("BeaconToClassRoomUsingClassRoom", testBeaconToOneClassRoomUsingClassRoom)
	t.Run("ClassRoomTagToClassRoomUsingClassRoom", testClassRoomTagToOneClassRoomUsingClassRoom)
	t.Run("ClassRoomTagToTagUsingTag", testClassRoomTagToOneTagUsingTag)
	t.Run("SessionToUserUsingUser", testSessionToOneUserUsingUser)
	t.Run("UserPositionToUserUsingUser", testUserPositionToOneUserUsingUser)
	t.Run("UserPositionToClassRoomUsingClassRoom", testUserPositionToOneClassRoomUsingClassRoom)
	t.Run("UserTagToUserUsingUser", testUserTagToOneUserUsingUser)
	t.Run("UserTagToTagUsingTag", testUserTagToOneTagUsingTag)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ClassRoomToBeacons", testClassRoomToManyBeacons)
	t.Run("ClassRoomToClassRoomTags", testClassRoomToManyClassRoomTags)
	t.Run("ClassRoomToUserPositions", testClassRoomToManyUserPositions)
	t.Run("TagToClassRoomTags", testTagToManyClassRoomTags)
	t.Run("TagToUserTags", testTagToManyUserTags)
	t.Run("UserToSessions", testUserToManySessions)
	t.Run("UserToUserPositions", testUserToManyUserPositions)
	t.Run("UserToUserTags", testUserToManyUserTags)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("BeaconToClassRoomUsingBeacons", testBeaconToOneSetOpClassRoomUsingClassRoom)
	t.Run("ClassRoomTagToClassRoomUsingClassRoomTags", testClassRoomTagToOneSetOpClassRoomUsingClassRoom)
	t.Run("ClassRoomTagToTagUsingClassRoomTags", testClassRoomTagToOneSetOpTagUsingTag)
	t.Run("SessionToUserUsingSessions", testSessionToOneSetOpUserUsingUser)
	t.Run("UserPositionToUserUsingUserPositions", testUserPositionToOneSetOpUserUsingUser)
	t.Run("UserPositionToClassRoomUsingUserPositions", testUserPositionToOneSetOpClassRoomUsingClassRoom)
	t.Run("UserTagToUserUsingUserTags", testUserTagToOneSetOpUserUsingUser)
	t.Run("UserTagToTagUsingUserTags", testUserTagToOneSetOpTagUsingTag)
}

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
func TestToManyAdd(t *testing.T) {
	t.Run("ClassRoomToBeacons", testClassRoomToManyAddOpBeacons)
	t.Run("ClassRoomToClassRoomTags", testClassRoomToManyAddOpClassRoomTags)
	t.Run("ClassRoomToUserPositions", testClassRoomToManyAddOpUserPositions)
	t.Run("TagToClassRoomTags", testTagToManyAddOpClassRoomTags)
	t.Run("TagToUserTags", testTagToManyAddOpUserTags)
	t.Run("UserToSessions", testUserToManyAddOpSessions)
	t.Run("UserToUserPositions", testUserToManyAddOpUserPositions)
	t.Run("UserToUserTags", testUserToManyAddOpUserTags)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Beacons", testBeaconsReload)
	t.Run("ClassRoomTags", testClassRoomTagsReload)
	t.Run("ClassRooms", testClassRoomsReload)
	t.Run("Sessions", testSessionsReload)
	t.Run("Tags", testTagsReload)
	t.Run("UserPositions", testUserPositionsReload)
	t.Run("UserTags", testUserTagsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Beacons", testBeaconsReloadAll)
	t.Run("ClassRoomTags", testClassRoomTagsReloadAll)
	t.Run("ClassRooms", testClassRoomsReloadAll)
	t.Run("Sessions", testSessionsReloadAll)
	t.Run("Tags", testTagsReloadAll)
	t.Run("UserPositions", testUserPositionsReloadAll)
	t.Run("UserTags", testUserTagsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Beacons", testBeaconsSelect)
	t.Run("ClassRoomTags", testClassRoomTagsSelect)
	t.Run("ClassRooms", testClassRoomsSelect)
	t.Run("Sessions", testSessionsSelect)
	t.Run("Tags", testTagsSelect)
	t.Run("UserPositions", testUserPositionsSelect)
	t.Run("UserTags", testUserTagsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Beacons", testBeaconsUpdate)
	t.Run("ClassRoomTags", testClassRoomTagsUpdate)
	t.Run("ClassRooms", testClassRoomsUpdate)
	t.Run("Sessions", testSessionsUpdate)
	t.Run("Tags", testTagsUpdate)
	t.Run("UserPositions", testUserPositionsUpdate)
	t.Run("UserTags", testUserTagsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Beacons", testBeaconsSliceUpdateAll)
	t.Run("ClassRoomTags", testClassRoomTagsSliceUpdateAll)
	t.Run("ClassRooms", testClassRoomsSliceUpdateAll)
	t.Run("Sessions", testSessionsSliceUpdateAll)
	t.Run("Tags", testTagsSliceUpdateAll)
	t.Run("UserPositions", testUserPositionsSliceUpdateAll)
	t.Run("UserTags", testUserTagsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
