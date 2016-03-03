package kkpm

import (
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx"
	"github.com/satori/go.uuid"
)

// MessageInfo for private message.
type MessageInfo struct {
	MessageID string `json:"messageid"`
	Message   string `json:"message"`
	FromUser  int32  `json:"fromuser,omitempty"`
	ToUser    int32  `json:"touser,omitempty"`
	At        int32  `json:"at"`
}

// Use the pool to do further operations.
func Use(pool *pgx.ConnPool) error {
	dbPool = pool
	return prepareDB()
}

// InsertMessage to send a message.
func InsertMessage(fromid, toid int32, message string) error {
	var msg MessageInfo

	if fromid <= 0 || toid <= 0 {
		return errors.New("id must larger than 0")
	}

	if fromid == toid {
		return errors.New("Can't send message to self.")
	}

	if len(strings.Trim(message, " ")) == 0 {
		return errors.New("message can't be empty")
	}

	msg.MessageID = uuid.NewV1().String()
	msg.At = int32(time.Now().Unix())
	msg.FromUser = fromid
	msg.ToUser = toid
	msg.Message = message

	return insertMessage(&msg)
}

// GetSentMessages to get all messages sent.
// utime the unixtime, the messages will be got after that time.
func GetSentMessages(userid, utime int32) ([]MessageInfo, error) {
	return getMessagesFrom(userid, utime)
}

// GetReveivedMessages to get all received messages.
// utime the unixtime, the messages will be got after that time.
func GetReveivedMessages(userid, utime int32) ([]MessageInfo, error) {
	return getMessagesTo(userid, utime)
}

// GetPeerChat to get the messages between two users.
// utime the unixtime, the messages will be got after that time.
func GetPeerChat(fromid, toid, utime int32) ([]MessageInfo, error) {
	return getMessagesFromTo(fromid, toid, utime)
}
