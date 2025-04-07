package pgsql

import (
	"context"
	sqldb "database/sql"
	"fmt"
	"github.com/goradd/all"
	"github.com/goradd/orm/pkg/db"
	sql2 "github.com/goradd/orm/pkg/db/sql"
	. "github.com/goradd/orm/pkg/query"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/stdlib"
	"strings"
)

// DB is the goradd driver for postgresql databases.
type DB struct {
	sql2.DbHelper
}

// NewDB returns a new Postgresql DB database object based on the pgx driver
// that you can add to the datastore.
// If connectionString is set, it will be used to create the configuration. Otherwise,
// use a config setting. Using a configSetting can potentially give you access to the
// underlying pgx database for advanced operations.
//
// The postgres driver specifies that you must use ParseConfig
// to create the initial configuration, although that can be sent a blank string to
// gather initial values from environment variables. You can then change items in
// the configuration structure. For example:
//
//	config,_ := pgx.ParseConfig(connectionString)
//	config.Password = "mysecret"
//	db := pgsql.NewDB(key, "", config)
func NewDB(dbKey string,
	connectionString string,
	config *pgx.ConnConfig) (*DB, error) {
	if connectionString == "" && config == nil {
		return nil, fmt.Errorf("must specify how to connect to the database")
	}

	if connectionString == "" {
		connectionString = stdlib.RegisterConnConfig(config)
	}

	db3, err := sqldb.Open("pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}
	err = db3.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not ping database: %w", err)
	}

	m := new(DB)
	m.DbHelper = sql2.NewSqlHelper(dbKey, db3, m)
	return m, nil
}

// OverrideConfigSettings will use a map read in from a json file to modify
// the given config settings
func OverrideConfigSettings(config *pgx.ConnConfig, jsonContent map[string]interface{}) {
	for k, v := range jsonContent {
		switch k {
		case "database":
			config.Database = v.(string)
		case "user":
			config.User = v.(string)
		case "password":
			config.Password = v.(string)
		case "host":
			config.Host = v.(string) // Typically, tcp or unix (for unix sockets).
		case "port":
			config.Port = uint16(v.(float64))
		case "runtimeParams":
			config.RuntimeParams = all.StringMap(v.(map[string]interface{}))
		case "kerberosServerName":
			config.KerberosSrvName = v.(string)
		case "kerberosSPN":
			config.KerberosSpn = v.(string)
		}
	}
}

// QuoteIdentifier surrounds the given identifier with quote characters
// appropriate for Postgres
func (m *DB) QuoteIdentifier(v string) string {
	var b strings.Builder
	b.WriteRune('"')

	if i := strings.Index(v, "."); i != -1 {
		b.WriteString(v[:i])
		b.WriteString(`"."`)
		b.WriteString(v[i+1:])
	} else {
		b.WriteString(v)
	}

	b.WriteRune('"')
	return b.String()
}

// FormatArgument formats the given argument number for embedding in a SQL statement.
func (m *DB) FormatArgument(n int) string {
	return fmt.Sprintf(`$%d`, n)
}

// OperationSql provides Postgres specific SQL for certain operators.
func (m *DB) OperationSql(op Operator, operandStrings []string) (sql string) {
	switch op {
	case OpDateAddSeconds:
		// Modifying a datetime in the query
		// Only works on date, datetime and timestamps. Not times.
		s := operandStrings[0]
		s2 := operandStrings[1]
		return fmt.Sprintf(`(%s + make_interval(seconds => %s))`, s, s2)
	}
	return
}

// Insert inserts the given data as a new record in the database.
// It returns the record id of the new record, and possibly an error if an error occurred.
func (m *DB) Insert(ctx context.Context, table string, pkName string, fields map[string]interface{}) (string, error) {
	sql, args := sql2.GenerateInsert(m, table, fields)
	if pkName != "" {
		sql += " RETURNING "
		sql += pkName
	}
	if rows, err := m.SqlQuery(ctx, sql, args...); err != nil {
		if pgErr, ok := all.As[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" {
				return "", db.NewUniqueValueError(table, "", "", err)
			}
		}
		return "", db.NewQueryError("SqlQuery", sql, args, err)
	} else {
		var id string
		defer sql2.RowClose(rows)
		for rows.Next() {
			err = rows.Scan(&id)
			return "", db.NewQueryError("Scan", sql, args, err)
		}
		if err = rows.Err(); err != nil {
			return "", db.NewQueryError("rows.Err", sql, args, err)
		} else {
			return id, nil
		}
	}
}

// Update sets specific fields of a record that already exists in the database.
// optLockFieldName is the name of a version field that will implement an optimistic locking check before executing the update.
// Note that if the database is not currently in a transaction, then the optimistic lock
// will be a cursory check, but will not be able to definitively prevent a prior write.
// If optLockFieldName is provided, that field will be updated with a new version number value during the update process.
func (m *DB) Update(ctx context.Context,
	table string,
	pkName string,
	pkValue any,
	fields map[string]any,
	optLockFieldName string,
	optLockFieldValue int64,
) (newLock int64, err error) {
	if len(fields) == 0 {
		panic("fields must not be empty")
	}
	newLock, err = m.DbHelper.CheckLock(ctx, table, pkName, pkValue, optLockFieldName, optLockFieldValue)
	if err != nil {
		return
	}
	if newLock != 0 {
		fields[optLockFieldName] = newLock
	}
	s, args := sql2.GenerateUpdate(m, table, fields, map[string]any{pkName: pkValue})
	_, err = m.SqlExec(ctx, s, args...)
	if err != nil {
		if pgErr, ok := all.As[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" {
				return 0, db.NewUniqueValueError(table, "", "", err)
			}
		}
		return 0, db.NewQueryError("SqlExec", s, args, err)
	}
	return
}

func (m *DB) SupportsForUpdate() bool {
	return true
}
