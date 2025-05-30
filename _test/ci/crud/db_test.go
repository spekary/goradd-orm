package crud

import (
	"fmt"
	"github.com/goradd/orm/_test/config"
	"github.com/goradd/orm/_test/gen/orm/goradd_unit"
	"github.com/goradd/orm/pkg/db"
	"os"
	"testing"
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

	config.InitDB()
	ctx := db.NewContext(nil)
	goradd_unit.ClearAll(ctx)
}

func teardown() {
	// Cleanup logic here
	fmt.Println("Cleaning up after tests...")
}
