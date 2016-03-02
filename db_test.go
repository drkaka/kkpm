package kkpm

import (
	"testing"

	"github.com/jackc/pgx"
	"github.com/satori/go.uuid"
)

func prepare(t *testing.T) {
	prepareDB()

	var dbname pgx.NullString
	if err := dbPool.QueryRow("SELECT 'public.private_msg'::regclass;").Scan(&dbname); err != nil {
		t.Fatal(err)
	}

	if dbname.String != "private_msg" {
		t.Fatal("dbname is not correct.")
	}
}

func insertSomeMessages(t *testing.T) {
	var err error

	var one MessageInfo
	one.At = 2016
	one.FromUser = 2
	one.ToUser = 3
	one.MessageID = uuid.NewV1().String()
	one.Message = "message"

	if err = insertMessage(&one); err != nil {
		t.Error(err)
	}

	one.MessageID = uuid.NewV1().String()
	if err = insertMessage(&one); err != nil {
		t.Error(err)
	}
}

func testGetMessageFrom(t *testing.T) {
	if result, err := getMessagesFrom(2); err != nil {
		t.Error(err)
	} else {
		if len(result) != 2 {
			t.Error("result not correct.")
		}
	}
}

func truncate(t *testing.T) {
	if _, err := dbPool.Exec("TRUNCATE private_msg"); err != nil {
		t.Error(err)
	}
}
