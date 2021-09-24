package dbcon

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	sqlConn := os.Getenv("DB_URL")
	//connStr := "user:passwd@tcp(127.0.0.1:5432)/dbname"
	db, err := sql.Open("postgres", sqlConn)
	if err != nil {
		log.Fatal("can't connect to database : ", err)
	}
	return db
}
