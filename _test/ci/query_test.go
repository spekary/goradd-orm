package ci

import (
	"github.com/goradd/goradd/pkg/time"
	"github.com/goradd/orm/_test/gen/orm/goradd"
	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/query"
	"github.com/stretchr/testify/assert"
	"slices"
	"strconv"
	"testing"
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

func TestReference(t *testing.T) {
	ctx := db.NewContext(nil)
	projects := goradd.QueryProjects(ctx).
		Join(node.Project().Manager()).
		OrderBy(node.Project().ID()).
		Load()

	if projects[0].Manager().FirstName() != "Karen" {
		t.Error("Person found not Karen, found " + projects[0].Manager().FirstName())
	}

}

func TestManyMany(t *testing.T) {
	ctx := db.NewContext(nil)
	projects := goradd.QueryProjects(ctx).
		Join(node.Project().TeamMembers()).
		OrderBy(node.Project().ID()).
		Load()

	if len(projects[0].TeamMembers()) != 5 {
		t.Error("Did not find 5 team members in project 1. Found: " + strconv.Itoa(len(projects[0].TeamMembers())))
	}

}

func TestReverseReference(t *testing.T) {
	ctx := db.NewContext(nil)
	people := goradd.QueryPeople(ctx).
		Join(node.Person().ManagerProjects()).
		OrderBy(node.Person().ID()).
		Load()

	if people[0].FirstName() != "John" {
		t.Error("Did not find person 0.")
	}

	if len(people[6].ManagerProjects()) != 2 {
		t.Error("Did not find 2 ProjectsAsManagers.")
	}

}

func TestBasicEnum(t *testing.T) {
	ctx := db.NewContext(nil)
	projects := goradd.QueryProjects(ctx).
		OrderBy(node.Project().ID()).
		Load()

	if projects[0].Status() != goradd.ProjectStatusCompleted {
		t.Error("Did not find correct project type.")
	}
}

func TestManyEnum(t *testing.T) {
	ctx := db.NewContext(nil)
	people := goradd.QueryPeople(ctx).
		OrderBy(node.Person().ID()).
		Join(node.Person().PersonTypes()).
		Load()

	if len(people[0].PersonTypes()) != 2 {
		t.Error("Did not expand to 2 person types.")
	}

	if !slices.Contains(people[0].PersonTypes(), goradd.PersonTypeInactive) {
		t.Error("Did not find correct person type.")
	}
}

func TestManyManySingles(t *testing.T) {
	ctx := db.NewContext(nil)
	projects := goradd.QueryProjects(ctx).
		Expand(node.Project().TeamMembers()).
		OrderBy(node.Project().ID(), node.Project().TeamMembers().FirstName()).
		Load()

	if projects[4].Name() != "ACME Website Redesign" { // should have 5 lines here that are all project 1
		t.Error("Did not find expanded project ACME Website Redesign.")
	}

	if projects[5].Name() != "State College HR System" { // should have 5 lines here that are all project 1
		t.Error("Did not find expanded project State College HR System.")
	}

	if projects[3].TeamMembers()[0].FirstName() != "Samantha" {
		t.Error("Did not find Samantha. Found: " + projects[3].TeamMembers()[0].FirstName())
	}
}

func TestReverseReferenceSingles(t *testing.T) {
	ctx := db.NewContext(nil)
	people := goradd.QueryPeople(ctx).
		Expand(node.Person().ManagerProjects()).
		OrderBy(node.Person().ID()).
		Load()

	if people[7].FirstName() != "Karen" {
		t.Error("Did not find expanded person Karen Wolfe.")
	}

	if people[7].ManagerProjects()[0].ManagerID() != people[7].ID() {
		t.Error("Did not find Project 2.")
	}
}

func TestManyEnumSingles(t *testing.T) {
	ctx := db.NewContext(nil)
	people := goradd.QueryPeople(ctx).
		OrderBy(node.Person().ID(), node.Person().PersonTypes().ID().Descending()).
		Expand(node.Person().PersonTypes()).
		Load()

	if people[1].PersonTypes()[0] != goradd.PersonTypeManager {
		t.Error("Did not find correct person type.")
	}

}

func TestAlias(t *testing.T) {
	ctx := db.NewContext(nil)
	projects := goradd.QueryProjects(ctx).
		Where(op.Equal(node.Project().ID(), 1)).
		Alias("Difference", op.Subtract(node.Project().Budget(), node.Project().Spent())).
		Load()

	v := projects[0].GetAlias("Difference").Float()
	assert.EqualValues(t, -690.5, v)
}

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

func TestCount(t *testing.T) {
	ctx := db.NewContext(nil)

	count := goradd.QueryProjects(ctx).
		Count(false)

	assert.EqualValues(t, 4, count)
}

func TestGroupBy(t *testing.T) {
	ctx := db.NewContext(nil)

	projects := goradd.QueryProjects(ctx).
		Alias("teamMemberCount", op.Count(node.Project().TeamMembers())).
		GroupBy(node.Project()).
		Load()

	assert.EqualValues(t, 5, projects[0].GetAlias("teamMemberCount").Int())

}

func TestSelect(t *testing.T) {
	ctx := db.NewContext(nil)

	projects := goradd.QueryProjects(ctx).
		Select(node.Project().Name()).
		Load()

	project := projects[0]
	assert.True(t, project.NameIsValid())
	assert.False(t, project.ManagerIDIsValid())
	assert.True(t, project.IDIsValid())
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

func TestDeleteQuery(t *testing.T) {
	ctx := db.NewContext(nil)

	person := goradd.NewPerson()
	person.SetFirstName("Test1")
	person.SetLastName("Last1")
	person.Save(ctx)

	goradd.QueryPeople(ctx).
		Where(
			op.And(
				op.Equal(
					node.Person().FirstName(), "Test1"),
				op.Equal(
					node.Person().LastName(), "Last1"))).
		Delete()

	people := goradd.QueryPeople(ctx).
		Where(
			op.And(
				op.Equal(
					node.Person().FirstName(), "Test1"),
				op.Equal(
					node.Person().LastName(), "Last1"))).
		Load()

	assert.Len(t, people, 0, "Deleted the person")
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
		Alias("team_member_count", op.Count(node.Project().TeamMembers())).
		Having(op.GreaterThan(query.Alias("team_member_count"), 5)).
		Load()

	assert.Len(t, projects, 2)
	assert.Equal(t, "State College HR System", projects[0].Name())
	assert.Equal(t, 6, projects[0].GetAlias("team_member_count").Int())
}

func TestFailedJoins(t *testing.T) {
	ctx := db.NewContext(nil)

	assert.Panics(t, func() { goradd.QueryProjects(ctx).Join(node.Person()) })
	assert.Panics(t, func() { goradd.QueryProjects(ctx).Join(node.Project().ManagerID()) })
}

func TestFailedExpand(t *testing.T) {
	ctx := db.NewContext(nil)

	assert.Panics(t, func() { goradd.QueryProjects(ctx).Expand(node.Person()) })
	assert.Panics(t, func() { goradd.QueryProjects(ctx).Expand(node.Project().Manager()) })
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