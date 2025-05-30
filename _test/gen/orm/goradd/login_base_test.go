// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"bytes"
	"context"
	"encoding/gob"
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
func createMaximalSampleLogin(ctx context.Context) *Login {
	obj := NewLogin()
	updateMaximalSampleLogin(ctx, obj)
	return obj
}

// updateMaximalSampleLogin sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleLogin(ctx context.Context, obj *Login) {
	updateMinimalSampleLogin(obj)
	obj.SetPerson(createMinimalSamplePerson())

}

// deleteSampleLogin deletes an object created and saved by one of the sample creator functions.
func deleteSampleLogin(ctx context.Context, obj *Login) {
	if obj == nil {
		return
	}

	_ = obj.Delete(ctx)

	deleteSamplePerson(ctx, obj.Person())

}

// assertEqualFieldsLogin compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsLogin(t *testing.T, obj1, obj2 *Login) {
	if obj1.IDIsLoaded() && obj2.IDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.ID(), obj2.ID())
	}
	if obj1.PersonIDIsLoaded() && obj2.PersonIDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.PersonID(), obj2.PersonID())
	}
	if obj1.UsernameIsLoaded() && obj2.UsernameIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Username(), obj2.Username())
	}
	if obj1.PasswordIsLoaded() && obj2.PasswordIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Password(), obj2.Password())
	}
	if obj1.IsEnabledIsLoaded() && obj2.IsEnabledIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.IsEnabled(), obj2.IsEnabled())
	}

}

func TestLogin_SetID(t *testing.T) {

	obj := NewLogin()

	assert.True(t, obj.IsNew())
	val := test.RandomNumberString()
	obj.SetID(val)
	assert.Equal(t, val, obj.ID())

	// test default
	obj.SetID("")
	assert.EqualValues(t, "", obj.ID(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](33)
	assert.Panics(t, func() {
		obj.SetID(val)
	})
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
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLogin(ctx, obj)

	// Test retrieval
	obj2, err := LoadLogin(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)
	assert.NoError(t, err)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.IDIsLoaded())
	assert.Panics(t, func() {
		obj2.SetID(obj2.ID())
	})

	assert.True(t, obj2.UsernameIsLoaded())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.usernameIsDirty)
	obj2.SetUsername(obj2.Username())
	assert.False(t, obj2.usernameIsDirty)

	assert.True(t, obj2.PasswordIsLoaded())
	assert.False(t, obj2.PasswordIsNull())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.passwordIsDirty)
	obj2.SetPassword(obj2.Password())
	assert.False(t, obj2.passwordIsDirty)

	assert.True(t, obj2.IsEnabledIsLoaded())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.isEnabledIsDirty)
	obj2.SetIsEnabled(obj2.IsEnabled())
	assert.False(t, obj2.isEnabledIsDirty)

}

func TestLogin_InsertPanics(t *testing.T) {
	obj := createMinimalSampleLogin()
	_ = obj
	ctx := db.NewContext(nil)
	_ = ctx

	obj.usernameIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.usernameIsLoaded = true

	obj.isEnabledIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.isEnabledIsLoaded = true

}

func TestLogin_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleLogin()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLogin(ctx, obj)
	updateMinimalSampleLogin(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2, err := LoadLogin(ctx, obj.PrimaryKey())
	assert.NoError(t, err)

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.Username(), obj.Username(), "Username did not update")
	assert.Equal(t, obj2.Password(), obj.Password(), "Password did not update")
	assert.Equal(t, obj2.IsEnabled(), obj.IsEnabled(), "IsEnabled did not update")
}

func TestLogin_ReferenceLoad(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLogin(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLogin(ctx, obj)

	// Test that referenced objects were saved and assigned ids
	assert.NotNil(t, obj.Person())
	assert.NotEqual(t, '-', obj.Person().PrimaryKey()[0])

	// Test lazy loading
	obj2, err := LoadLogin(ctx, obj.PrimaryKey())
	assert.NoError(t, err)
	objPkOnly, err2 := LoadLogin(ctx, obj.PrimaryKey(), node.Login().PrimaryKey())
	assert.NoError(t, err2)
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	assert.Nil(t, obj2.Person(), "Person is not loaded initially")
	v_Person, _ := obj2.LoadPerson(ctx)
	assert.NotNil(t, v_Person)
	assert.Equal(t, v_Person.PrimaryKey(), obj2.Person().PrimaryKey())
	assert.Equal(t, obj.Person().PrimaryKey(), obj2.Person().PrimaryKey())
	assert.True(t, obj2.PersonIDIsLoaded())

	assert.False(t, objPkOnly.PersonIDIsLoaded())
	assert.Panics(t, func() { _, _ = objPkOnly.LoadPerson(ctx) })

	// test eager loading
	obj3, _ := LoadLogin(ctx, obj.PrimaryKey(), node.Login().Person())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.Person().PrimaryKey(), obj3.Person().PrimaryKey())

}

func TestLogin_ReferenceUpdateNewObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLogin(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLogin(ctx, obj)

	obj2, _ := LoadLogin(ctx, obj.PrimaryKey())
	updateMaximalSampleLogin(ctx, obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleLogin(ctx, obj2)

	obj3, _ := LoadLogin(ctx, obj2.PrimaryKey(), node.Login().Person())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.Person().PrimaryKey(), obj3.Person().PrimaryKey())

}

