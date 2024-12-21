package mysql

import (
	"context"
	sqldb "database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/goradd/all"
	sql2 "github.com/goradd/orm/pkg/db/sql"
	. "github.com/goradd/orm/pkg/query"
	"strings"
	"time"
)

// DB is the goradd driver for mysql databases. It works through the excellent go-sql-driver driver,
// to supply functionality above go's built in driver. To use it, call NewDB, but afterwards,
// work through the DB parent interface so that the underlying database can be swapped out later if needed.
//
// Timezones
// Timezones are always tricky. Mysql has some interesting quirks:
//   - Datetime types are internally stored in the timezone of the server, and then returned based on the
//
// timezone of the client.
//   - Timestamp types are internally stored in UTC and returned in the timezone of the client.
//
// The benefit of this is that you can move your database
// to a server in another timezone, and the times will automatically change to the correct timezone.
//   - The mysql-go-driver has the ability to set a default timezone in the Loc configuration parameter
//
// It appears to convert all times to this timezone before sending them
// to the database, and then when receiving times, it will set this as the timezone of the date.
//
// These issues are further compounded by the fact that MYSQL can initialize date and time values to what it
// believes is the current date and time in its server's timezone, but will not save the timezone itself.
// If the database gets replicated around the world, you must explicitly set the timezone of each database
// master and slave to keep datetimes in sync. Also be aware that if you are using a scaling service that is global,
// it too may change the local timezone of the server, which may be different from the timezone of the database.
// Add to this the possibility that your users may be accessing the servers from different timezones than either the
// database or server, and you get quite a tangle.
//
// Add to that the TIMESTAMP has a max year of 2038, so TIMESTAMP itself is going to have to change soon.
//
// So, as a general rule, use DATETIME types to represent a date combined with a time, like an appointment in
// a calendar or a recurring event that happens is entered in the current timezone is and that is editable. If you
// change timezones, the time will change too.
// Use TIMESTAMP or DATETIME types to store data that records when an event happened in world time. Use separate DATE and TIME
// values to record a date and time that should always be thought of in the perspective of the viewer, and
// that if the viewer changes timezones, the time will not change. 9 am in one timezone is 9 am in the other(An alarm
// for example.)
//
// Also, set the Loc configuration parameter to be the same as the server's timezone. By default, its UTC.
// That will make it so all dates and times are in the same timezone as those automatically generated by MYSQL.
// It is best to set this and your database to UTC, as this will make your database portable to other timezones.
//
// Set the ParseTime configuration parameter to TRUE so that the driver will parse the times into the correct
// timezone, navigating the GO server and database server timezones. Otherwise, we
// can only assume that the database is in UTC time, since we will not get any timezone info from the server.
//
// The driver will return times in the timezone of the mysql server. This will mean that you can save data in local time,
// but you will need to convert to local time in some situations. Be particularly careful of DATE and TIME types, since
// these have no timezone information, and will always be in server time; converting to local time may have unintended
// effects.
//
// You need to be aware that when you view the data in the SQL, it will appear in whatever
// timezone the MYSQL server is set to.
type DB struct {
	sql2.DbHelper
	databaseName string
}

// NewDB returns a new DB database object that you can add to the datastore.
// If connectionString is set, it will be used to create the configuration. Otherwise,
// use a config setting.
//
// The postgres driver specifies that you must use ParseConfig
// to create the initial configuration, although that can be sent a blank string to
// gather initial values from environment variables. You can then change items in
// the configuration structure. For example:
//
//	config,_ := pgx.ParseConfig(connectionString)
//	config.Password = "mysecret"
//	db := pgsql.NewDB(key, "", config)
func NewDB(dbKey string, connectionString string, config *mysql.Config) *DB {
	if connectionString == "" && config == nil {
		panic("must specify how to connect to the database")
	}
	if connectionString == "" {
		connectionString = config.FormatDSN()
	}

	db3, err := sqldb.Open("mysql", connectionString)
	if err != nil {
		panic("Could not open database: " + err.Error())
	}
	err = db3.Ping()
	if err != nil {
		panic("Could not ping database " + dbKey + ":" + err.Error())
	}

	m := DB{
		DbHelper: sql2.NewSqlDb(dbKey, db3),
	}
	m.databaseName = config.DBName // save off the database name for later use
	return &m
}

