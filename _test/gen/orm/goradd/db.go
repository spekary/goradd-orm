// Code generated by goradd-orm. DO NOT EDIT.

// Package goradd contains the object relational model for the goradd database.
//
// Queries use a builder pattern, started with a Query* function. Add functions to the builder to further constrain the query,
// using nodes from the [node] package to refer to tables and columns in the database. End the query with either a Load call to get a
// list of items, a Get call to get one item, or a Count call to count the number of items in the query.
//
// Some Examples
//
//	projects := model.QueryProjects().Load()
//
// Returns all the projects in the database.
//
//	projects := model.QueryProjects().
//	    Join(node.Project().Manager()).
//	    Where(op.GreaterOrEqual(node.Project().StartDate(), time.NewDate(2006, 1, 1)).
//	    OrderBy(node.Project().Num()).
//	    Load()
//
// Returns the projects that started in 2006 or later, with the manager objects attached, and ordered by project number.
// To get the manager of the first project returned, you can do this:
//
//	firstManager := projects[0].Manager()
//
// See the goradd-orm documentation for more information.
package goradd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/query"
)

// Database returns the database object corresponding to "goradd" in the global database cluster.
// Use this to call directly into the database through the DatabaseI interface, or if you want to call functions
// specific to the database, cast the interface to the driver.
func Database() db.DatabaseI {
	return db.GetDatabase("goradd")
}

// ClearAll deletes all the data in the database, except for data in Enum tables.
func ClearAll(ctx context.Context) {
	db := Database()

	_ = db.Delete(ctx, "related_project_assn", nil)
	_ = db.Delete(ctx, "team_member_project_assn", nil)

	_ = db.Delete(ctx, "milestone", nil)
	_ = db.Delete(ctx, "login", nil)
	_ = db.Delete(ctx, "employee_info", nil)
	_ = db.Delete(ctx, "address", nil)
	_ = db.Delete(ctx, "project", nil)
	_ = db.Delete(ctx, "person_with_lock", nil)
	_ = db.Delete(ctx, "person", nil)
	_ = db.Delete(ctx, "gift", nil)

}

