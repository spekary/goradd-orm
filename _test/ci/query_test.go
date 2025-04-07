package ci

import (
	"encoding/json"
	"github.com/goradd/orm/_test/gen/orm/goradd"
	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/_test/gen/orm/goradd_unit"
	node2 "github.com/goradd/orm/_test/gen/orm/goradd_unit/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	ctx := db.NewContext(nil)

	people := goradd.QueryPeople(ctx).
		OrderBy(node.Person().ID()).
		Load()
	if len(people) != 12 {
		t.Error("12 people not found")
	}
	if people[0].FirstName() != "John" {
		t.Error("First person is not John")
	}
}

func TestLoad(t *testing.T) {
	ctx := db.NewContext(nil)

	people := goradd.QueryPeople(ctx).
		Where(op.Equal(node.Person().ID(), "7")).
		Select(
			node.Person().Projects(),        // many many
			node.Person().ManagerProjects(), // reverse
			node.Person().Login(),
		).
		Load()

	person := people[0]
	assert.Equal(t, "Wolfe", person.LastName(), "Last name was not selected by default")
	assert.Len(t, person.Projects(), 2)
	assert.Len(t, person.ManagerProjects(), 2)
	assert.Equal(t, person.Types().Len(), 2)
	assert.Equal(t, "kwolfe", person.Login().Username())

	person = goradd.LoadPerson(ctx, "7",
		node.Person().Projects(),        // many many
		node.Person().ManagerProjects(), // reverse
		node.Person().Login(),           // reverse unique
	)

	// Serialize and deserialize
	assert.Len(t, person.Projects(), 2)
	assert.Len(t, person.ManagerProjects(), 2)
	assert.Equal(t, person.Types().Len(), 2)
	assert.Equal(t, "kwolfe", person.Login().Username())
}

func TestSort(t *testing.T) {
	ctx := db.NewContext(nil)
	people := goradd.QueryPeople(ctx).
		OrderBy(node.Person().LastName()).
		Load()

	if people[0].LastName() != "Brady" {
		t.Error("Person found not Brady, found " + people[0].LastName())
	}

	people = goradd.QueryPeople(ctx).
		OrderBy(node.Person().FirstName()).
		Load()
	if people[0].FirstName() != "Alex" {
		t.Error("Person found not Alex, found " + people[0].FirstName())
	}

	// Testing for regression bug with multiple sorts
	people = goradd.QueryPeople(ctx).
		OrderBy(node.Person().LastName().Descending(), node.Person().FirstName().Ascending()).
		Load()
	if people[0].FirstName() != "Karen" || people[0].LastName() != "Wolfe" {
		t.Error("Person found not Karen Wolfe, found " + people[0].FirstName() + " " + people[0].LastName())
	}
}

func TestWhere(t *testing.T) {
	ctx := db.NewContext(nil)
	people := goradd.QueryPeople(ctx).
		Where(op.Equal(node.Person().LastName(), "Smith")).
		OrderBy(node.Person().FirstName().Descending(), node.Person().LastName()).
		Load()

	if people[0].FirstName() != "Wendy" {
		t.Error("Person found not Wendy, found " + people[0].FirstName())
	}
}

func TestAlias(t *testing.T) {
	ctx := db.NewContext(nil)
	projects := goradd.QueryProjects(ctx).
		Where(op.Equal(node.Project().ID(), 1)).
		Calculation(node.Project(), "Difference", op.Subtract(node.Project().Budget(), node.Project().Spent())).
		Load()

	v := projects[0].GetAlias("Difference").Float()
	assert.EqualValues(t, -690.5, v)
}

func TestCursor(t *testing.T) {
	ctx := db.NewContext(nil)
	projectCursor := goradd.QueryProjects(ctx).
		LoadCursor()

	var projects []*goradd.Project
	for project := projectCursor.Next(); project != nil; project = projectCursor.Next() {
		projects = append(projects, project)
	}

	assert.Len(t, projects, 4)
}

/*
	func TestAlias2(t *testing.T) {
		ctx := db.NewContext(nil)
		projects := goradd.QueryProjects(ctx).
			Alias("a", node.Project().Num()).
			Alias("b", node.Project().Name()).
			Alias("c", node.Project().Spent()).
			Alias("d", node.Project().StartDate()).
			Alias("e", op.IsNull(node.Project().EndDate())).
			OrderBy(node.Project().ID()).
			Load()

		project := projects[0]
		assert.Equal(t, 1, project.GetAlias("a").Int())
		assert.Equal(t, 1, project.Num())
		assert.Equal(t, "ACME Website Redesign", project.GetAlias("b").String())
		assert.Equal(t, 10250.75, project.GetAlias("c").Float())
		d := time.FromSqlDateTime("2004-03-01")
		assert.EqualValues(t, d, project.GetAlias("d").Time())
		//assert.EqualValues(t, d, project.StartDate())
		assert.Equal(t, false, project.GetAlias("e").Bool())
	}
*/
func TestCount(t *testing.T) {
	ctx := db.NewContext(nil)

	count := goradd.QueryProjects(ctx).
		Count()

	assert.EqualValues(t, 4, count)
}

func TestGroupBy(t *testing.T) {
	ctx := db.NewContext(nil)

	projects := goradd.QueryProjects(ctx).
		Calculation(node.Project(), "teamMemberCount", op.Count(node.Project().TeamMembers())).
		GroupBy(node.Project()).
		OrderBy(node.Project().ID()).
		Load()

	assert.EqualValues(t, 5, projects[0].GetAlias("teamMemberCount").Int())
}

