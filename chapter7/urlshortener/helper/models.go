package helper

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // sql behavior modified
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "gituser"
	password = "passme"
	dbname   = "mydb"
)

// InitDB initializes database table
func InitDB() (*sql.DB, error) {
	var connectionString = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS web_url(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec()

	if err != nil {
		return nil, err
	}

	return db, nil
}
