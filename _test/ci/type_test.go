package ci

import (
	"github.com/goradd/orm/_test/gen/orm/goradd_unit"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/strings"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAutoGeneratedValues(t *testing.T) {
	ctx := db.NewContext(nil)
	defer goradd_unit.ClearAll(ctx)

	r := goradd_unit.NewAutoGen()
	r.SetName("test")

	id := r.ID()
	v := strings.AtoI[int](id)
	assert.True(t, v < 0)
	assert.True(t, r.GroLock() == 0)
	assert.True(t, r.GroTimestamp() == 0)
	assert.True(t, r.Created().IsZero())
	assert.True(t, r.Modified().IsZero())
	err := r.Save(ctx)
	assert.NoError(t, err)

	id = r.ID()
	v = strings.AtoI[int](id)
	assert.True(t, v >= 0)

	lock := r.GroLock()
	ts := r.GroTimestamp()
	assert.True(t, lock != 0)
	assert.True(t, ts != 0)
	assert.False(t, r.Created().IsZero())
	assert.False(t, r.Modified().IsZero())
	assert.WithinDuration(t, r.Created(), r.Modified(), time.Millisecond)

	time.Sleep(2 * time.Millisecond)
	r.SetName("test2")
	err = r.Save(ctx)
	assert.NoError(t, err)

	assert.True(t, r.Created().Add(time.Millisecond).Before(r.Modified()))
	assert.NotEqual(t, lock, r.GroLock())
	assert.NotEqual(t, ts, r.GroTimestamp())
	assert.EqualValues(t, id, r.ID())
}
