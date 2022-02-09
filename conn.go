package sqlxplus

type DBConnect struct {
	transaction bool
	conn        interface{}
}

func (db *DBConnect) IsTransaction() bool {
	return db.transaction
}

func (db *DBConnect) Transaction() *Tx {
	return db.conn.(*Tx)
}

func (db *DBConnect) DB() *DB {
	return db.conn.(*DB)
}
