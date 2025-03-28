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
	{ // Write Leafs
		if _, err := io.WriteString(writer, "["); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, `"leaf"`); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return err
		}

		cursor := QueryLeafs(ctx).LoadCursor()
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
	{ // Write LeafLocks
		if _, err := io.WriteString(writer, "["); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, `"leaf_lock"`); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return err
		}

		cursor := QueryLeafLocks(ctx).LoadCursor()
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
	{ // Write Roots
		if _, err := io.WriteString(writer, "["); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, `"root"`); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, ",\n["); err != nil {
			return err
		}

		cursor := QueryRoots(ctx).LoadCursor()
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
		case "leaf":
			err = jsonDecodeLeafs(ctx, decoder)
		case "leaf_lock":
			err = jsonDecodeLeafLocks(ctx, decoder)
		case "root":
			err = jsonDecodeRoots(ctx, decoder)
		case "type_test":
			err = jsonDecodeTypeTests(ctx, decoder)
		case "unsupported_type":
			err = jsonDecodeUnsupportedTypes(ctx, decoder)

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
func jsonDecodeLeafs(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the Leaf list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewLeaf()
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
func jsonDecodeLeafLocks(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the LeafLock list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewLeafLock()
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
func jsonDecodeRoots(ctx context.Context, decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the Root list to start with an array")
		return err
	}

	for decoder.More() {
		obj := NewRoot()
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
