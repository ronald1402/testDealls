package database

import (
	"database/sql"
	"fmt"
	"github.com/apex/log"
	_ "github.com/go-sql-driver/mysql"
	"testDealls/config"
	"time"
)

func Connect(c *config.Database) *sql.DB {
	var dbConn *sql.DB
	success := false
	for i := 1; i <= c.MaxReconnectRetry; i++ {
		db, err := connect(c.Mysql)
		if err != nil {
			continue
		}
		dbConn = db
		success = true
		break
	}

	if !success {
		log.Fatal("can't connect to mysql")
	}

	log.Info("Successfully connected to the database!")

	return dbConn
}

func connect(connConfig *config.MySqlConnConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		connConfig.Username, connConfig.Password, connConfig.Host, connConfig.Port, connConfig.Schema)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// Ping the database to check if the connection is alive
	if err := db.Ping(); err != nil {
		log.Fatalf("Cannot connect to the database: %v", err)
	}

	if connConfig.MaxOpenConn > 0 {
		db.SetMaxOpenConns(connConfig.MaxOpenConn)
	}
	if connConfig.MaxConnLifeTime > 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(connConfig.MaxConnLifeTime))
	}
	if connConfig.MaxIdleConns > 0 {
		db.SetMaxIdleConns(connConfig.MaxIdleConns)
	}

	return db, nil
}

func Disconnect(db *sql.DB) {
	log.Info("disconnecting mysql connection")
	err := db.Close()
	if err != nil {
		log.WithError(err).Error("error on closing mysql connection")
		return
	}
	log.Info("disconnected from mysql")
}
