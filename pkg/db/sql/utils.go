package sql

// Helper utilities for extracting a description out of a database

import (
	"encoding/json"
	"strconv"
	"strings"
)

// ExtractOptions finds the json encoded list of options in the given string
func ExtractOptions(comment string) (options map[string]interface{}, remainingComment string, err error) {
	var optionString string
	firstIndex := strings.Index(comment, "{")
	lastIndex := strings.LastIndex(comment, "}")
	options = make(map[string]interface{})

	if firstIndex != -1 &&
		lastIndex != -1 &&
		lastIndex > firstIndex {

		optionString = comment[firstIndex : lastIndex+1]
		remainingComment = strings.TrimSpace(comment[:firstIndex] + comment[lastIndex+1:])

		err = json.Unmarshal([]byte(optionString), &options)
	}
	return
}

// GetDataDefLength will extract the length from the definition given a data definition description of the table.
// Example:
//
//	bigint(21) -> 21
//
// varchar(50) -> 50
// decimal(10,2) -> 10
func GetDataDefLength(description string) (l int, subLen int) {
	var lastPos, lenPos int
	var size string
	if lenPos = strings.Index(description, "("); lenPos != -1 {
		lastPos = strings.LastIndex(description, ")")
		size = description[lenPos+1 : lastPos]
		sizes := strings.Split(size, ",")
		l, _ = strconv.Atoi(sizes[0])
		if len(sizes) > 1 {
			subLen, _ = strconv.Atoi(sizes[1])
		}
		return
	}
	return
}

// Retrieves a numeric value from the options, which is always going to return a float64
func getNumericOption(o map[string]interface{}, option string, defaultValue float64) (float64, bool) {
	if v := o[option]; v != nil {
		if v2, ok := v.(float64); !ok {
			return defaultValue, false
		} else {
			return v2, true
		}
	} else {
		return defaultValue, true
	}
}

// Retrieves a boolean value from the options
/*
func getBooleanOption(o *maps.SliceMap, option string) (val bool, ok bool) {
	val, ok = o.LoadBool(option)
	return
}
*/
