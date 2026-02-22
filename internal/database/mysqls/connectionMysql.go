// Package myslqs is
package mysqls

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

func configConnectionDatabaseMysql() *mysql.Config {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = os.Getenv("DBNET")
	cfg.Addr = os.Getenv("DBADDR")
	cfg.DBName = os.Getenv("DBNAME")
	return cfg
}

func configHandleDatabaseMysql() (sqls *sql.DB, err error) {
	cfg := configConnectionDatabaseMysql()
	sqls, err = sql.Open(os.Getenv("DBDATABASE"), cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return sqls, nil
}

func ConnectionDatabaseMysql() (*sql.DB, error) {
	db, err := configHandleDatabaseMysql()
	if err != nil {
		return nil, err
	}
	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	return db, nil
}
