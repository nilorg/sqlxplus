package sqlxplus

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Tx struct {
	sqlxTx sqlxer
	log    Logger
}

func (that *Tx) SqlxTx() *sqlx.Tx {
	return that.sqlxTx.(*sqlx.Tx)
}

func (that *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	that.log.Printf(ctx, query, args...)
	rows, err = that.sqlxTx.QueryContext(ctx, query, args...)
	if err != nil {
		that.log.Errorln(ctx, err)
	}
	return
}

func (that *Tx) QueryxContext(ctx context.Context, query string, args ...interface{}) (rows *sqlx.Rows, err error) {
	that.log.Printf(ctx, query, args...)
	rows, err = that.sqlxTx.QueryxContext(ctx, query, args...)
	if err != nil {
		that.log.Errorln(ctx, err)
	}
	return
}

func (that *Tx) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	that.log.Printf(ctx, query, args...)
	return that.sqlxTx.QueryRowxContext(ctx, query, args...)
}

func (that *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (result sql.Result, err error) {
	that.log.Printf(ctx, query, args...)
	result, err = that.sqlxTx.ExecContext(ctx, query, args...)
	if err != nil {
		that.log.Errorln(ctx, err)
	}
	return
}

func (that *Tx) PrepareContext(ctx context.Context, query string) (stmt *sql.Stmt, err error) {
	that.log.Println(ctx, query)
	stmt, err = that.sqlxTx.PrepareContext(ctx, query)
	if err != nil {
		that.log.Errorln(ctx, err)
	}
	return
}

func (that *Tx) Query(query string, args ...interface{}) (rows *sql.Rows, err error) {
	that.log.Printf(context.Background(), query, args...)
	rows, err = that.sqlxTx.Query(query, args...)
	if err != nil {
		that.log.Errorln(context.Background(), err)
	}
	return
}

func (that *Tx) Queryx(query string, args ...interface{}) (rows *sqlx.Rows, err error) {
	that.log.Printf(context.Background(), query, args...)
	rows, err = that.sqlxTx.Queryx(query, args...)
	if err != nil {
		that.log.Errorln(context.Background(), err)
	}
	return
}

func (that *Tx) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	that.log.Printf(context.Background(), query, args...)
	return that.sqlxTx.QueryRowx(query, args...)
}

func (that *Tx) Exec(query string, args ...interface{}) (result sql.Result, err error) {
	that.log.Printf(context.Background(), query, args...)
	result, err = that.sqlxTx.Exec(query, args...)
	if err != nil {
		that.log.Errorln(context.Background(), err)
	}
	return
}

func (that *Tx) Prepare(query string) (stmt *sql.Stmt, err error) {
	that.log.Println(context.Background(), query)
	stmt, err = that.sqlxTx.Prepare(query)
	if err != nil {
		that.log.Errorln(context.Background(), err)
	}
	return
}
