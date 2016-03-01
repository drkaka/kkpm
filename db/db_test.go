package db

import "testing"

var (
	// DBName the PostgresSQL database name.
	DBName string
	// DBHost the host address.
	DBHost string
	// DBUser the db user.
	DBUser string
	//DBPassword the db password.
	DBPassword string
)

func prepareDB(t *testing.T) {
	PrepareDB()
}

func TestMain(t *testing.T) {

}
