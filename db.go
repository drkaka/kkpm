package kkpm

import "github.com/jackc/pgx"

const (
	insert = "INSERT INTO private_msg(id,from_userid,to_userid,message,at) VALUES($1,$2,$3,$4,$5)"
)

// dbPool the pgx database pool.
var dbPool *pgx.ConnPool

// prepareDB to prepare the database.
func prepareDB() error {
	s := `CREATE TABLE IF NOT EXISTS private_msg (
	id uuid primary key,
	from_userid integer,
    to_userid integer,
	read boolean DEFAULT false, 
    message text,
    at integer);
    CREATE INDEX IF NOT EXISTS index_private_msg_to_userid ON private_msg (to_userid);
    CREATE INDEX IF NOT EXISTS index_private_msg_from_userid ON private_msg (from_userid);
    CREATE INDEX IF NOT EXISTS index_private_msg_at ON private_msg (at);`

	_, err := dbPool.Exec(s)
	return err
}

// insertMessage to insert a message to database.
func insertMessage(info *MessageInfo) error {
	_, err := dbPool.Exec(insert, info.MessageID, info.FromUser, info.ToUser, info.Message, info.At)
	return err
}

// getUnreadCount to get the count of unread messages.
func getUnreadCount(toid int32) (int32, error) {
	s := "SELECT COUNT(1) FROM private_msg WHERE to_userid=$1 AND read=false"
	var count int64
	if err := dbPool.QueryRow(s, toid).Scan(&count); err != nil {
		return 0, err
	}
	return int32(count), nil
}

// readFrom to mark all as read form the id.
func readFrom(toid, fromid int32) error {
	s := "UPDATE private_msg SET read=true WHERE to_userid=$1 AND from_userid=$2 AND read=false"
	if _, err := dbPool.Exec(s, toid, fromid); err != nil {
		return err
	}
	return nil
}

// getMessagesFromUser to get the messages sent by the user with fromid.
// utime the unixtime, the messages will be got after that time.
func getMessagesFrom(fromid, utime int32) ([]MessageInfo, error) {
	s := "SELECT id,to_userid,message,at FROM private_msg WHERE from_userid=$1 AND at>$2"
	rows, _ := dbPool.Query(s, fromid, utime)

	var result []MessageInfo
	for rows.Next() {
		var one MessageInfo
		err := rows.Scan(&(one.MessageID), &(one.ToUser), &(one.Message), &(one.At))
		if err != nil {
			return result, err
		}
		one.FromUser = fromid
		result = append(result, one)
	}

	return result, rows.Err()
}

// getMessagesToUser to get the messages received by the user with toid.
// utime the unixtime, the messages will be got after that time.
func getMessagesTo(toid, utime int32) ([]MessageInfo, error) {
	s := "SELECT id,from_userid,message,at FROM private_msg WHERE to_userid=$1 AND at>$2"
	rows, _ := dbPool.Query(s, toid, utime)

	var result []MessageInfo
	for rows.Next() {
		var one MessageInfo
		err := rows.Scan(&(one.MessageID), &(one.FromUser), &(one.Message), &(one.At))
		if err != nil {
			return result, err
		}
		one.ToUser = toid
		result = append(result, one)
	}

	return result, rows.Err()
}

// getMessagesFromTo to get messages with a single user.
// utime the unixtime, the messages will be got after that time.
func getMessagesFromTo(fromid, toid, utime int32) ([]MessageInfo, error) {
	s := "SELECT id,message,at FROM private_msg WHERE to_userid=$1 AND from_userid=$2 AND at>$3"
	rows, _ := dbPool.Query(s, toid, fromid, utime)

	var result []MessageInfo
	for rows.Next() {
		var one MessageInfo
		err := rows.Scan(&(one.MessageID), &(one.Message), &(one.At))
		if err != nil {
			return result, err
		}
		one.FromUser = fromid
		one.ToUser = toid
		result = append(result, one)
	}

	return result, rows.Err()
}
