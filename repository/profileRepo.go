package repository

import (
	"log"

	dbcon "gosoft.co.th/mygo-lab/dbcon"
	st "gosoft.co.th/mygo-lab/structs"
)

func InsertProfileRepo(req st.Profile) (*int64, error) {
	db := dbcon.GetConnection()
	defer db.Close()

	sqlStr := `insert into profiles(id,name,email,age) 
	values(?,?,?,?)`
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
	id, _ := resp.LastInsertId()
	return &id, nil
}

func ListProfileRepo() ([]st.Profile, error) {
	db := dbcon.GetConnection()
	defer db.Close()

	sqlStr := `select * from profiles order by id asc`
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
