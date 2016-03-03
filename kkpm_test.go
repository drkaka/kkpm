package kkpm

import (
	"net"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx"
)

func testInsert(t *testing.T) {
	if err := InsertMessage(0, 1, "abc"); err == nil {
		t.Error("Should have error that from id is 0.")
	}

	if err := InsertMessage(1, 0, "abc"); err == nil {
		t.Error("Should have error that to id is 0.")
	}

	if err := InsertMessage(1, 1, "abc"); err == nil {
		t.Error("Should have error that to ids are the same.")
	}

	if err := InsertMessage(1, 2, ""); err == nil {
		t.Error("Should have error that message is empty.")
	}
}

func TestMain(t *testing.T) {
	DBName := os.Getenv("dbname")
	DBHost := os.Getenv("dbhost")
	DBUser := os.Getenv("dbuser")
	DBPassword := os.Getenv("dbpassword")

	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     DBHost,
			User:     DBUser,
			Password: DBPassword,
			Database: DBName,
			Dial:     (&net.Dialer{KeepAlive: 5 * time.Minute, Timeout: 5 * time.Second}).Dial,
		},
		MaxConnections: 10,
	}

	var err error
	var pool *pgx.ConnPool
	if pool, err = pgx.NewConnPool(connPoolConfig); err != nil {
		t.Fatal(err)
	}

	if err = Use(pool); err != nil {
		t.Fatal(err)
	}
	testTableGeneration(t)

	// test the db methods.
	// insert some data.
	insertSomeMessages(t)

	testGetMessageFrom(t)
	testGetMessageTo(t)
	testGetMessageFromTo(t)

	truncate(t)

	// test the public methods
	testInsert(t)
}
