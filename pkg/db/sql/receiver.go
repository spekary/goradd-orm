package sql

import (
	"fmt"
	"github.com/goradd/orm/pkg/db"
	. "github.com/goradd/orm/pkg/query"
	"github.com/goradd/orm/pkg/schema"
	strings2 "github.com/goradd/strings"
	"log"
	"log/slog"
	"strconv"
	"strings"
	"time"
)

// SqlReceiver is an encapsulation of a way of receiving data from sql queries as interface{} pointers. This allows you
// to get data without knowing the type of data you are asking for ahead of time, and is easier for dealing with NULL fields.
// Some database drivers (MySql for one) return different results in fields depending on how you call the query (using
// a prepared statement can return different results than without one), or if the data does not quite fit (UInt64 in particular
// will return a string if the returned value is bigger than MaxInt64, but smaller than MaxUint64.)
//
// Pass the address of the R member to the sql.Scan method when using an object of this type,
// because there are some idiosyncrasies with
// how Go treats return values that prevents returning an address of R from a function
type SqlReceiver struct {
	R interface{}
}

// IntI returns the receiver as an interface to an int.
func (r SqlReceiver) IntI() interface{} {
	if r.R == nil {
		return nil
	}
	switch v := r.R.(type) {
	case int64:
		return int(v)
	case int:
		return r.R
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			// Note that some databases allow formulas for default integers, and we do not want to panic for those.
			i = 0
		}
		return i
	case []byte:
		v2 := string(v[:])
		if v2 == "NULL" {
			return nil
		} // MariaDB does this for default values
		i, err := strconv.Atoi(v2)
		if err != nil {
			slog.Error("Attempting to scan a non-integer column into an integer",
				slog.String(db.LogComponent, "SQL receiver"),
				slog.Any(db.LogError, err))
			panic(err)
		}
		return i

	default:
		slog.Error("Unknown type returned from sql driver",
			slog.String(db.LogComponent, "SQL receiver"))
		return nil
	}
}

// UintI converts the value to an interface to a GO uint.
//
// Some drivers (like MySQL) return all integers as Int64. Its up to you to make sure
// you only use this on 32-bit uints or smaller.
func (r SqlReceiver) UintI() interface{} {
	if r.R == nil {
		return nil
	}
	switch r.R.(type) {
	case int64:
		return uint(r.R.(int64))
	case int:
		return uint(r.R.(int))
	case uint:
		return r.R
	case string:
		i, err := strconv.ParseUint(r.R.(string), 10, 32)
		if err != nil {
			// Note that some databases allow formulas for default integers, and we do not want to panic for those.
			i = 0
		}
		return uint(i)
	case []byte:
		v := string(r.R.([]byte)[:])
		if v == "NULL" {
			return nil
		} // MariaDB does this for default values
		i, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			slog.Error("Attempting to scan a non-integer column into an unisgned integer",
				slog.String(db.LogComponent, "SQL receiver"),
				slog.Any(db.LogError, err))
		}
		return uint(i)
	default:
		slog.Error("Unknown type returned from sql driver",
			slog.String(db.LogComponent, "SQL receiver"))
		return nil
	}
}

// Int64I returns the given value as an interface to an Int64
func (r SqlReceiver) Int64I() interface{} {
	if r.R == nil {
		return nil
	}
	switch r.R.(type) {
	case int64:
		return r.R
	case int:
		return int64(r.R.(int))
	case string:
		i, err := strconv.ParseInt(r.R.(string), 10, 64)
		if err != nil {
			// Note that some databases allow formulas for default integers, and we do not want to panic for those.
			i = 0
		}
		return i
	case []byte:
		v := string(r.R.([]byte)[:])
		if v == "NULL" {
			return nil
		} // MariaDB does this for default values
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			slog.Error("Attempting to scan a non-integer column into an int64",
				slog.String(db.LogComponent, "SQL receiver"),
				slog.Any(db.LogError, err))
		}
		return i
	default:
		slog.Error("Unknown type returned from sql driver",
			slog.String(db.LogComponent, "SQL receiver"))
		return nil
	}
}

