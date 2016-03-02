package kkpm

const (
	insert = "INSERT INTO private_msg(id,from_userid,to_userid,message,at) VALUES($1,$2,$3,$4,$5)"
)

// prepareDB to prepare the database.
func prepareDB() {
	s := `CREATE TABLE IF NOT EXISTS private_msg (
	id uuid primary key,
	from_userid integer,
    to_userid integer,
    message text,
    at integer);
    CREATE INDEX IF NOT EXISTS index_private_msg_to_userid ON private_msg (to_userid);
    CREATE INDEX IF NOT EXISTS index_private_msg_from_userid ON private_msg (from_userid);`

	if _, err := dbPool.Exec(s); err != nil {
		panic(err)
	}
}

// insertMessage to insert a message to database.
func insertMessage(info *MessageInfo) {
	if _, err := dbPool.Exec(insert, info.MessageID, info.FromUser, info.ToUser, info.Message, info.At); err != nil {
		panic(err)
	}
}
