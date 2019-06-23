package helper

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // sql behavior modified
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "gituser"
	password = "passme123"
	dbname   = "mydb"
)

func InitDB() (*sql.DB, error) {
	var connectionString = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	} else {
		stmt, err := db.Prepare("CREATE TABLE WEB_URL(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		res, err := stmt.Exec()
		log.Println(res)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		fmt.Println("Table has been created successfully!")
		return db, nil
	}
}