// Uint64I returns a value as an interface to a UInt64.
//
// Some drivers (like MySQL) return all integers as Int64. This converts to uint64. Its up to you to make sure
// you only use this on 64-bit uints or smaller.
func (r SqlReceiver) Uint64I() interface{} {
	if r.R == nil {
		return nil
	}
	switch r.R.(type) {
	case int64:
		return uint64(r.R.(int64))
	case int:
		return uint64(r.R.(int))
	case uint64:
		return r.R
	case string: // Mysql returns this if the detected value is greater than int64 size
		i, err := strconv.ParseUint(r.R.(string), 10, 64)
		if err != nil {
			// Note that some databases allow formulas for default integers, and we do not want to panic for those.
			i = 0
		}
		return i
	case []byte:
		v := string(r.R.([]byte)[:])
		if v == "NULL" {
			return nil
		} // MariaDB does this for default values
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			slog.Error("Attempting to scan a non-integer column into a uint64",
				slog.String(db.LogComponent, "SQL receiver"),
				slog.Any(db.LogError, err))
		}
		return i
	default:
		slog.Error("Unknown type returned from sql driver",
			slog.String(db.LogComponent, "SQL receiver"))
		return nil
	}
}

// BoolI returns the value as an interface to a boolean
func (r SqlReceiver) BoolI() interface{} {
	if r.R == nil {
		return nil
	}
	switch v := r.R.(type) {
	case bool:
		return r.R
	case int:
		return v != 0
	case int64:
		return v != 0
	case string:
		b, err := strconv.ParseBool(r.R.(string))
		if err != nil {
			b = false
			if strings2.EqualCaseInsensitive(v, "YES") {
				b = true
			}
			if strings2.EqualCaseInsensitive(v, "TRUE") {
				b = true
			}
		}
		return b
	case []byte:
		v2 := string(v[:])
		if v2 == "NULL" {
			return nil
		} // MariaDB does this for default values
		b, err := strconv.ParseBool(v2)
		if err != nil {
			slog.Error("Attempting to scan a non-boolean into a bool",
				slog.String(db.LogComponent, "SQL receiver"),
				slog.Any(db.LogError, err))
		}
		return b

	default:
		slog.Error("Unknown type returned from sql driver",
			slog.String(db.LogComponent, "SQL receiver"))
		return nil
	}
}

// StringI returns the value as an interface to a string
func (r SqlReceiver) StringI() interface{} {
	if r.R == nil {
		return nil
	}
	switch r.R.(type) {
	case string:
		return r.R
	case []byte:
		v := string(r.R.([]byte)[:])
		return v
	default:
		return fmt.Sprint(r.R)
	}
}

// ByteI returns the value as an interface to a byte array
func (r SqlReceiver) ByteI() interface{} {
	if r.R == nil {
		return nil
	}
	switch v := r.R.(type) {
	case string:
		return []byte(v)
	case []byte:
		return v
	default:
		return r.R // let it pass through driver
	}
}

// FloatI returns the value as an interface to a float32 value.
func (r SqlReceiver) FloatI() interface{} {
	if r.R == nil {
		return nil
	}
	switch r.R.(type) {
	case float32:
		return r.R
	case float64:
		return float32(r.R.(float64))
	case string:
		f, err := strconv.ParseFloat(r.R.(string), 32)
		if err != nil {
			slog.Error("Attempting to scan a non-float into a float",
				slog.String(db.LogComponent, "SQL receiver"),
				slog.Any(db.LogError, err))
		}
		return f
	case []byte:
		v := string(r.R.([]byte)[:])
		if v == "NULL" {
			return nil
		} // MariaDB does this for default values
		f, err := strconv.ParseFloat(v, 32)
		if err != nil {
			slog.Error("Attempting to scan a non-float into a float",
				slog.String(db.LogComponent, "SQL receiver"),
				slog.Any(db.LogError, err))
		}
		return float32(f)
	default:
		slog.Error("Unknown type returned from sql driver",
			slog.String(db.LogComponent, "SQL receiver"))
		return nil
	}
}

