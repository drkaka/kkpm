package kkpm

import (
	"testing"

	"github.com/jackc/pgx"
	"github.com/satori/go.uuid"
)

var messages []MessageInfo

func testTableGeneration(t *testing.T) {
	var dbname pgx.NullString
	if err := dbPool.QueryRow("SELECT 'public.private_msg'::regclass;").Scan(&dbname); err != nil {
		t.Fatal(err)
	}

	if dbname.String != "private_msg" {
		t.Fatal("dbname is not correct.")
	}
}

func testDBMethods(t *testing.T) {
	insertSomeMessages(t)

	testGetMessageFrom(t)
	testGetPartialMessageFrom(t)
	testGetMessageTo(t)
	testGetPartialMessageTo(t)
	testGetMessageFromTo(t)

	truncate(t)
}

func insertSomeMessages(t *testing.T) {
	var err error

	var one MessageInfo
	one.At = 2016
	one.FromUser = 2
	one.ToUser = 3
	one.MessageID = uuid.NewV1().String()
	one.Message = "message"
	messages = append(messages, one)

	if err = insertMessage(&one); err != nil {
		t.Error(err)
	}

	var two MessageInfo
	two.At = 2018
	two.FromUser = 2
	two.ToUser = 3
	two.Message = "message"
	two.MessageID = uuid.NewV1().String()
	messages = append(messages, two)

	if err = insertMessage(&two); err != nil {
		t.Error(err)
	}

	var three MessageInfo
	three.At = 2016
	three.FromUser = 2
	three.ToUser = 4
	three.Message = "message"
	three.MessageID = uuid.NewV1().String()
	messages = append(messages, three)

	if err = insertMessage(&three); err != nil {
		t.Error(err)
	}
}

func testGetMessageFrom(t *testing.T) {
	if result, err := getMessagesFrom(2, 0); err != nil {
		t.Error(err)
	} else {
		if len(result) != 3 {
			t.Error("result not correct.")
		}
	}
}

func testGetPartialMessageFrom(t *testing.T) {
	if result, err := getMessagesFrom(2, 2017); err != nil {
		t.Error(err)
	} else {
		if len(result) != 1 {
			t.Error("result not correct.")
		}
	}
}

func testGetMessageTo(t *testing.T) {
	if result, err := getMessagesTo(3, 0); err != nil {
		t.Error(err)
	} else {
		if len(result) != 2 {
			t.Error("result not correct.")
		}
	}
}

func testGetPartialMessageTo(t *testing.T) {
	if result, err := getMessagesTo(3, 2017); err != nil {
		t.Error(err)
	} else {
		if len(result) != 1 {
			t.Error("result not correct.")
		}
	}
}

func testGetMessageFromTo(t *testing.T) {
	if result, err := getMessagesFromTo(2, 4, 0); err != nil {
		t.Error(err)
	} else {
		if len(result) != 1 {
			t.Error("result not correct.")
		}
	}
}

func truncate(t *testing.T) {
	if _, err := dbPool.Exec("TRUNCATE private_msg"); err != nil {
		t.Error(err)
	}
}
