package sqlxplus

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type sqlxer interface {
	sqlx.QueryerContext
	sqlx.ExecerContext
	sqlx.PreparerContext
	sqlx.Queryer
	sqlx.Execer
	sqlx.Preparer
}

type DB struct {
	SqlxDB sqlxer
	Log    Logger
}

func (that *DB) GetSqlxDB() *sqlx.DB {
	return that.SqlxDB.(*sqlx.DB)
}

func (that *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	that.Log.Printf(ctx, query, args...)
	rows, err = that.SqlxDB.QueryContext(ctx, query, args...)
	if err != nil {
		that.Log.Errorln(ctx, err)
	}
	return
}

func (that *DB) QueryxContext(ctx context.Context, query string, args ...interface{}) (rows *sqlx.Rows, err error) {
	that.Log.Printf(ctx, query, args...)
	rows, err = that.SqlxDB.QueryxContext(ctx, query, args...)
	if err != nil {
		that.Log.Errorln(ctx, err)
	}
	return
}

func (that *DB) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	that.Log.Printf(ctx, query, args...)
	return that.SqlxDB.QueryRowxContext(ctx, query, args...)
}

func (that *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (result sql.Result, err error) {
	that.Log.Printf(ctx, query, args...)
	result, err = that.SqlxDB.ExecContext(ctx, query, args...)
	if err != nil {
		that.Log.Errorln(ctx, err)
	}
	return
}

func (that *DB) PrepareContext(ctx context.Context, query string) (stmt *sql.Stmt, err error) {
	that.Log.Println(ctx, query)
	stmt, err = that.SqlxDB.PrepareContext(ctx, query)
	if err != nil {
		that.Log.Errorln(ctx, err)
	}
	return
}

func (that *DB) Query(query string, args ...interface{}) (rows *sql.Rows, err error) {
	that.Log.Printf(context.Background(), query, args...)
	rows, err = that.SqlxDB.Query(query, args...)
	if err != nil {
		that.Log.Errorln(context.Background(), err)
	}
	return
}

func (that *DB) Queryx(query string, args ...interface{}) (rows *sqlx.Rows, err error) {
	that.Log.Printf(context.Background(), query, args...)
	rows, err = that.SqlxDB.Queryx(query, args...)
	if err != nil {
		that.Log.Errorln(context.Background(), err)
	}
	return
}

func (that *DB) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	that.Log.Printf(context.Background(), query, args...)
	return that.SqlxDB.QueryRowx(query, args...)
}

func (that *DB) Exec(query string, args ...interface{}) (result sql.Result, err error) {
	that.Log.Printf(context.Background(), query, args...)
	result, err = that.SqlxDB.Exec(query, args...)
	if err != nil {
		that.Log.Errorln(context.Background(), err)
	}
	return
}

func (that *DB) Prepare(query string) (stmt *sql.Stmt, err error) {
	that.Log.Println(context.Background(), query)
	stmt, err = that.SqlxDB.Prepare(query)
	if err != nil {
		that.Log.Errorln(context.Background(), err)
	}
	return
}

// ExecContext error handling of sqlx.MustExecContext.
func ExecContext(ctx context.Context, e sqlx.ExecerContext, query string, args ...interface{}) (sql.Result, error) {
	return e.ExecContext(ctx, query, args...)
}

// Exec error handling of sqlx.MustExec.
func Exec(e sqlx.Execer, query string, args ...interface{}) (sql.Result, error) {
	return e.Exec(query, args...)
}
