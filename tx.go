package sqlxplus

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Tx struct {
	SqlxTx sqlxer
	Log    Logger
}

func (that *Tx) GetSqlxTx() *sqlx.Tx {
	return that.SqlxTx.(*sqlx.Tx)
}

func (that *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	that.Log.Printf(ctx, query, args...)
	rows, err = that.SqlxTx.QueryContext(ctx, query, args...)
	if err != nil {
		that.Log.Errorln(ctx, err)
	}
	return
}

func (that *Tx) QueryxContext(ctx context.Context, query string, args ...interface{}) (rows *sqlx.Rows, err error) {
	that.Log.Printf(ctx, query, args...)
	rows, err = that.SqlxTx.QueryxContext(ctx, query, args...)
	if err != nil {
		that.Log.Errorln(ctx, err)
	}
	return
}

func (that *Tx) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	that.Log.Printf(ctx, query, args...)
	return that.SqlxTx.QueryRowxContext(ctx, query, args...)
}

func (that *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (result sql.Result, err error) {
	that.Log.Printf(ctx, query, args...)
	result, err = that.SqlxTx.ExecContext(ctx, query, args...)
	if err != nil {
		that.Log.Errorln(ctx, err)
	}
	return
}

func (that *Tx) PrepareContext(ctx context.Context, query string) (stmt *sql.Stmt, err error) {
	that.Log.Println(ctx, query)
	stmt, err = that.SqlxTx.PrepareContext(ctx, query)
	if err != nil {
		that.Log.Errorln(ctx, err)
	}
	return
}

func (that *Tx) Query(query string, args ...interface{}) (rows *sql.Rows, err error) {
	that.Log.Printf(context.Background(), query, args...)
	rows, err = that.SqlxTx.Query(query, args...)
	if err != nil {
		that.Log.Errorln(context.Background(), err)
	}
	return
}

func (that *Tx) Queryx(query string, args ...interface{}) (rows *sqlx.Rows, err error) {
	that.Log.Printf(context.Background(), query, args...)
	rows, err = that.SqlxTx.Queryx(query, args...)
	if err != nil {
		that.Log.Errorln(context.Background(), err)
	}
	return
}

func (that *Tx) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	that.Log.Printf(context.Background(), query, args...)
	return that.SqlxTx.QueryRowx(query, args...)
}

func (that *Tx) Exec(query string, args ...interface{}) (result sql.Result, err error) {
	that.Log.Printf(context.Background(), query, args...)
	result, err = that.SqlxTx.Exec(query, args...)
	if err != nil {
		that.Log.Errorln(context.Background(), err)
	}
	return
}

func (that *Tx) Prepare(query string) (stmt *sql.Stmt, err error) {
	that.Log.Println(context.Background(), query)
	stmt, err = that.SqlxTx.Prepare(query)
	if err != nil {
		that.Log.Errorln(context.Background(), err)
	}
	return
}
