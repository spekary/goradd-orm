// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"context"
	"strconv"
	"testing"

	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createMinimalSampleLogin creates an unsaved minimal version of a Login object
// for testing.
func createMinimalSampleLogin() *Login {
	obj := NewLogin()
	updateMinimalSampleLogin(obj)

	return obj
}

// updateMinimalSampleLogin sets the values of a minimal sample to new, random values.
func updateMinimalSampleLogin(obj *Login) {

	obj.SetUsername(test.RandomValue[string](20))

	obj.SetPassword(test.RandomValue[string](20))

	obj.SetIsEnabled(test.RandomValue[bool](0))

}

// createMaximalSampleLogin creates an unsaved version of a Login object
// for testing that includes references to minimal objects.
func createMaximalSampleLogin() *Login {
	obj := NewLogin()
	updateMaximalSampleLogin(obj)
	return obj
}

// updateMaximalSampleLogin sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleLogin(obj *Login) {
	updateMinimalSampleLogin(obj)

	obj.SetPerson(createMinimalSamplePerson())

}

// deleteSampleLogin deletes an object created and saved by one of the sample creator functions.
func deleteSampleLogin(ctx context.Context, obj *Login) {
	if obj == nil {
		return
	}

	obj.Delete(ctx)

	deleteSamplePerson(ctx, obj.Person())

}

// assertEqualFieldsLogin compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsLogin(t *testing.T, obj1, obj2 *Login) {
	assert.EqualValues(t, obj1.ID(), obj2.ID())

	assert.EqualValues(t, obj1.PersonID(), obj2.PersonID())

	assert.EqualValues(t, obj1.Username(), obj2.Username())

	assert.EqualValues(t, obj1.Password(), obj2.Password())

	assert.EqualValues(t, obj1.IsEnabled(), obj2.IsEnabled())

}

func TestLogin_SetPersonID(t *testing.T) {

	obj := NewLogin()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](0)
	obj.SetPersonID(val)
	assert.Equal(t, val, obj.PersonID())
	assert.False(t, obj.PersonIDIsNull())

	// Test NULL
	obj.SetPersonIDToNull()
	assert.EqualValues(t, "", obj.PersonID())
	assert.True(t, obj.PersonIDIsNull())

	// test default
	obj.SetPersonID("")
	assert.EqualValues(t, "", obj.PersonID(), "set default")

}
func TestLogin_SetUsername(t *testing.T) {

	obj := NewLogin()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](20)
	obj.SetUsername(val)
	assert.Equal(t, val, obj.Username())

	// test default
	obj.SetUsername("")
	assert.EqualValues(t, "", obj.Username(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](21)
	assert.Panics(t, func() {
		obj.SetUsername(val)
	})
}
func TestLogin_SetPassword(t *testing.T) {

	obj := NewLogin()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](20)
	obj.SetPassword(val)
	assert.Equal(t, val, obj.Password())
	assert.False(t, obj.PasswordIsNull())

	// Test NULL
	obj.SetPasswordToNull()
	assert.EqualValues(t, "", obj.Password())
	assert.True(t, obj.PasswordIsNull())

	// test default
	obj.SetPassword("")
	assert.EqualValues(t, "", obj.Password(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](21)
	assert.Panics(t, func() {
		obj.SetPassword(val)
	})
}
func TestLogin_SetIsEnabled(t *testing.T) {

	obj := NewLogin()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[bool](0)
	obj.SetIsEnabled(val)
	assert.Equal(t, val, obj.IsEnabled())

	// test default
	obj.SetIsEnabled(true)
	assert.EqualValues(t, true, obj.IsEnabled(), "set default")

}

func TestLogin_Copy(t *testing.T) {
	obj := createMinimalSampleLogin()

	obj2 := obj.Copy()

	assert.Equal(t, obj.PersonID(), obj2.PersonID())
	assert.Equal(t, obj.Username(), obj2.Username())
	assert.Equal(t, obj.Password(), obj2.Password())
	assert.Equal(t, obj.IsEnabled(), obj2.IsEnabled())

}

func TestLogin_BasicInsert(t *testing.T) {
	obj := createMinimalSampleLogin()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLogin(ctx, obj)

	// Test retrieval
	obj2 := LoadLogin(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.IDIsValid())

	assert.True(t, obj2.UsernameIsValid())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.usernameIsDirty)
	obj2.SetUsername(obj2.Username())
	assert.False(t, obj2.usernameIsDirty)

	assert.True(t, obj2.PasswordIsValid())
	assert.False(t, obj2.PasswordIsNull())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.passwordIsDirty)
	obj2.SetPassword(obj2.Password())
	assert.False(t, obj2.passwordIsDirty)

	assert.True(t, obj2.IsEnabledIsValid())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.isEnabledIsDirty)
	obj2.SetIsEnabled(obj2.IsEnabled())
	assert.False(t, obj2.isEnabledIsDirty)

}

