package kkpm

import (
	"net"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx"
)

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
	if dbPool, err = pgx.NewConnPool(connPoolConfig); err != nil {
		t.Fatal(err)
	}

	prepare(t)
	insertSomeMessages(t)

	testGetMessageFrom(t)

	truncate(t)
}
