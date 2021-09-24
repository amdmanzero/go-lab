package dbcon

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	sqlConn := "xhbbzrmj:kVBlt0mq5dUO61PcebjEh240_vmTiIr1@tcp(arjuna.db.elephantsql.com)/xhbbzrmj"
	db, err := sql.Open("postgres", sqlConn)
	if err != nil {
		log.Fatal("can't connect to database : ", err)
	}
	return db
}
