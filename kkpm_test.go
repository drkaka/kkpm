package kkpm

import (
	"net"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx"
)

func TestInvalidInsert(t *testing.T) {
	if err := InsertMessage(1, 1, "abc"); err == nil {
		t.Error("Should have error that to ids are the same.")
	}

	if err := InsertMessage(1, 2, "   "); err == nil {
		t.Error("Should have error that message is empty.")
	}
}

func testValidInsert(t *testing.T) {
	if err := InsertMessage(3, 2, "abc"); err != nil {
		t.Error(err)
	}

	if err := InsertMessage(3, 2, "abc"); err != nil {
		t.Error(err)
	}

	if err := InsertMessage(3, 1, "abc"); err != nil {
		t.Error(err)
	}
}

func testPubReadFunctions(t *testing.T) {
	if count, err := GetUnreadCount(2); err != nil {
		t.Error(err)
	} else if count != 2 {
		t.Error("Count is wrong.")
	}

	if err := ReadFrom(2, 3); err != nil {
		t.Error(err)
	}

	if count, err := GetUnreadCount(2); err != nil {
		t.Error(err)
	} else if count != 0 {
		t.Error("Count is wrong.")
	}
}

func testGetSentMessages(t *testing.T) {
	if result, err := GetSentMessages(3, 0); err != nil {
		t.Error(err)
	} else {
		if len(result) != 3 {
			t.Error("result not correct.")
		}
	}
}

func testGetReceivedMessages(t *testing.T) {
	if result, err := GetReveivedMessages(2, 0); err != nil {
		t.Error(err)
	} else {
		if len(result) != 2 {
			t.Error("result not correct.")
		}
	}
}

func testGetPeerChat(t *testing.T) {
	if result, err := GetPeerChat(3, 1, 0); err != nil {
		t.Error(err)
	} else {
		if len(result) != 1 {
			t.Error("result not correct.")
		}
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
	testDBMethods(t)

	// test the public methods
	testValidInsert(t)
	testPubReadFunctions(t)
	testGetSentMessages(t)
	testGetReceivedMessages(t)
	testGetPeerChat(t)

	if _, err = dbPool.Exec("DROP TABLE private_msg"); err != nil {
		t.Error(err)
	}
}