// OverrideConfigSettings will use a map read in from a json file to modify
// the given config settings
func OverrideConfigSettings(config *mysql.Config, jsonContent map[string]interface{}) {
	for k, v := range jsonContent {
		switch k {
		case "database":
			config.DBName = v.(string)
		case "user":
			config.User = v.(string)
		case "password":
			config.Passwd = v.(string)
		case "net":
			config.Net = v.(string) // Typically, tcp or unix (for unix sockets).
		case "address":
			config.Addr = v.(string) // Note: if you set address, you MUST set net also.
		case "params":
			// Convert from map[string]any to map[string]string
			config.Params = all.StringMap(v.(map[string]interface{}))
		case "collation":
			config.Collation = v.(string)
		case "maxAllowedPacket":
			config.MaxAllowedPacket = int(v.(float64))
		case "serverPubKey":
			config.ServerPubKey = v.(string)
		case "tlsConfig":
			config.TLSConfig = v.(string)
		case "timeout":
			config.Timeout = time.Duration(int(v.(float64))) * time.Second
		case "readTimeout":
			config.ReadTimeout = time.Duration(int(v.(float64))) * time.Second
		case "writeTimeout":
			config.WriteTimeout = time.Duration(int(v.(float64))) * time.Second
		case "allowAllFiles":
			config.AllowAllFiles = v.(bool)
		case "allowCleartextPasswords":
			config.AllowCleartextPasswords = v.(bool)
		case "allowNativePasswords":
			config.AllowNativePasswords = v.(bool)
		case "allowOldPasswords":
			config.AllowOldPasswords = v.(bool)
		}
	}

	// The other config options effect how queries work, and so should be set before
	// calling this function, as they will change how the GO code for these queries will
	// need to be written.
}

// NewBuilder returns a new query builder to build a query that will be processed by the database.
func (m *DB) NewBuilder(ctx context.Context) QueryBuilderI {
	return sql2.NewSqlBuilder(ctx, m)
}

// iq surrounds the given value with sql identifier quotes.
func iq(v string) string {
	return string('`') + v + string('`')
}

// QuoteIdentifier surrounds the given identifier with quote characters
// appropriate for Postgres
func (m *DB) QuoteIdentifier(v string) string {
	return iq(v)
}

// FormatArgument formats the given argument number for embedding in a SQL statement.
// Mysql just uses a question mark as a placeholder.
func (m *DB) FormatArgument(n int) string {
	return "?"
}

// DeleteUsesAlias indicates the database requires the alias of a table after
// a delete clause when using aliases in the delete
func (m *DB) DeleteUsesAlias() bool {
	return true
}

// OperationSql provides Postgres specific SQL for certain operators.
func (m *DB) OperationSql(op Operator, operandStrings []string) (sql string) {
	switch op {
	case OpDateAddSeconds:
		// Modifying a datetime in the query
		// Only works on date, datetime and timestamps. Not times.
		s := operandStrings[0]
		s2 := operandStrings[1]
		sql = fmt.Sprintf(`DATE_ADD(%s, INTERVAL (%s) SECOND)`, s, s2)
	case OpXor:
		sOp := " " + op.String() + " "
		sql = " (" + strings.Join(operandStrings, sOp) + ") "
	}
	return
}

// Update sets specific fields of a record that already exists in the database to the given data.
func (m *DB) Update(ctx context.Context,
	table string,
	fields map[string]any,
	where map[string]any,
) {

	sql, args := sql2.GenerateUpdate(m, table, fields, where)
	_, e := m.Exec(ctx, sql, args...)
	if e != nil {
		panic(e.Error())
	}
}

// Insert inserts the given data as a new record in the database.
// It returns the record id of the new record.
func (m *DB) Insert(ctx context.Context, table string, fields map[string]interface{}) string {
	sql, args := sql2.GenerateInsert(m, table, fields)
	if r, err := m.Exec(ctx, sql, args...); err != nil {
		panic(err.Error())
	} else {
		if id, err2 := r.LastInsertId(); err2 != nil {
			panic(err2.Error())
			return ""
		} else {
			return fmt.Sprint(id)
		}
	}
}

// Delete deletes the indicated record from the database.
func (m *DB) Delete(ctx context.Context, table string, where map[string]any) {
	sql, args := sql2.GenerateDelete(m, table, where)
	_, e := m.Exec(ctx, sql, args...)
	if e != nil {
		panic(e.Error())
	}
}

// Associate sets up the many-many association pointing from the given table and column to another table and column.
// table is the name of the association table.
// column is the name of the column in the association table that contains the pk for the record we are associating.
// pk is the value of the primary key.
// relatedTable is the table the association is pointing to.
// relatedColumn is the column in the association table that points to the relatedTable's pk.
// relatedPks are the new primary keys in the relatedTable we are associating.
func (m *DB) Associate(ctx context.Context,
	table string,
	column string,
	pk interface{},
	_ string,
	relatedColumn string,
	relatedPks interface{}) {

	sql2.Associate(ctx, m, table, column, pk, relatedColumn, relatedPks)

}