// JsonEncodeAll sends the entire database to writer as JSON.
func JsonEncodeAll(ctx context.Context, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	if _, err := io.WriteString(writer, "[\n"); err != nil {
		return fmt.Errorf("writer error: %w", err)
	}

	{ // Write Gifts
		if _, err := io.WriteString(writer, "["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, `"gift"`); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		cursor, err := QueryGifts(ctx).LoadCursor()
		if err != nil {
			return fmt.Errorf("query error: %w", err)
		}
		db.WalkCursor(cursor, func(i int, item *Gift) error {
			
		})
		defer cursor.Close()
		obj, err2 := cursor.Next()
		if err2 != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}
		if obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}

		for obj, err = cursor.Next(); obj != nil && err == nil; obj, err = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return fmt.Errorf("writer error: %w", err)
			}
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}
		if err != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
	}
	{ // Write People
		if _, err := io.WriteString(writer, "["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, `"person"`); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		cursor, err := QueryPeople(ctx).LoadCursor()
		if err != nil {
			return fmt.Errorf("query error: %w", err)
		}
		defer cursor.Close()
		obj, err2 := cursor.Next()
		if err2 != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}
		if obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}

		for obj, err = cursor.Next(); obj != nil && err == nil; obj, err = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return fmt.Errorf("writer error: %w", err)
			}
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}
		if err != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
	}
	{ // Write PersonWithLocks
		if _, err := io.WriteString(writer, "["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, `"person_with_lock"`); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		cursor, err := QueryPersonWithLocks(ctx).LoadCursor()
		if err != nil {
			return fmt.Errorf("query error: %w", err)
		}
		defer cursor.Close()
		obj, err2 := cursor.Next()
		if err2 != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}
		if obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}

		for obj, err = cursor.Next(); obj != nil && err == nil; obj, err = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return fmt.Errorf("writer error: %w", err)
			}
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}
		if err != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
	}
	{ // Write Projects
		if _, err := io.WriteString(writer, "["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, `"project"`); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		cursor, err := QueryProjects(ctx).LoadCursor()
		if err != nil {
			return fmt.Errorf("query error: %w", err)
		}
		defer cursor.Close()
		obj, err2 := cursor.Next()
		if err2 != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}
		if obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}

		for obj, err = cursor.Next(); obj != nil && err == nil; obj, err = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return fmt.Errorf("writer error: %w", err)
			}
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}
		if err != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
	}
	{ // Write Addresses
		if _, err := io.WriteString(writer, "["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, `"address"`); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		cursor, err := QueryAddresses(ctx).LoadCursor()
		if err != nil {
			return fmt.Errorf("query error: %w", err)
		}
		defer cursor.Close()
		obj, err2 := cursor.Next()
		if err2 != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}
		if obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}

		for obj, err = cursor.Next(); obj != nil && err == nil; obj, err = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return fmt.Errorf("writer error: %w", err)
			}
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}
		if err != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
	}
	{ // Write EmployeeInfos
		if _, err := io.WriteString(writer, "["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, `"employee_info"`); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		cursor, err := QueryEmployeeInfos(ctx).LoadCursor()
		if err != nil {
			return fmt.Errorf("query error: %w", err)
		}
		defer cursor.Close()
		obj, err2 := cursor.Next()
		if err2 != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}
		if obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}

		for obj, err = cursor.Next(); obj != nil && err == nil; obj, err = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return fmt.Errorf("writer error: %w", err)
			}
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}
		if err != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
	}
	{ // Write Logins
		if _, err := io.WriteString(writer, "["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, `"login"`); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		cursor, err := QueryLogins(ctx).LoadCursor()
		if err != nil {
			return fmt.Errorf("query error: %w", err)
		}
		defer cursor.Close()
		obj, err2 := cursor.Next()
		if err2 != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}
		if obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}

		for obj, err = cursor.Next(); obj != nil && err == nil; obj, err = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return fmt.Errorf("writer error: %w", err)
			}
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}
		if err != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
	}
	{ // Write Milestones
		if _, err := io.WriteString(writer, "["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, `"milestone"`); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		cursor, err := QueryMilestones(ctx).LoadCursor()
		if err != nil {
			return fmt.Errorf("query error: %w", err)
		}
		defer cursor.Close()
		obj, err2 := cursor.Next()
		if err2 != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}
		if obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}

		for obj, err = cursor.Next(); obj != nil && err == nil; obj, err = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return fmt.Errorf("writer error: %w", err)
			}
			if err := encoder.Encode(obj); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}
		if err != nil {
			return fmt.Errorf("database cursor error: %w", err)
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		if _, err := io.WriteString(writer, "\n"); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
	}

	db := Database()
	{
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, `"related_project_assn",[`); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		cursor, err := db.Query(ctx, "related_project_assn",
			map[string]query.ReceiverType{
				"child_id":  query.ColTypeString,
				"parent_id": query.ColTypeString,
			},
			nil,
			nil)
		if err != nil {
			return fmt.Errorf("query error: %w", err)
		}
		if rec, err2 := cursor.Next(); err2 != nil {
			return fmt.Errorf("database cursor error: %w", err2)
		} else if rec != nil {
			if err = encoder.Encode(rec); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}
		for {
			rec, err2 := cursor.Next()
			if err2 != nil {
				return fmt.Errorf("database cursor error: %w", err2)
			}
			if rec == nil {
				break
			}

			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return fmt.Errorf("writer error: %w", err)
			}
			if err := encoder.Encode(rec); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}
		if _, err := io.WriteString(writer, `]]`); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
	}
	{
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
		if _, err := io.WriteString(writer, `"team_member_project_assn",[`); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}

		cursor, err := db.Query(ctx, "team_member_project_assn",
			map[string]query.ReceiverType{
				"project_id":     query.ColTypeString,
				"team_member_id": query.ColTypeString,
			},
			nil,
			nil)
		if err != nil {
			return fmt.Errorf("query error: %w", err)
		}
		if rec, err2 := cursor.Next(); err2 != nil {
			return fmt.Errorf("database cursor error: %w", err2)
		} else if rec != nil {
			if err = encoder.Encode(rec); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}
		for {
			rec, err2 := cursor.Next()
			if err2 != nil {
				return fmt.Errorf("database cursor error: %w", err2)
			}
			if rec == nil {
				break
			}

			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return fmt.Errorf("writer error: %w", err)
			}
			if err := encoder.Encode(rec); err != nil {
				return fmt.Errorf("encoding error: %w", err)
			}
		}
		if _, err := io.WriteString(writer, `]]`); err != nil {
			return fmt.Errorf("writer error: %w", err)
		}
	}

	if _, err := io.WriteString(writer, "]"); err != nil {
		return fmt.Errorf("writer error: %w", err)
	}
	return nil
}

