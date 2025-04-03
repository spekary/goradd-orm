// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createMinimalSampleAddress creates an unsaved minimal version of a Address object
// for testing.
func createMinimalSampleAddress() *Address {
	obj := NewAddress()
	updateMinimalSampleAddress(obj)

	// A required forward reference will need to be fulfilled just to save the minimal version of this object
	// If the database is configured so that the referenced object points back here, either directly or through multiple
	// forward references, it possible this could create an endless loop. Such a structure could not be saved anyways.
	obj.SetPerson(createMinimalSamplePerson())

	return obj
}

// updateMinimalSampleAddress sets the values of a minimal sample to new, random values.
func updateMinimalSampleAddress(obj *Address) {

	obj.SetStreet(test.RandomValue[string](100))

	obj.SetCity(test.RandomValue[string](100))

}

// createMaximalSampleAddress creates an unsaved version of a Address object
// for testing that includes references to minimal objects.
func createMaximalSampleAddress(ctx context.Context) *Address {
	obj := NewAddress()
	updateMaximalSampleAddress(ctx, obj)
	return obj
}

// updateMaximalSampleAddress sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleAddress(ctx context.Context, obj *Address) {
	updateMinimalSampleAddress(obj)
	obj.SetPerson(createMinimalSamplePerson())

}

// deleteSampleAddress deletes an object created and saved by one of the sample creator functions.
func deleteSampleAddress(ctx context.Context, obj *Address) {
	if obj == nil {
		return
	}

	obj.Delete(ctx)

	deleteSamplePerson(ctx, obj.Person())

}

// assertEqualFieldsAddress compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsAddress(t *testing.T, obj1, obj2 *Address) {
	if obj1.IDIsLoaded() && obj2.IDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.ID(), obj2.ID())
	}
	if obj1.PersonIDIsLoaded() && obj2.PersonIDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.PersonID(), obj2.PersonID())
	}
	if obj1.StreetIsLoaded() && obj2.StreetIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Street(), obj2.Street())
	}
	if obj1.CityIsLoaded() && obj2.CityIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.City(), obj2.City())
	}

}

func TestAddress_SetID(t *testing.T) {

	obj := NewAddress()

	assert.True(t, obj.IsNew())
	val := test.RandomNumberString()
	obj.SetID(val)
	assert.Equal(t, val, obj.ID())

	// test default
	obj.SetID("")
	assert.EqualValues(t, "", obj.ID(), "set default")

}
func TestAddress_SetPersonID(t *testing.T) {

	obj := NewAddress()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](0)
	obj.SetPersonID(val)
	assert.Equal(t, val, obj.PersonID())

	// test default
	obj.SetPersonID("")
	assert.EqualValues(t, "", obj.PersonID(), "set default")

}
func TestAddress_SetStreet(t *testing.T) {

	obj := NewAddress()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](100)
	obj.SetStreet(val)
	assert.Equal(t, val, obj.Street())

	// test default
	obj.SetStreet("")
	assert.EqualValues(t, "", obj.Street(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](101)
	assert.Panics(t, func() {
		obj.SetStreet(val)
	})
}
func TestAddress_SetCity(t *testing.T) {

	obj := NewAddress()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](100)
	obj.SetCity(val)
	assert.Equal(t, val, obj.City())
	assert.False(t, obj.CityIsNull())

	// Test NULL
	obj.SetCityToNull()
	assert.EqualValues(t, "BOB", obj.City())
	assert.True(t, obj.CityIsNull())

	// test default
	obj.SetCity("BOB")
	assert.EqualValues(t, "BOB", obj.City(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](101)
	assert.Panics(t, func() {
		obj.SetCity(val)
	})
}

func TestAddress_Copy(t *testing.T) {
	obj := createMinimalSampleAddress()

	obj2 := obj.Copy()

	assert.Equal(t, obj.PersonID(), obj2.PersonID())
	assert.Equal(t, obj.Street(), obj2.Street())
	assert.Equal(t, obj.City(), obj2.City())

}

func TestAddress_BasicInsert(t *testing.T) {
	obj := createMinimalSampleAddress()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleAddress(ctx, obj)

	// Test retrieval
	obj2 := LoadAddress(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.IDIsLoaded())
	assert.Panics(t, func() {
		obj2.SetID(obj2.ID())
	})

	assert.True(t, obj2.StreetIsLoaded())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.streetIsDirty)
	obj2.SetStreet(obj2.Street())
	assert.False(t, obj2.streetIsDirty)

	assert.True(t, obj2.CityIsLoaded())
	assert.False(t, obj2.CityIsNull())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.cityIsDirty)
	obj2.SetCity(obj2.City())
	assert.False(t, obj2.cityIsDirty)

}

func TestAddress_InsertPanics(t *testing.T) {
	obj := createMinimalSampleAddress()
	_ = obj
	ctx := db.NewContext(nil)
	_ = ctx

	obj.personIDIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.personIDIsLoaded = true

	obj.streetIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.streetIsLoaded = true

}

func TestAddress_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleAddress()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleAddress(ctx, obj)
	updateMinimalSampleAddress(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2 := LoadAddress(ctx, obj.PrimaryKey())

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.Street(), obj.Street(), "Street did not update")
	assert.Equal(t, obj2.City(), obj.City(), "City did not update")
}

