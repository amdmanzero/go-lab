package repository

import (
	"log"

	dbcon "gosoft.co.th/mygo-lab/dbcon"
	st "gosoft.co.th/mygo-lab/structs"
)

func InsertProfileRepo(req st.Profile) (*int64, error) {
	db := dbcon.GetConnection()
	defer db.Close()

	sqlStr := `insert into profiles(id,name,email,age) values($1,$2,$3,$4)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	resp, err := stmt.Exec(req.ID, req.Name, req.Email, req.Age)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var id *int64
	numRows, _ := resp.RowsAffected()
	if numRows > 0 {
		id = req.ID
	}
	return id, nil
}

func ListProfileRepo() ([]st.Profile, error) {
	db := dbcon.GetConnection()
	defer db.Close()

	sqlStr := `select id,name,email,age from profiles order by id asc`
	var datas []st.Profile
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var tmpItem st.Profile
	for rows.Next() {
		rows.Scan(&tmpItem.ID, &tmpItem.Name, &tmpItem.Email, &tmpItem.Age)
		datas = append(datas, tmpItem)
	}
	return datas, nil
}

func GetProfileRepo(id int64) (*st.Profile, error) {
	db := dbcon.GetConnection()
	defer db.Close()

	sqlStr := `select id,name,email,age from profiles where id = $1`
	var data st.Profile
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&data.ID, &data.Name, &data.Email, &data.Age)
	}
	return &data, nil
}

func DeleteProfileRepo(id int64) (*int64, error) {
	db := dbcon.GetConnection()
	defer db.Close()

	sqlStr := `delete from profiles where id=$1`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	resp, err := stmt.Exec(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	rowNum, rowErr := resp.RowsAffected()
	if rowErr != nil {
		log.Println(rowErr.Error())
		return nil, err
	}
	return &rowNum, nil
}
