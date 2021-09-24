package service

import (
	"encoding/json"
	"log"
	"net/http"

	repo "gosoft.co.th/mygo-lab/repository"
	st "gosoft.co.th/mygo-lab/structs"
)

func ProfileSubmitSv(w http.ResponseWriter, r *http.Request) {
	var reqParam st.Profile
	err := json.NewDecoder(r.Body).Decode(&reqParam)
	if nil != err {
		log.Panic(err.Error())
		return
	}

	id, sqlErr := repo.InsertProfileRepo(reqParam)
	if nil != sqlErr {
		log.Panic(sqlErr.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
}

func ProfileDetailSv(w http.ResponseWriter, r *http.Request) {
	datas, sqlErr := repo.ListProfileRepo()
	if nil != sqlErr {
		log.Panic(sqlErr.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datas)
}

func ProfileListSv(w http.ResponseWriter, r *http.Request) {
	datas, sqlErr := repo.ListProfileRepo()
	if nil != sqlErr {
		log.Panic(sqlErr.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datas)
}

func ProfileDeleteSv(w http.ResponseWriter, r *http.Request) {
	datas, sqlErr := repo.ListProfileRepo()
	if nil != sqlErr {
		log.Panic(sqlErr.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datas)
}
