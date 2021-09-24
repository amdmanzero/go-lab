package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	repo "gosoft.co.th/mygo-lab/repository"
	st "gosoft.co.th/mygo-lab/structs"
)

func ProfileSubmitSv(w http.ResponseWriter, r *http.Request) {
	var reqParam st.Profile
	err := json.NewDecoder(r.Body).Decode(&reqParam)
	if err != nil {
		log.Panic(err.Error())
		return
	}

	id, sqlErr := repo.InsertProfileRepo(reqParam)
	if sqlErr != nil {
		log.Panic(sqlErr.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
}

func ProfileDetailSv(w http.ResponseWriter, r *http.Request) {
	strID := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		log.Panic(err.Error())
		return
	}
	profile, sqlErr := repo.GetProfileRepo(id)
	if sqlErr != nil {
		log.Panic(sqlErr.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func ProfileListSv(w http.ResponseWriter, r *http.Request) {
	datas, sqlErr := repo.ListProfileRepo()
	if sqlErr != nil {
		log.Panic(sqlErr.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datas)
}

func ProfileDeleteSv(w http.ResponseWriter, r *http.Request) {
	strID := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		log.Panic(err.Error())
		return
	}
	rowNum, sqlErr := repo.DeleteProfileRepo(id)
	if sqlErr != nil {
		log.Panic(sqlErr.Error())
		return
	}

	resp := st.DelResult{
		RowNum: *rowNum,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
