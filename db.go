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
	sqlxDB sqlxer
	log    Logger
}

func (that *DB) SqlxTx() *sqlx.DB {
	return that.sqlxDB.(*sqlx.DB)
}

func (that *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	that.log.Printf(ctx, query, args...)
	rows, err = that.sqlxDB.QueryContext(ctx, query, args...)
	if err != nil {
		that.log.Errorln(ctx, err)
	}
	return
}

func (that *DB) QueryxContext(ctx context.Context, query string, args ...interface{}) (rows *sqlx.Rows, err error) {
	that.log.Printf(ctx, query, args...)
	rows, err = that.sqlxDB.QueryxContext(ctx, query, args...)
	if err != nil {
		that.log.Errorln(ctx, err)
	}
	return
}

func (that *DB) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	that.log.Printf(ctx, query, args...)
	return that.sqlxDB.QueryRowxContext(ctx, query, args...)
}

func (that *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (result sql.Result, err error) {
	that.log.Printf(ctx, query, args...)
	result, err = that.sqlxDB.ExecContext(ctx, query, args...)
	if err != nil {
		that.log.Errorln(ctx, err)
	}
	return
}

func (that *DB) PrepareContext(ctx context.Context, query string) (stmt *sql.Stmt, err error) {
	that.log.Println(ctx, query)
	stmt, err = that.sqlxDB.PrepareContext(ctx, query)
	if err != nil {
		that.log.Errorln(ctx, err)
	}
	return
}

func (that *DB) Query(query string, args ...interface{}) (rows *sql.Rows, err error) {
	that.log.Printf(context.Background(), query, args...)
	rows, err = that.sqlxDB.Query(query, args...)
	if err != nil {
		that.log.Errorln(context.Background(), err)
	}
	return
}

func (that *DB) Queryx(query string, args ...interface{}) (rows *sqlx.Rows, err error) {
	that.log.Printf(context.Background(), query, args...)
	rows, err = that.sqlxDB.Queryx(query, args...)
	if err != nil {
		that.log.Errorln(context.Background(), err)
	}
	return
}

func (that *DB) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	that.log.Printf(context.Background(), query, args...)
	return that.sqlxDB.QueryRowx(query, args...)
}

func (that *DB) Exec(query string, args ...interface{}) (result sql.Result, err error) {
	that.log.Printf(context.Background(), query, args...)
	result, err = that.sqlxDB.Exec(query, args...)
	if err != nil {
		that.log.Errorln(context.Background(), err)
	}
	return
}

func (that *DB) Prepare(query string) (stmt *sql.Stmt, err error) {
	that.log.Println(context.Background(), query)
	stmt, err = that.sqlxDB.Prepare(query)
	if err != nil {
		that.log.Errorln(context.Background(), err)
	}
	return
}

// ExecContext error handling of sqlx.MustExecContext
func ExecContext(ctx context.Context, e sqlx.ExecerContext, query string, args ...interface{}) (sql.Result, error) {
	return e.ExecContext(ctx, query, args...)
}
