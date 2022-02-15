package sqldb

import (
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	connStringParams = "parseTime=true"

	DriverMySQL = "mysql"
)

// SqlDB interface for db
type SqlDB interface {
	Connect(connStr string, maxConn int) error
	Connection() *sqlx.DB
	ConnectionString() string
}

// db implements the SqlDB interface
type db struct {
	dbconn
}

// dbconn contains one connection
type dbconn struct {
	connStr string
	conn    *sqlx.DB
}

func NewDB() SqlDB {
	return &db{}
}

func (d *db) Connect(connStr string, maxConn int) error {
	c, err := d.createConnection(connStr, connStringParams, maxConn)
	if err != nil {
		return err
	}
	d.dbconn = c
	return nil
}

func (d *db) Connection() *sqlx.DB {
	return d.conn
}

func (d *db) ConnectionString() string {
	return d.connStr
}

func (d *db) createConnection(connStr string, params string, maxConn int) (dbconn, error) {

	fullConnStr := connStr
	if strings.Contains(fullConnStr, "?") {
		if fullConnStr[len(fullConnStr)-1] != '?' {
			fullConnStr += "&"
		}
		fullConnStr += params
	} else {
		fullConnStr += "?" + params
	}

	conn, err := sqlx.Connect(DriverMySQL, fullConnStr)
	if err != nil {
		return dbconn{}, err
	}

	conn.SetMaxOpenConns(maxConn)
	conn.SetMaxIdleConns(0)

	return dbconn{
		conn:    conn,
		connStr: connStr,
	}, nil
}
