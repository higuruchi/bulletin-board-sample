package database

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/higuruchi/bulletin-board-sample.git/internal/db_handler/mysql_gateway"
)

type mysql struct {
	Conn *sql.DB
}

type TableRow struct {
	Rows *sql.Rows
}

type SQLResult struct {
	Result sql.Result
}

func NewMySQL(user, password, ip, port, name string) (mysql_gateway.MySQLHandler, func(), error) {
	conn, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@(%s:%s)/%s",
		user,
		password,
		ip,
		port,
		name,
	))
	if err != nil {
		log.Fatal("database connection error: ", err)
	}
	
	err = conn.Ping()
	if err!= nil {
		log.Fatal("database ping error: ", err)
	}
	return &mysql{
		Conn: conn,
	}, func() { conn.Close() }, nil
}

func (handler *mysql) Query(
	statement string,
	args ...interface{},
) (mysql_gateway.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return nil, fmt.Errorf("calling handler.Conn.Query: %w", err)
	}

	return &TableRow{
		Rows: rows,
	}, nil
}

func (tableRow *TableRow) Scan(dest ...interface{}) error {
	err := tableRow.Rows.Scan(dest...)
	if err != nil {
		return fmt.Errorf("calling handler.Rows.Scan: %w", err)
	}

	return nil
}

func (tableRow *TableRow) Next() bool {
	return tableRow.Rows.Next()
}

func (handler *mysql) Execute(statement string, args ...interface{}) (mysql_gateway.Result, error) {
	res := new(SQLResult)

	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return nil, err
	}

	res.Result = result
	return res, nil
}

func (r *SQLResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r *SQLResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

func (tableRow *TableRow) Close() error {
	err := tableRow.Rows.Close()
	if err != nil {
		return fmt.Errorf("calling tableRow.Row.Close: %w", err)
	}

	return nil
}