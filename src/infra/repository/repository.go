package repository

import (
	"database/sql"
	"fmt"
	"os"
	"projeto-docker/src/core/domain/errors"
	"projeto-docker/src/infra"
	"projeto-docker/src/utils"
	"regexp"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"

	"github.com/jmoiron/sqlx"
)

var logger = infra.Logger()
var keyConstraintCompiler = regexp.MustCompile(`^.+?_(.*)_key`)
var primaryKeyConstraintCompiler = regexp.MustCompile(`"(\w+)_pkey?`)
var foreignKeyConstraintCompiler = regexp.MustCompile(`(\w+)_fk`)

type SQLTransaction struct {
	Query     func(query string, args ...interface{}) (*sql.Rows, errors.Error)
	QueryRow  func(query string, args ...interface{}) *sql.Row
	ExecQuery func(query string, args ...interface{}) (sql.Result, errors.Error)
	Rollback  func() errors.Error
	CloseConn func() error
	Commit    func() errors.Error
}

func init() {
	logger = logger.With().Str("schema", getDatabaseSchema()).Logger()
}

func getDatabaseSchema() string {
	return utils.GetenvWithDefault("DATABASE_SCHEMA", "postgres")
}

func getDatabaseURI() string {
	schema := getDatabaseSchema()
	user := os.Getenv("POSTGRES_USER")
	pwd := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("POSTGRES_NAME")
	authentication := fmt.Sprintf("%s:%s", user, pwd)
	dst := fmt.Sprintf("%s:%s/%s", host, port, name)
	sslMode := utils.GetenvWithDefault("DATABASE_SSL_MODE", "disable")
	return fmt.Sprintf("%s://%s@%s?sslmode=%s", schema, authentication, dst, sslMode)
}

func getConnection() (*sqlx.DB, errors.Error) {
	schema := getDatabaseSchema()
	connection, err := sqlx.Open(schema, getDatabaseURI())
	if err != nil {
		logger.Error().Msg("Error while acessing database: " + err.Error())
		return nil, TranslateError(err)
	}
	connection.SetConnMaxLifetime(time.Minute * 3)
	connection.SetMaxOpenConns(10)
	connection.SetMaxIdleConns(10)
	return connection, nil
}

func closeConnection(conn *sqlx.DB) error {
	err := conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func internalRollback(tx *sql.Tx) errors.Error {
	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		return TranslateError(rollbackErr)
	}
	return nil
}

func internalCommit(tx *sql.Tx) errors.Error {
	err := tx.Commit()
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return TranslateError(rbErr)
		}
		return TranslateError(err)
	}
	return nil
}

func Queryx(sqlQuery string, args ...interface{}) (*sqlx.Rows, errors.Error) {
	conn, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer closeConnection(conn)
	rows, queryErr := conn.Queryx(sqlQuery, args...)
	if queryErr != nil {
		return nil, TranslateError(queryErr)
	}
	return rows, nil
}

func ExecQuery(sqlQuery string, args ...interface{}) (sql.Result, errors.Error) {
	conn, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer closeConnection(conn)
	result, queryErr := conn.Exec(sqlQuery, args...)
	if queryErr != nil {
		return nil, TranslateError(queryErr)
	}
	return result, nil
}

func BeginTransaction() (*SQLTransaction, errors.Error) {
	conn, err := getConnection()
	if err != nil {
		return nil, err
	}
	tx, beginErr := conn.Begin()
	if beginErr != nil {
		return nil, TranslateError(beginErr)
	}
	queryRow := func(query string, args ...interface{}) *sql.Row {
		return tx.QueryRow(query, args...)
	}
	execQuery := func(query string, args ...interface{}) (sql.Result, errors.Error) {
		result, err := tx.Exec(query, args...)
		return result, TranslateError(err)
	}
	query := func(query string, args ...interface{}) (*sql.Rows, errors.Error) {
		rows, err := tx.Query(query, args...)
		return rows, TranslateError(err)
	}
	closeConnection := func() error {
		return closeConnection(conn)
	}
	rollBack := func() errors.Error {
		defer closeConnection()
		return internalRollback(tx)
	}
	commit := func() errors.Error {
		defer closeConnection()
		return internalCommit(tx)
	}
	transaction := &SQLTransaction{
		Query:     query,
		QueryRow:  queryRow,
		ExecQuery: execQuery,
		Rollback:  rollBack,
		CloseConn: closeConnection,
		Commit:    commit,
	}
	return transaction, nil
}

func TranslateError(err error) errors.Error {
	if err == nil {
		return nil
	}
	logger.Error().Caller(zerolog.CallerSkipFrameCount).Msg(err.Error())
	if strings.Contains(err.Error(), "duplicate key") {
		keyMatch := keyConstraintCompiler.FindStringSubmatch(err.Error())
		var key = ""
		if keyMatch == nil {
			pkeyMatch := primaryKeyConstraintCompiler.FindStringSubmatch(err.Error())
			if pkeyMatch == nil {
				return errors.NewInternal(err)
			}
			key = pkeyMatch[1]
		} else {
			key = keyMatch[1]
		}
		return errors.NewFromString(fmt.Sprintf("the value provided for the \"%s\" field is already in use", key))
	} else if strings.Contains(err.Error(), "violates foreign key constraint") {
		keyMatch := foreignKeyConstraintCompiler.FindStringSubmatch(err.Error())
		if keyMatch == nil {
			return errors.NewInternal(err)
		}
		key := keyMatch[1]
		return errors.NewFromString(fmt.Sprintf("the value provided for the \"%s\" key doesn't exists", key))
	} else if err == sql.ErrNoRows {
		return errors.New(err)
	}
	return errors.NewUnexpected()
}
