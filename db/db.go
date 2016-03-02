package db

import "github.com/jackc/pgx"

const (
	insert = "INSERT INTO private_msg(id,from_userid,to_userid,message,at) VALUES($1,$2,$3,$4,$5)"
)

// DBPool the pgx database pool.
var DBPool *pgx.ConnPool

// MessageInfo for private message.
type MessageInfo struct {
	MessageID string `json:"messageid" binding:"required"`
	Message   string `json:"message" binding:"required"`
	FromUser  int    `json:"fromuser" binding:"required"`
	ToUser    int    `json:"touser" binding:"required"`
	At        int    `json:"at" binding:"required"`
}

// PrepareDB to prepare the database.
func PrepareDB() {
	s := `CREATE TABLE IF NOT EXISTS private_msg (
	id uuid primary key,
	from_userid integer,
    to_userid integer,
    message text,
    at integer);
    CREATE INDEX IF NOT EXISTS index_private_msg_to_userid ON private_msg (to_userid);
    CREATE INDEX IF NOT EXISTS index_private_msg_from_userid ON private_msg (from_userid);`

	if _, err := DBPool.Exec(s); err != nil {
		panic(err)
	}
}

// InsertMessage to insert a message to database.
func InsertMessage(info *MessageInfo) {
	if _, err := DBPool.Exec(insert, info.MessageID, info.FromUser, info.ToUser, info.Message, info.At); err != nil {
		panic(err)
	}
}
