// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/goradd/orm/_test/gen/orm/goradd_unit/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(runTests(m))
}

func runTests(m *testing.M) int {
	setup(m)
	defer teardown()
	return m.Run()
}

func setup(m *testing.M) {
	fmt.Println("Setting up tests...")

	//fmt.Println(test.NewSeed())

	// uncomment the next line to re-use a previously created seed to recreate a test
	test.UseSeed(1000)

	InitDB()
}

func teardown() {
	// Cleanup logic here
	fmt.Println("Cleaning up after tests...")
}

// TestDbJson will export the entire database as JSON into a memory buffer, clear the database, then
// import the entire database from the buffer. It will then do some sanity checks.
func TestDbJson(t *testing.T) {
	return
	ctx := db.NewContext(nil)

	// get single comparison objects and data sizes
	// database must be pre-populated for test

	v_DoubleIndex, _ := QueryDoubleIndices(ctx).OrderBy(node.DoubleIndex().PrimaryKey()).Get()            // gets first record
	v_Leaf, _ := QueryLeafs(ctx).OrderBy(node.Leaf().PrimaryKey()).Get()                                  // gets first record
	v_LeafLock, _ := QueryLeafLocks(ctx).OrderBy(node.LeafLock().PrimaryKey()).Get()                      // gets first record
	v_MultiParent, _ := QueryMultiParents(ctx).OrderBy(node.MultiParent().PrimaryKey()).Get()             // gets first record
	v_Root, _ := QueryRoots(ctx).OrderBy(node.Root().PrimaryKey()).Get()                                  // gets first record
	v_TypeTest, _ := QueryTypeTests(ctx).OrderBy(node.TypeTest().PrimaryKey()).Get()                      // gets first record
	v_UnsupportedType, _ := QueryUnsupportedTypes(ctx).OrderBy(node.UnsupportedType().PrimaryKey()).Get() // gets first record
	v_DoubleIndexCount, _ := CountDoubleIndices(ctx)
	v_LeafCount, _ := CountLeafs(ctx)
	v_LeafLockCount, _ := CountLeafLocks(ctx)
	v_MultiParentCount, _ := CountMultiParents(ctx)
	v_RootCount, _ := CountRoots(ctx)
	v_TypeTestCount, _ := CountTypeTests(ctx)
	v_UnsupportedTypeCount, _ := CountUnsupportedTypes(ctx)

	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	assert.NoError(t, JsonEncodeAll(ctx, w))

	ClearAll(ctx)
	assert.Equal(t, 0, func() int { i, _ := CountDoubleIndices(ctx); return i }())
	assert.Equal(t, 0, func() int { i, _ := CountLeafs(ctx); return i }())
	assert.Equal(t, 0, func() int { i, _ := CountLeafLocks(ctx); return i }())
	assert.Equal(t, 0, func() int { i, _ := CountMultiParents(ctx); return i }())
	assert.Equal(t, 0, func() int { i, _ := CountRoots(ctx); return i }())
	assert.Equal(t, 0, func() int { i, _ := CountTypeTests(ctx); return i }())
	assert.Equal(t, 0, func() int { i, _ := CountUnsupportedTypes(ctx); return i }())

	r := bufio.NewReader(&b)
	assert.NoError(t, JsonDecodeAll(ctx, r))

	if v_DoubleIndex != nil {
		obj, _ := QueryDoubleIndices(ctx).OrderBy(node.DoubleIndex().PrimaryKey()).Get()
		assertEqualFieldsDoubleIndex(t, v_DoubleIndex, obj)
	}
	if v_Leaf != nil {
		obj, _ := QueryLeafs(ctx).OrderBy(node.Leaf().PrimaryKey()).Get()
		assertEqualFieldsLeaf(t, v_Leaf, obj)
	}
	if v_LeafLock != nil {
		obj, _ := QueryLeafLocks(ctx).OrderBy(node.LeafLock().PrimaryKey()).Get()
		assertEqualFieldsLeafLock(t, v_LeafLock, obj)
	}
	if v_MultiParent != nil {
		obj, _ := QueryMultiParents(ctx).OrderBy(node.MultiParent().PrimaryKey()).Get()
		assertEqualFieldsMultiParent(t, v_MultiParent, obj)
	}
	if v_Root != nil {
		obj, _ := QueryRoots(ctx).OrderBy(node.Root().PrimaryKey()).Get()
		assertEqualFieldsRoot(t, v_Root, obj)
	}
	if v_TypeTest != nil {
		obj, _ := QueryTypeTests(ctx).OrderBy(node.TypeTest().PrimaryKey()).Get()
		assertEqualFieldsTypeTest(t, v_TypeTest, obj)
	}
	if v_UnsupportedType != nil {
		obj, _ := QueryUnsupportedTypes(ctx).OrderBy(node.UnsupportedType().PrimaryKey()).Get()
		assertEqualFieldsUnsupportedType(t, v_UnsupportedType, obj)
	}
	assert.Equal(t, v_DoubleIndexCount, func() int { i, _ := CountDoubleIndices(ctx); return i }())
	assert.Equal(t, v_LeafCount, func() int { i, _ := CountLeafs(ctx); return i }())
	assert.Equal(t, v_LeafLockCount, func() int { i, _ := CountLeafLocks(ctx); return i }())
	assert.Equal(t, v_MultiParentCount, func() int { i, _ := CountMultiParents(ctx); return i }())
	assert.Equal(t, v_RootCount, func() int { i, _ := CountRoots(ctx); return i }())
	assert.Equal(t, v_TypeTestCount, func() int { i, _ := CountTypeTests(ctx); return i }())
	assert.Equal(t, v_UnsupportedTypeCount, func() int { i, _ := CountUnsupportedTypes(ctx); return i }())
}
