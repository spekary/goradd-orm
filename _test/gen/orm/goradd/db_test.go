// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/goradd/orm/pkg/test"
)

// ClearAll deletes all the data in the database, except for data in Enum tables.
func ClearAll(ctx context.Context) {
	db := Database()

	db.Delete(ctx, "related_project_assn", nil)
	db.Delete(ctx, "team_member_project_assn", nil)

	db.Delete(ctx, "project", nil)
	db.Delete(ctx, "person_with_lock", nil)
	db.Delete(ctx, "person", nil)
	db.Delete(ctx, "milestone", nil)
	db.Delete(ctx, "login", nil)
	db.Delete(ctx, "gift", nil)
	db.Delete(ctx, "employee_info", nil)
	db.Delete(ctx, "address", nil)

}

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

	fmt.Println(test.NewSeed())

	// uncomment the next line to re-use a previously created seed to recreate a test
	// UseSeed(seed)

	InitDB()
}

func teardown() {
	// Cleanup logic here
	fmt.Println("Cleaning up after tests...")
}
