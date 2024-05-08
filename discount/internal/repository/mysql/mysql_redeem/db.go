package mysql_redeem

import "github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/repository/mysql"

type DB struct {
	conn *mysql.MySQLDB
}

func New(conn *mysql.MySQLDB) *DB {
	return &DB{
		conn: conn,
	}
}
