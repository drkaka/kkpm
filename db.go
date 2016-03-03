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
    message text,
    at integer);
    CREATE INDEX IF NOT EXISTS index_private_msg_to_userid ON private_msg (to_userid);
    CREATE INDEX IF NOT EXISTS index_private_msg_from_userid ON private_msg (from_userid);`

	_, err := dbPool.Exec(s)
	return err
}

// insertMessage to insert a message to database.
func insertMessage(info *MessageInfo) error {
	_, err := dbPool.Exec(insert, info.MessageID, info.FromUser, info.ToUser, info.Message, info.At)
	return err
}

// getMessagesFromUser to get the messages sent by the user with fromid.
func getMessagesFrom(fromid int32) ([]MessageInfo, error) {
	rows, _ := dbPool.Query("select id,to_userid,message,at from private_msg where from_userid=$1", fromid)

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
func getMessagesTo(toid int32) ([]MessageInfo, error) {
	rows, _ := dbPool.Query("select id,from_userid,message,at from private_msg where to_userid=$1", toid)

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

func getMessagesFromTo(fromid, toid int32) ([]MessageInfo, error) {
	rows, _ := dbPool.Query("select id,message,at from private_msg where to_userid=$1 and from_userid=$2", toid, fromid)

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
