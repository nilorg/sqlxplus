package sqlxplus

import (
	"context"
	"errors"
)

var (
	// ErrContextNotFoundSqlx 上下文不存在Sqlx错误
	ErrContextNotFoundSqlx = errors.New("上下文中没有获取到Sqlx")
)

type SqlxKey struct{}
type SqlxTranKey struct{}

// FromSqlxContext 从上下文中获取Sqlx
func FromSqlxContext(ctx context.Context) (conn *DBConnect, err error) {
	var (
		db *DB
		tx *Tx
		ok bool
	)
	tx, ok = ctx.Value(SqlxTranKey{}).(*Tx)
	if ok {
		conn = &DBConnect{
			transaction: true,
			conn:        tx,
		}
	} else {
		db, ok = ctx.Value(SqlxKey{}).(*DB)
		if ok {
			conn = &DBConnect{
				transaction: false,
				conn:        db,
			}
		} else {
			err = ErrContextNotFoundSqlx
		}
	}
	return
}

// CheckSqlxTranContextExist 检查sqlx是否存在
func CheckSqlxTranContextExist(ctx context.Context) bool {
	_, ok := ctx.Value(SqlxTranKey{}).(*Tx)
	return ok
}

// NewSqlxContext 创建Sqlx上下文
func NewSqlxContext(ctx context.Context, xdb *DB) context.Context {
	return context.WithValue(ctx, SqlxKey{}, xdb)
}

// NewSqlxTranContext 创建sqlx事务上下文
func NewSqlxTranContext(ctx context.Context, xtx *Tx) context.Context {
	return context.WithValue(ctx, SqlxTranKey{}, xtx)
}
