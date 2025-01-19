// Code generated by goradd-orm. DO NOT EDIT.

// Package goradd_unit contains the object relational model for the goradd_unit database.
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
package goradd_unit

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/goradd/orm/pkg/db"
)

// The newObjectPkCounter is used to assign a temporary primary key to an auto generated primary key that
// has not been saved. Auto generated primary keys are generated by the database driver, and until the record
// is saved, it does not have the final value. But having unique values for the primary key of new values helps
// with association tables and in user interfaces with associating rows of a table for example, with new objects
// that are yet to be saved. This does not need to be a thread safe object since we just need to guarantee uniqueness
// within a thread. Duplicates or skipped values between threads are OK.
var newObjectPkCounter int64

// Database returns the database object corresponding to goradd_unit in the global database cluster.
func Database() db.DatabaseI {
	return db.GetDatabase("goradd_unit")
}

// JsonEncodeAll sends the entire database to writer as JSON.
func JsonEncodeAll(ctx context.Context, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	if _, err := io.WriteString(writer, "[\n"); err != nil {
		return err
	}

	{ // Write DoubleIndices
		if _, err := io.WriteString(writer, "["); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, `"double_index"`); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return err
		}

		cursor := QueryDoubleIndices(ctx).LoadCursor()
		defer cursor.Close()
		if obj := cursor.Next(); obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		for obj := cursor.Next(); obj != nil; obj = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return err
			}
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return err
		}
	}
	{ // Write Reverses
		if _, err := io.WriteString(writer, "["); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, `"reverse"`); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return err
		}

		cursor := QueryReverses(ctx).LoadCursor()
		defer cursor.Close()
		if obj := cursor.Next(); obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		for obj := cursor.Next(); obj != nil; obj = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return err
			}
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return err
		}
	}
	{ // Write TypeTests
		if _, err := io.WriteString(writer, "["); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, `"type_test"`); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return err
		}

		cursor := QueryTypeTests(ctx).LoadCursor()
		defer cursor.Close()
		if obj := cursor.Next(); obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		for obj := cursor.Next(); obj != nil; obj = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return err
			}
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return err
		}
	}
	{ // Write UnsupportedTypes
		if _, err := io.WriteString(writer, "["); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, `"unsupported_type"`); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return err
		}

		cursor := QueryUnsupportedTypes(ctx).LoadCursor()
		defer cursor.Close()
		if obj := cursor.Next(); obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		for obj := cursor.Next(); obj != nil; obj = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return err
			}
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return err
		}
	}
	{ // Write ForwardCascades
		if _, err := io.WriteString(writer, "["); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, `"forward_cascade"`); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return err
		}

		cursor := QueryForwardCascades(ctx).LoadCursor()
		defer cursor.Close()
		if obj := cursor.Next(); obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		for obj := cursor.Next(); obj != nil; obj = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return err
			}
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return err
		}
	}
	{ // Write ForwardCascadeUniques
		if _, err := io.WriteString(writer, "["); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, `"forward_cascade_unique"`); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return err
		}

		cursor := QueryForwardCascadeUniques(ctx).LoadCursor()
		defer cursor.Close()
		if obj := cursor.Next(); obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		for obj := cursor.Next(); obj != nil; obj = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return err
			}
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return err
		}
	}
	{ // Write ForwardNulls
		if _, err := io.WriteString(writer, "["); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, `"forward_null"`); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return err
		}

		cursor := QueryForwardNulls(ctx).LoadCursor()
		defer cursor.Close()
		if obj := cursor.Next(); obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		for obj := cursor.Next(); obj != nil; obj = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return err
			}
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return err
		}
	}
	{ // Write ForwardNullUniques
		if _, err := io.WriteString(writer, "["); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, `"forward_null_unique"`); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return err
		}

		cursor := QueryForwardNullUniques(ctx).LoadCursor()
		defer cursor.Close()
		if obj := cursor.Next(); obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		for obj := cursor.Next(); obj != nil; obj = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return err
			}
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return err
		}
	}
	{ // Write ForwardRestricts
		if _, err := io.WriteString(writer, "["); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, `"forward_restrict"`); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return err
		}

		cursor := QueryForwardRestricts(ctx).LoadCursor()
		defer cursor.Close()
		if obj := cursor.Next(); obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		for obj := cursor.Next(); obj != nil; obj = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return err
			}
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, ","); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return err
		}
	}
	{ // Write ForwardRestrictUniques
		if _, err := io.WriteString(writer, "["); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, `"forward_restrict_unique"`); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return err
		}

		cursor := QueryForwardRestrictUniques(ctx).LoadCursor()
		defer cursor.Close()
		if obj := cursor.Next(); obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		for obj := cursor.Next(); obj != nil; obj = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return err
			}
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		if _, err := io.WriteString(writer, "]\n]"); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, "\n"); err != nil {
			return err
		}
	}

	_, err := io.WriteString(writer, "]")
	return err
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
		case "double_index":
			err = jsonDecodeDoubleIndices(ctx, decoder)
		case "forward_cascade":
			err = jsonDecodeForwardCascades(ctx, decoder)
		case "forward_cascade_unique":
			err = jsonDecodeForwardCascadeUniques(ctx, decoder)
		case "forward_null_unique":
			err = jsonDecodeForwardNullUniques(ctx, decoder)
		case "forward_restrict":
			err = jsonDecodeForwardRestricts(ctx, decoder)
		case "forward_restrict_unique":
			err = jsonDecodeForwardRestrictUniques(ctx, decoder)
		case "reverse":
			err = jsonDecodeReverses(ctx, decoder)
		case "type_test":
			err = jsonDecodeTypeTests(ctx, decoder)
		case "unsupported_type":
			err = jsonDecodeUnsupportedTypes(ctx, decoder)
		case "forward_null":
			err = jsonDecodeForwardNulls(ctx, decoder)

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

func jsonDecodeDoubleIndices(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the DoubleIndex list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewDoubleIndex()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		obj.Save(ctx)
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
func jsonDecodeForwardCascades(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the ForwardCascade list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewForwardCascade()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		obj.Save(ctx)
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
func jsonDecodeForwardCascadeUniques(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the ForwardCascadeUnique list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewForwardCascadeUnique()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		obj.Save(ctx)
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
func jsonDecodeForwardNullUniques(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the ForwardNullUnique list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewForwardNullUnique()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		obj.Save(ctx)
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
func jsonDecodeForwardRestricts(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the ForwardRestrict list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewForwardRestrict()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		obj.Save(ctx)
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
func jsonDecodeForwardRestrictUniques(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the ForwardRestrictUnique list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewForwardRestrictUnique()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		obj.Save(ctx)
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
func jsonDecodeReverses(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the Reverse list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewReverse()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		obj.Save(ctx)
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
func jsonDecodeTypeTests(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the TypeTest list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewTypeTest()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		obj.Save(ctx)
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
func jsonDecodeUnsupportedTypes(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the UnsupportedType list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewUnsupportedType()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		obj.Save(ctx)
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
func jsonDecodeForwardNulls(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the ForwardNull list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewForwardNull()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		obj.Save(ctx)
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