// JsonDecodeAll imports the entire database from JSON that was created using JsonEncodeAll.
func JsonDecodeAll(ctx context.Context, reader io.Reader) error {
	decoder := json.NewDecoder(reader)

	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the JSON to start with an array")
		return err
	}

	for decoder.More() {
		if err := jsonDecodeTable(ctx, decoder); err != nil {
			return err
		}
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}

func jsonDecodeTable(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the duple to start with an array")
		return err
	}

	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading name of table:", err)
		return err
	}

	if tableName, ok := token.(string); !ok {
		fmt.Println("Error: Expected a name of a table.")
		return err
	} else {
		switch tableName {
		case "gift":
			err = jsonDecodeGifts(ctx, decoder)
		case "person":
			err = jsonDecodePeople(ctx, decoder)
		case "person_with_lock":
			err = jsonDecodePersonWithLocks(ctx, decoder)
		case "project":
			err = jsonDecodeProjects(ctx, decoder)
		case "address":
			err = jsonDecodeAddresses(ctx, decoder)
		case "employee_info":
			err = jsonDecodeEmployeeInfos(ctx, decoder)
		case "login":
			err = jsonDecodeLogins(ctx, decoder)
		case "milestone":
			err = jsonDecodeMilestones(ctx, decoder)
		case "related_project_assn":
			err = jsonDecodeRelatedProjectAssn(ctx, decoder)
		case "team_member_project_assn":
			err = jsonDecodeTeamMemberProjectAssn(ctx, decoder)

		}
		if err != nil {
			return err
		}
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}

func jsonDecodeGifts(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the Gift list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewGift()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		if err = obj.Save(ctx); err != nil {
			return err
		}
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}
func jsonDecodePeople(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the Person list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewPerson()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		if err = obj.Save(ctx); err != nil {
			return err
		}
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}
func jsonDecodePersonWithLocks(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the PersonWithLock list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewPersonWithLock()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		if err = obj.Save(ctx); err != nil {
			return err
		}
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}
func jsonDecodeProjects(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the Project list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewProject()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		if err = obj.Save(ctx); err != nil {
			return err
		}
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}
func jsonDecodeAddresses(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the Address list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewAddress()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		if err = obj.Save(ctx); err != nil {
			return err
		}
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}
func jsonDecodeEmployeeInfos(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the EmployeeInfo list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewEmployeeInfo()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		if err = obj.Save(ctx); err != nil {
			return err
		}
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}
func jsonDecodeLogins(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the Login list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewLogin()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		if err = obj.Save(ctx); err != nil {
			return err
		}
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}
func jsonDecodeMilestones(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the Milestone list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewMilestone()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		if err = obj.Save(ctx); err != nil {
			return err
		}
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}

func jsonDecodeRelatedProjectAssn(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the RelatedProjectAssn list to start with an array")
		return err
	}

	database := Database()
	for decoder.More() {
		var imp struct {
			Src  string `json:"child_id"`
			Dest string `json:"parent_id"`
		}

		if err = decoder.Decode(&imp); err != nil {
			return err
		}
		db.Associate(ctx, database, "related_project_assn", "child_id", imp.Src, "parent_id", imp.Dest)
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}
func jsonDecodeTeamMemberProjectAssn(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the TeamMemberProjectAssn list to start with an array")
		return err
	}

	database := Database()
	for decoder.More() {
		var imp struct {
			Src  string `json:"project_id"`
			Dest string `json:"team_member_id"`
		}

		if err = decoder.Decode(&imp); err != nil {
			return err
		}
		db.Associate(ctx, database, "team_member_project_assn", "project_id", imp.Src, "team_member_id", imp.Dest)
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}