func TestLogin_InsertPanics(t *testing.T) {
	obj := createMinimalSampleLogin()
	ctx := db.NewContext(nil)

	obj.usernameIsValid = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.usernameIsValid = true

}

func TestLogin_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleLogin()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLogin(ctx, obj)
	updateMinimalSampleLogin(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2 := LoadLogin(ctx, obj.PrimaryKey())

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.PersonID(), obj.PersonID(), "PersonID did not update")
	assert.Equal(t, obj2.Username(), obj.Username(), "Username did not update")
	assert.Equal(t, obj2.Password(), obj.Password(), "Password did not update")
	assert.Equal(t, obj2.IsEnabled(), obj.IsEnabled(), "IsEnabled did not update")
}

func TestLogin_ReferenceLoad(t *testing.T) {
	obj := createMaximalSampleLogin()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSampleLogin(ctx, obj)

	// Test that referenced objects were saved and assigned ids
	assert.NotNil(t, obj.Person())
	assert.NotEqual(t, '-', obj.Person().PrimaryKey()[0])

	// Test lazy loading
	obj2 := LoadLogin(ctx, obj.PrimaryKey())
	objPkOnly := LoadLogin(ctx, obj.PrimaryKey(), node.Login().PrimaryKey())
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	assert.Nil(t, obj2.Person(), "Person is not loaded initially")
	v_Person := obj2.LoadPerson(ctx)
	assert.NotNil(t, v_Person)
	assert.Equal(t, v_Person.PrimaryKey(), obj2.Person().PrimaryKey())
	assert.Equal(t, obj.Person().PrimaryKey(), obj2.Person().PrimaryKey())
	assert.True(t, obj2.PersonIDIsValid())

	assert.False(t, objPkOnly.PersonIDIsValid())
	assert.Nil(t, objPkOnly.LoadPerson(ctx))

	// test eager loading
	obj3 := LoadLogin(ctx, obj.PrimaryKey(), node.Login().Person())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.Person().PrimaryKey(), obj3.Person().PrimaryKey())

}

func TestLogin_ReferenceUpdateNewObjects(t *testing.T) {
	obj := createMaximalSampleLogin()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSampleLogin(ctx, obj)

	obj2 := LoadLogin(ctx, obj.PrimaryKey())
	updateMaximalSampleLogin(obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleLogin(ctx, obj2)

	obj3 := LoadLogin(ctx, obj2.PrimaryKey(), node.Login().Person())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.Person().PrimaryKey(), obj3.Person().PrimaryKey())

}

func TestLogin_ReferenceUpdateOldObjects(t *testing.T) {
	obj := createMaximalSampleLogin()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLogin(ctx, obj)

	updateMinimalSamplePerson(obj.Person())

	assert.NoError(t, obj.Save(ctx))

	obj2 := LoadLogin(ctx, obj.PrimaryKey(),
		node.Login().Person(),
	)
	_ = obj2 // avoid error if there are no references

	assertEqualFieldsPerson(t, obj2.Person(), obj.Person())

}
func TestLogin_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewLogin()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)
}

func TestLogin_Getters(t *testing.T) {
	obj := createMinimalSampleLogin()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSampleLogin(ctx, obj)

	assert.True(t, HasLogin(ctx, obj.PrimaryKey()))

	obj2 := LoadLogin(ctx, obj.PrimaryKey(), node.Login().PrimaryKey())

	assert.Panics(t, func() { obj2.PersonID() })
	assert.Panics(t, func() { obj2.Username() })
	assert.Panics(t, func() { obj2.Password() })
	assert.Panics(t, func() { obj2.IsEnabled() })
}

func TestLogin_QueryLoad(t *testing.T) {
	obj := createMinimalSampleLogin()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLogin(ctx, obj)

	objs := QueryLogins(ctx).
		Where(op.Equal(node.Login().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.Login().PrimaryKey()). // exercise order by
		Limit(1, 0).                        // exercise limit
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
}
func TestLogin_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleLogin()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLogin(ctx, obj)

	objs := QueryLogins(ctx).
		Where(op.Equal(node.Login().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestLogin_QueryCursor(t *testing.T) {
	obj := createMinimalSampleLogin()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLogin(ctx, obj)

	cursor := QueryLogins(ctx).
		Where(op.Equal(node.Login().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()

	obj2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.Nil(t, cursor.Next())

	// test empty cursor result
	cursor = QueryLogins(ctx).
		Where(op.Equal(1, 0)).
		LoadCursor()
	assert.Nil(t, cursor.Next())

}
func TestLogin_Count(t *testing.T) {
	obj := createMaximalSampleLogin()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLogin(ctx, obj)

	assert.Less(t, 0, CountLoginsByID(ctx, obj.ID()))
	assert.Less(t, 0, CountLoginsByPersonID(ctx, obj.PersonID()))
	assert.Less(t, 0, CountLoginsByUsername(ctx, obj.Username()))
	assert.Less(t, 0, CountLoginsByPassword(ctx, obj.Password()))
	assert.Less(t, 0, CountLoginsByIsEnabled(ctx, obj.IsEnabled()))
}
