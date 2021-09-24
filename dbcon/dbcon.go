package dbcon

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	sqlConn := "postgres://xhbbzrmj:kVBlt0mq5dUO61PcebjEh240_vmTiIr1@arjuna.db.elephantsql.com/xhbbzrmj"
	//connStr := "user:passwd@tcp(127.0.0.1:5432)/dbname"
	//connStr := "user:passwd@tcp(127.0.0.1:3306)/dbname"
	db, err := sql.Open("postgres", sqlConn)
	if err != nil {
		log.Fatal("can't connect to database : ", err)
	}
	return db
}
