package sqlxplus

import (
	"context"
	"database/sql"
	"errors"

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

type CreateSqlxTranContextFunc func(ctx context.Context, tx *sqlx.Tx) context.Context
type TransactionFunc func(context.Context) error

type ExecTransactionOptions struct {
	DB                        *sqlx.DB
	Log                       Logger
	CreateSqlxTranContextFunc CreateSqlxTranContextFunc
}
type ExecTransactionOption func(o *ExecTransactionOptions)

func newExecTransactionOptions(opts ...ExecTransactionOption) ExecTransactionOptions {
	var o ExecTransactionOptions
	for _, opt := range opts {
		opt(&o)
	}
	if o.Log == nil {
		o.Log = &StdLogger{}
	}
	if o.CreateSqlxTranContextFunc == nil {
		o.CreateSqlxTranContextFunc = func(ctx context.Context, tx *sqlx.Tx) context.Context {
			return NewSqlxTranContext(ctx, &Tx{SqlxTx: tx, Log: o.Log})
		}
	}
	return o
}

func WithExecTransactionDB(db *sqlx.DB) ExecTransactionOption {
	return func(o *ExecTransactionOptions) {
		o.DB = db
	}
}

func WithExecTransactionLogger(log Logger) ExecTransactionOption {
	return func(o *ExecTransactionOptions) {
		o.Log = log
	}
}

func WithExecTransactionCreateSqlxTranContextFunc(f CreateSqlxTranContextFunc) ExecTransactionOption {
	return func(o *ExecTransactionOptions) {
		o.CreateSqlxTranContextFunc = f
	}
}

var ErrExecTransactionOptionDBIsNil = errors.New("db in options is nil")

func ExecTransaction(ctx context.Context, f TransactionFunc, opts ...ExecTransactionOption) (err error) {
	op := newExecTransactionOptions(opts...)
	var tran *sqlx.Tx
	// 判断上下文是否存在事务，如果不存在事务，则开启事务
	tranCreateFlag := false
	if !CheckSqlxTranContextExist(ctx) {
		// 判断上下文是否存在数据库，如果不存在数据库，则使用可选参数中的数据库
		var conn *DBConnect
		if conn, err = FromSqlxContext(ctx); err != nil {
			if errors.Is(err, ErrContextNotFoundSqlx) {
				err = nil
				if op.DB == nil {
					err = ErrExecTransactionOptionDBIsNil
					return
				}
				if tran, err = op.DB.Beginx(); err != nil {
					return
				}
			} else {
				return
			}
		} else {
			if tran, err = conn.DB().GetSqlxDB().Beginx(); err != nil {
				return
			}
		}
		defer func() {
			if reErr := recover(); reErr != nil {
				err = reErr.(error)
				tran.Rollback()
			} else if err != nil {
				tran.Rollback()
			}
		}()
		ctx = op.CreateSqlxTranContextFunc(ctx, tran)
		tranCreateFlag = true
	}
	if err = f(ctx); err != nil {
		return
	}
	if tranCreateFlag {
		err = tran.Commit()
	}
	return
}