func TestAddress_ReferenceLoad(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleAddress(ctx)
	obj.Save(ctx)
	defer deleteSampleAddress(ctx, obj)

	// Test that referenced objects were saved and assigned ids
	assert.NotNil(t, obj.Person())
	assert.NotEqual(t, '-', obj.Person().PrimaryKey()[0])

	// Test lazy loading
	obj2 := LoadAddress(ctx, obj.PrimaryKey())
	objPkOnly := LoadAddress(ctx, obj.PrimaryKey(), node.Address().PrimaryKey())
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	assert.Nil(t, obj2.Person(), "Person is not loaded initially")
	v_Person := obj2.LoadPerson(ctx)
	assert.NotNil(t, v_Person)
	assert.Equal(t, v_Person.PrimaryKey(), obj2.Person().PrimaryKey())
	assert.Equal(t, obj.Person().PrimaryKey(), obj2.Person().PrimaryKey())
	assert.True(t, obj2.PersonIDIsLoaded())

	assert.False(t, objPkOnly.PersonIDIsLoaded())
	assert.Nil(t, objPkOnly.LoadPerson(ctx))

	assert.Panics(t, func() {
		objPkOnly.SetPerson(nil)
	})

	// test eager loading
	obj3 := LoadAddress(ctx, obj.PrimaryKey(), node.Address().Person())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.Person().PrimaryKey(), obj3.Person().PrimaryKey())

}

func TestAddress_ReferenceUpdateNewObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleAddress(ctx)
	obj.Save(ctx)
	defer deleteSampleAddress(ctx, obj)

	obj2 := LoadAddress(ctx, obj.PrimaryKey())
	updateMaximalSampleAddress(ctx, obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleAddress(ctx, obj2)

	obj3 := LoadAddress(ctx, obj2.PrimaryKey(), node.Address().Person())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.Person().PrimaryKey(), obj3.Person().PrimaryKey())

}

func TestAddress_ReferenceUpdateOldObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleAddress(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleAddress(ctx, obj)

	updateMinimalSamplePerson(obj.Person())

	assert.NoError(t, obj.Save(ctx))

	obj2 := LoadAddress(ctx, obj.PrimaryKey(),
		node.Address().Person(),
	)
	_ = obj2 // avoid error if there are no references

	assertEqualFieldsPerson(t, obj2.Person(), obj.Person())

}
func TestAddress_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewAddress()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)
}

func TestAddress_Getters(t *testing.T) {
	obj := createMinimalSampleAddress()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSampleAddress(ctx, obj)

	assert.True(t, HasAddress(ctx, obj.PrimaryKey()))

	obj2 := LoadAddress(ctx, obj.PrimaryKey(), node.Address().PrimaryKey())

	assert.Equal(t, obj.ID(), obj.Get(node.Address().ID().Identifier))
	assert.Panics(t, func() { obj2.PersonID() })
	assert.Nil(t, obj2.Get(node.Address().PersonID().Identifier))
	assert.Equal(t, obj.Street(), obj.Get(node.Address().Street().Identifier))
	assert.Panics(t, func() { obj2.Street() })
	assert.Nil(t, obj2.Get(node.Address().Street().Identifier))
	assert.Equal(t, obj.City(), obj.Get(node.Address().City().Identifier))
	assert.Panics(t, func() { obj2.City() })
	assert.Nil(t, obj2.Get(node.Address().City().Identifier))
}

func TestAddress_QueryLoad(t *testing.T) {
	obj := createMinimalSampleAddress()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleAddress(ctx, obj)

	objs := QueryAddresses(ctx).
		Where(op.Equal(node.Address().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.Address().PrimaryKey()). // exercise order by
		Limit(1, 0).                          // exercise limit
		Calculation(node.Address(), "IsTrue", op.Equal(1, 1)).
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
	assert.True(t, objs[0].GetAlias("IsTrue").Bool())
}
func TestAddress_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleAddress()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleAddress(ctx, obj)

	objs := QueryAddresses(ctx).
		Where(op.Equal(node.Address().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestAddress_QueryCursor(t *testing.T) {
	obj := createMinimalSampleAddress()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleAddress(ctx, obj)

	cursor := QueryAddresses(ctx).
		Where(op.Equal(node.Address().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()

	obj2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.Nil(t, cursor.Next())

	// test empty cursor result
	cursor = QueryAddresses(ctx).
		Where(op.Equal(1, 0)).
		LoadCursor()
	assert.Nil(t, cursor.Next())

}
func TestAddress_Count(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleAddress(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	// reread in case there are data limitations imposed by the database
	obj2 := LoadAddress(ctx, obj.PrimaryKey())
	defer deleteSampleAddress(ctx, obj)

	assert.Less(t, 0, CountAddresses(ctx))

	assert.Less(t, 0, CountAddressesByID(ctx, obj2.ID()))
	assert.Less(t, 0, CountAddressesByPersonID(ctx, obj2.PersonID()))
	assert.Less(t, 0, CountAddressesByStreet(ctx, obj2.Street()))
	assert.Less(t, 0, CountAddressesByCity(ctx, obj2.City()))

}
func TestAddress_MarshalJSON(t *testing.T) {
	obj := createMinimalSampleAddress()

	b, err := json.Marshal(obj)
	assert.NoError(t, err)

	obj2 := NewAddress()
	err = json.Unmarshal(b, &obj2)
	assert.NoError(t, err)

	assertEqualFieldsAddress(t, obj, obj2)
}

func TestAddress_MarshalBinary(t *testing.T) {
	obj := createMinimalSampleAddress()

	b, err := obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 := NewAddress()
	err = obj2.UnmarshalBinary(b)
	assert.NoError(t, err)

	assertEqualFieldsAddress(t, obj, obj2)
}
