package db

import "github.com/jackc/pgx"

// DBPool the pgx database pool.
var DBPool *pgx.ConnPool

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
