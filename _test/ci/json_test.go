package ci

import (
	"encoding/json"
	"github.com/goradd/orm/_test/gen/orm/goradd"
	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/_test/gen/orm/goradd_unit"
	"github.com/goradd/orm/pkg/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonMarshall1(t *testing.T) {
	ctx := db.NewContext(nil)

	p := goradd.LoadProject(ctx, "1",
		node.Project().Name(),
		node.Project().Status(),
		node.Project().Manager().FirstName())
	j, err := json.Marshal(p)
	assert.NoError(t, err)
	m := make(map[string]interface{})
	err = json.Unmarshal(j, &m)
	assert.NoError(t, err)
	assert.Exactly(t, "ACME Website Redesign", m["name"])
	assert.Exactly(t, "Karen", m["manager"].(map[string]interface{})["firstName"])
}

func TestJsonUnmarshall1(t *testing.T) {
	p := goradd.NewProject()
	err := json.Unmarshal([]byte(
		`{
	"name":"ACME Website Redesign",
	"status":3,
	"num":14,
	"startDate":"2020-11-01T00:00:00Z"
}
`),
		&p)
	assert.NoError(t, err)
	assert.Exactly(t, "ACME Website Redesign", p.Name())
	assert.Exactly(t, goradd.ProjectStatusCompleted, p.Status())
	assert.Exactly(t, 14, p.Num())
	assert.Exactly(t, 2020, p.StartDate().Year())
}

func TestJsonMarshall2(t *testing.T) {
	ctx := db.NewContext(nil)

	p := goradd.LoadPerson(ctx, "1",
		node.Person().FirstName(),
		node.Person().LastName(),
		node.Person().Types())
	j, err := json.Marshal(p)
	assert.NoError(t, err)
	m := make(map[string]interface{})
	err = json.Unmarshal(j, &m)
	assert.NoError(t, err)
	assert.Equal(t, "John", m["firstName"])
	assert.Equal(t, "Doe", m["lastName"])
	assert.ElementsMatch(t, []string{"manager", "inactive"}, m["types"])
}

func TestJsonUnmarshall2(t *testing.T) {
	p := goradd.NewPerson()
	err := json.Unmarshal([]byte(
		`{
	"firstName":"John",
	"lastName":"Doe",
	"types":[2, 3]
}
`),
		&p)
	assert.NoError(t, err)
	assert.Equal(t, "John", p.FirstName())
	assert.Equal(t, "Doe", p.LastName())
	assert.ElementsMatch(t, []goradd.PersonType{goradd.PersonTypeManager, goradd.PersonTypeInactive}, p.Types().Values())
}

func TestJsonMarshall3(t *testing.T) {
	ctx := db.NewContext(nil)

	r := goradd_unit.LoadTypeTest(ctx, "1")
	j, err := json.Marshal(r)
	assert.NoError(t, err)

	r2 := goradd_unit.NewTypeTest()
	err = r2.UnmarshalJSON(j)
	assert.NoError(t, err)

	assert.Equal(t, r.TestInt(), r2.TestInt())
	assert.Equal(t, string(r.TestBlob()), string(r2.TestBlob()))
}

func TestJsonMarshallReferences(t *testing.T) {
	ctx := db.NewContext(nil)
	project := goradd.LoadProject(ctx, "1", node.Project().Manager())

	b, err := json.Marshal(project)
	assert.NoError(t, err)

	project2 := goradd.NewProject()
	err = project2.UnmarshalJSON(b)
	assert.NoError(t, err)
	assert.Equal(t, project.ID(), project2.ID())
	assert.Equal(t, project.Manager().ID(), project2.Manager().ID())
}