func TestLogin_ReferenceUpdateOldObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLogin(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLogin(ctx, obj)

	updateMinimalSamplePerson(obj.Person())

	assert.NoError(t, obj.Save(ctx))

	obj2, _ := LoadLogin(ctx, obj.PrimaryKey(),
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

	has, _ := HasLogin(ctx, obj.PrimaryKey())
	assert.True(t, has)

	obj2, _ := LoadLogin(ctx, obj.PrimaryKey(), node.Login().PrimaryKey())

	assert.Equal(t, obj.ID(), obj.Get(node.Login().ID().Identifier))
	assert.Panics(t, func() { obj2.PersonID() })
	assert.Nil(t, obj2.Get(node.Login().PersonID().Identifier))
	assert.Equal(t, obj.Username(), obj.Get(node.Login().Username().Identifier))
	assert.Panics(t, func() { obj2.Username() })
	assert.Nil(t, obj2.Get(node.Login().Username().Identifier))
	assert.Equal(t, obj.Password(), obj.Get(node.Login().Password().Identifier))
	assert.Panics(t, func() { obj2.Password() })
	assert.Nil(t, obj2.Get(node.Login().Password().Identifier))
	assert.Equal(t, obj.IsEnabled(), obj.Get(node.Login().IsEnabled().Identifier))
	assert.Panics(t, func() { obj2.IsEnabled() })
	assert.Nil(t, obj2.Get(node.Login().IsEnabled().Identifier))
}

func TestLogin_QueryLoad(t *testing.T) {
	obj := createMinimalSampleLogin()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLogin(ctx, obj)

	objs, err := QueryLogins(ctx).
		Where(op.Equal(node.Login().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.Login().PrimaryKey()). // exercise order by
		Limit(1, 0).                        // exercise limit
		Calculation(node.Login(), "IsTrue", op.Equal("A", "A")).
		Load()
	assert.NoError(t, err)
	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
	assert.True(t, objs[0].GetAlias("IsTrue").Bool())
}
func TestLogin_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleLogin()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLogin(ctx, obj)

	objs, _ := QueryLogins(ctx).
		Where(op.Equal(node.Login().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestLogin_QueryCursor(t *testing.T) {
	obj := createMinimalSampleLogin()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLogin(ctx, obj)

	cursor, err := QueryLogins(ctx).
		Where(op.Equal(node.Login().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()
	require.NoError(t, err)
	obj2, err2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	require.NoError(t, err2)
	obj2, err2 = cursor.Next()
	assert.Nil(t, obj2)
	require.NoError(t, err2)
	assert.NoError(t, cursor.Close())

	// test empty cursor result
	cursor, err = QueryLogins(ctx).
		Where(op.Equal("B", "A")).
		LoadCursor()
	require.NoError(t, err)

	obj2, err = cursor.Next()
	assert.Nil(t, obj2)
	require.NoError(t, err)
	assert.NoError(t, cursor.Close())
}
func TestLogin_Count(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLogin(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLogin(ctx, obj)
	assert.Positive(t, func() int { i, _ := CountLogins(ctx); return i }())

	// reread in case there are data limitations imposed by the database
	obj2, _ := LoadLogin(ctx, obj.PrimaryKey())
	assert.Positive(t,
		func() int {
			i, _ := CountLoginsByPersonID(ctx,
				obj2.PersonID())
			return i
		}())
	assert.Positive(t,
		func() int {
			i, _ := CountLoginsByUsername(ctx,
				obj2.Username())
			return i
		}())

}

func TestLogin_MarshalJSON(t *testing.T) {
	obj := createMinimalSampleLogin()

	b, err := json.Marshal(obj)
	assert.NoError(t, err)

	obj2 := NewLogin()
	err = json.Unmarshal(b, &obj2)
	assert.NoError(t, err)

	assertEqualFieldsLogin(t, obj, obj2)
}

func TestLogin_MarshalBinary(t *testing.T) {
	obj := createMinimalSampleLogin()

	b, err := obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 := NewLogin()
	err = obj2.UnmarshalBinary(b)
	assert.NoError(t, err)

	assertEqualFieldsLogin(t, obj, obj2)
}

func TestLogin_FailingMarshalBinary(t *testing.T) {
	obj := createMinimalSampleLogin()
	var err error

	for i := 0; i < 20; i++ {
		enc := &test.GobEncoder{Count: i}
		err = obj.encodeTo(enc)
		assert.Error(t, err)
	}
	// do it again with aliases
	obj._aliases = make(map[string]any)
	for i := 0; i < 21; i++ {
		enc := &test.GobEncoder{Count: i}
		err = obj.encodeTo(enc)
		assert.Error(t, err)
	}
}

func TestLogin_FailingUnmarshalBinary(t *testing.T) {
	obj := createMinimalSampleLogin()
	b, err := obj.MarshalBinary()
	assert.NoError(t, err)
	obj2 := NewLogin()
	for i := 0; i < 20; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}

	// do it again with aliases
	obj = createMinimalSampleLogin()
	obj._aliases = map[string]any{"a": 1}
	b, err = obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 = NewLogin()
	for i := 0; i < 21; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}
}

func TestLogin_Indexes(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLogin(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLogin(ctx, obj)

	var obj2 *Login
	obj2, _ = LoadLoginByPersonID(ctx, obj.PersonID())
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.True(t, func() bool { h, _ := HasLoginByPersonID(ctx, obj.PersonID()); return h }())

	obj2, _ = LoadLoginByUsername(ctx, obj.Username())
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.True(t, func() bool { h, _ := HasLoginByUsername(ctx, obj.Username()); return h }())

}