func TestSelect(t *testing.T) {
	ctx := db.NewContext(nil)

	projects := goradd.QueryProjects(ctx).
		Select(node.Project().Name()).
		Load()

	project := projects[0]
	assert.True(t, project.NameIsLoaded())
	assert.False(t, project.ManagerIDIsLoaded())
	assert.True(t, project.IDIsLoaded())
}

func TestLimit(t *testing.T) {
	ctx := db.NewContext(nil)

	people := goradd.QueryPeople(ctx).
		OrderBy(node.Person().ID()).
		Limit(2, 3).
		Load()

	assert.EqualValues(t, "Mike", people[0].FirstName())
	assert.Len(t, people, 2)
}

func TestSaveAndDelete(t *testing.T) {
	ctx := db.NewContext(nil)

	person := goradd.NewPerson()
	person.SetFirstName("Test1")
	person.SetLastName("Last1")
	person.Save(ctx)

	people := goradd.QueryPeople(ctx).
		Where(
			op.And(
				op.Equal(
					node.Person().FirstName(), "Test1"),
				op.Equal(
					node.Person().LastName(), "Last1"))).
		Load()

	assert.EqualValues(t, person.ID(), people[0].ID())

	people[0].Delete(ctx)

	people = goradd.QueryPeople(ctx).
		Where(
			op.And(
				op.Equal(
					node.Person().FirstName(), "Test1"),
				op.Equal(
					node.Person().LastName(), "Last1"))).
		Load()

	assert.Len(t, people, 0, "Deleted the person")
}

func TestSingleEmpty(t *testing.T) {
	ctx := db.NewContext(nil)

	people := goradd.QueryPeople(ctx).
		Where(op.Equal(node.Person().ID(), 12345)).
		Load()

	assert.Len(t, people, 0)

}

func TestLazyLoad(t *testing.T) {
	ctx := db.NewContext(nil)

	projects := goradd.QueryProjects(ctx).
		Where(op.Equal(node.Project().ID(), 1)).
		Load()

	var mId string = projects[0].ID() // foreign keys are treated as strings for cross-database compatibility
	assert.Equal(t, "1", mId)

	manager := projects[0].LoadManager(ctx)
	assert.Equal(t, "7", manager.ID())
}

func TestHaving(t *testing.T) {
	// This particular test shows a quirk of SQL that requires:
	// 1) If you have an aggregate clause (like COUNT), you MUST have a GROUPBY clause, and
	// 2) If you have a GROUPBY, you MUST SELECT and only select the things you are grouping by.
	//
	// Sooo, when we see a GroupBy, we automatically also select the same nodes.
	ctx := db.NewContext(nil)

	projects := goradd.QueryProjects(ctx).
		GroupBy(node.Project().ID(), node.Project().Name()).
		OrderBy(node.Project().ID()).
		Calculation(node.Project(), "team_member_count", op.Count(node.Project().TeamMembers())).
		Having(op.GreaterThan(node.Alias("team_member_count"), 5)).
		Load()

	assert.Len(t, projects, 2)
	assert.Equal(t, "State College HR System", projects[0].Name())
	assert.Equal(t, 6, projects[0].GetAlias("team_member_count").Int())
}

func TestFailedSelects(t *testing.T) {
	ctx := db.NewContext(nil)

	assert.Panics(t, func() { goradd.QueryProjects(ctx).Select(node.Person()) })
}

func TestFailedGroupBy(t *testing.T) {
	ctx := db.NewContext(nil)

	assert.Panics(t, func() {
		goradd.
			QueryProjects(ctx).
			GroupBy(node.Project().Name()).
			Select(node.Project().Name())
	})
}

// Test that we can get from an integer keyed database
func TestIntKey(t *testing.T) {
	ctx := db.NewContext(nil)

	g := goradd.LoadGift(ctx, 2)
	assert.Equal(t, "Turtle doves", g.Name())
}

func TestMultiParent(t *testing.T) {
	ctx := db.NewContext(nil)

	baby := goradd_unit.NewMultiParent()
	baby.SetName("Baby")
	mom := goradd_unit.NewMultiParent()
	mom.SetName("Mom")
	dad := goradd_unit.NewMultiParent()
	dad.SetName("Dad")

	baby.SetParent1(mom)
	baby.SetParent2(dad)
	baby.Save(ctx)
	defer baby.Delete(ctx)
	defer mom.Delete(ctx)
	defer dad.Delete(ctx)

	mom2 := goradd_unit.LoadMultiParent(ctx, mom.ID(), node2.MultiParent().Parent1MultiParents())
	assert.Equal(t, mom2.Parent1MultiParents()[0].ID(), baby.ID())

	b, err := json.Marshal(baby)
	assert.NoError(t, err)

	baby2 := goradd_unit.NewMultiParent()
	err = baby2.UnmarshalJSON(b)
	assert.NoError(t, err)
	assert.Equal(t, baby.ID(), baby2.ID())
	assert.Equal(t, baby.Parent1().ID(), baby2.Parent1().ID())
	assert.Equal(t, baby.Parent2().ID(), baby2.Parent2().ID())
}

func TestTimeDefaults(t *testing.T) {
	ctx := db.NewContext(nil)
	p := goradd.NewPerson()
	defer p.Delete(ctx)
	p.SetFirstName("Bob")
	p.SetLastName("Dylan")
	assert.NoError(t, p.Save(ctx))
	now := time.Now()
	assert.WithinDuration(t, now, p.Created(), time.Second)
	assert.WithinDuration(t, p.Created(), p.Modified(), time.Second)
	time.Sleep(2 * time.Second)
	p.SetLastName("Hope")
	assert.NoError(t, p.Save(ctx))
	assert.WithinDuration(t, now, p.Created(), time.Second) // no change
	assert.WithinDuration(t, time.Now(), p.Modified(), time.Second)
}
