package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Network    string
	Driver     string
	Host       string
	Port       int
	UserName   string
	Password   string
	Database   string
	PoolSize   int
	Charset    string
	AutoCommit string
	ParseTime  string
}

type DBOption func(*DBConfig)

func Network(net string) DBOption {
	return func(cf *DBConfig) {
		cf.Network = net
	}
}
func Driver(driver string) DBOption {
	return func(cf *DBConfig) {
		cf.Driver = driver
	}
}
func Host(host string) DBOption {
	return func(cf *DBConfig) {
		cf.Host = host
	}
}
func Port(port int) DBOption {
	return func(cf *DBConfig) {
		cf.Port = port
	}
}
func UserName(username string) DBOption {
	return func(cf *DBConfig) {
		cf.UserName = username
	}
}
func Password(password string) DBOption {
	return func(cf *DBConfig) {
		cf.Password = password
	}
}
func Database(database string) DBOption {
	return func(cf *DBConfig) {
		cf.Database = database
	}
}
func PoolSize(size int) DBOption {
	return func(cf *DBConfig) {
		cf.PoolSize = size
	}
}
func Charset(charset string) DBOption {
	return func(cf *DBConfig) {
		cf.Charset = charset
	}
}
func AutoCommit(flag bool) DBOption {
	return func(cf *DBConfig) {
		cf.AutoCommit = fmt.Sprint(flag)
	}
}
func ParseTime(flag bool) DBOption {
	return func(cf *DBConfig) {
		cf.ParseTime = fmt.Sprint(flag)
	}
}

func CreateDB(opts ...DBOption) (*sql.DB, error) {
	cf := &DBConfig{
		Network:    "tcp",
		Driver:     "mysql",
		Host:       "127.0.0.1",
		Port:       3306,
		UserName:   "root",
		Password:   "",
		PoolSize:   4,
		Charset:    "utf8mb4",
		AutoCommit: "true",
		ParseTime:  "true",
	}
	for _, opt := range opts {
		opt(cf)
	}
	switch strings.ToLower(cf.Driver) {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&autocommit=%s&parseTime=%s",
			cf.UserName,
			cf.Password,
			cf.Network,
			cf.Host,
			cf.Port,
			cf.Database,
			cf.Charset,
			cf.AutoCommit,
			cf.ParseTime,
		)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, err
		}
		db.SetConnMaxLifetime(time.Hour)
		db.SetMaxIdleConns(cf.PoolSize)
		db.SetMaxOpenConns(cf.PoolSize)
		return db, nil
	}
	return nil, errors.New("unsupported database")
}
