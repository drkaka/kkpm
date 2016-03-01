package db

import (
	"net"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx"
)

func prepareDB(t *testing.T) {
	PrepareDB()

	var dbname pgx.NullString
	if err := DBPool.QueryRow("SELECT 'public.private_msg'::regclass;").Scan(&dbname); err != nil {
		t.Fatal(err)
	}

	if dbname.String != "private_msg" {
		t.Fatal("dbname is not correct.")
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
	if DBPool, err = pgx.NewConnPool(connPoolConfig); err != nil {
		t.Fatal(err)
	}

	prepareDB(t)
}