// DoubleI returns the value as a float64 interface
func (r SqlReceiver) DoubleI() interface{} {
	if r.R == nil {
		return nil
	}
	switch r.R.(type) {
	case float32:
		return float64(r.R.(float32))
	case float64:
		return r.R
	case string:
		f, err := strconv.ParseFloat(r.R.(string), 64)
		if err != nil {
			log.Panic(err)
		}
		return f
	case []byte:
		v := string(r.R.([]byte)[:])
		if v == "NULL" {
			return nil
		} // MariaDB does this for default values
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			slog.Error("Attempting to scan a non-float into a float",
				slog.String(db.LogComponent, "SQL receiver"),
				slog.Any(db.LogError, err))
		}
		return f
	default:
		slog.Error("Unknown type returned from sql driver",
			slog.String(db.LogComponent, "SQL receiver"))
		return nil
	}
}

// TimeI returns the value as a time.Time value in UTC, or in the case of CURRENT_TIME, a string "now".
func (r SqlReceiver) TimeI() interface{} {
	if r.R == nil {
		return nil
	}

	switch v := r.R.(type) {
	case time.Time:
		return v
	case string:
		return parseTime(v) // Note that this must always include timezone information if coming from a timestamp with timezone column
	case []byte:
		return parseTime(string(v))
		// TODO: SQL Lite, may return an int or float. Not sure we can support these.
	default:
		slog.Error("Unknown type returned from sql driver",
			slog.String(db.LogComponent, "SQL receiver"))
		return nil
	}
}

func parseTime(s string) interface{} {
	if s == "NULL" {
		return nil
	}
	u := strings.ToUpper(s)
	if strings2.StartsWith(u, "CURRENT_TIMESTAMP") {
		return "now"
	}
	t := ParseTime(s)
	if t.IsZero() {
		return t
	}
	return t.UTC()
}

// Unpack converts a SqlReceiver to a type corresponding to the given ReceiverType
func (r SqlReceiver) Unpack(typ ReceiverType) interface{} {
	switch typ {
	case ColTypeBytes:
		return r.R
	case ColTypeString:
		return r.StringI()
	case ColTypeInteger:
		return r.IntI()
	case ColTypeUnsigned:
		return r.UintI()
	case ColTypeInteger64:
		return r.Int64I()
	case ColTypeUnsigned64:
		return r.Uint64I()
	case ColTypeTime:
		return r.TimeI()
	case ColTypeFloat32:
		return r.FloatI()
	case ColTypeFloat64:
		return r.DoubleI()
	case ColTypeBool:
		return r.BoolI()
	default:
		return r.R
	}
}

// UnpackDefaultValue converts a SqlReceiver used to get the default value
// to a type corresponding to typ.
func (r SqlReceiver) UnpackDefaultValue(typ schema.ColumnType, size int) interface{} {
	switch typ {
	case schema.ColTypeBytes:
		fallthrough
	case schema.ColTypeUnknown:
		b := r.ByteI()
		if b == nil {
			return b
		}
		if string(b.([]byte)) == "NULL" {
			return nil
		}
		return b
	case schema.ColTypeString:
		s := r.StringI()
		if s == nil {
			return s
		}
		s2 := s.(string)
		if s2 == "NULL" {
			return nil
		}
		// Unwrap single quotes coming from mariadb and postgres
		s = strings2.Between(s2, `'`, `'`)
		return s
	case schema.ColTypeEnum:
		fallthrough
	case schema.ColTypeInt:
		if size == 64 {
			return r.Int64I()
		}
		return r.IntI()
	case schema.ColTypeUint:
		if size == 64 {
			return r.Uint64I()
		}
		return r.UintI()
	case schema.ColTypeTime:
		return r.TimeI()
	case schema.ColTypeFloat:
		if size == 32 {
			return r.FloatI()
		}
		return r.DoubleI()
	case schema.ColTypeBool:
		return r.BoolI()
	default:
		return r.R
	}
}
