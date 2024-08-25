package controller

import (
	"encoding/json"
	"github.com/hewo233/cxsj_homework1/common"
	"github.com/hewo233/cxsj_homework1/model"
	"net/http"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	email := queryParam.Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	userInfo, err := common.QueryRecordByEmail(email, common.UserDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	userList, err := common.ListAllRecords(common.UserDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(userList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	var userInfo model.CxsjUser
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = common.AddRecord(&userInfo, common.UserDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func ModifyHandler(w http.ResponseWriter, r *http.Request) {
	var userInfo model.CxsjUser
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = common.ModifyRecord(&userInfo, common.UserDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	email := queryParam.Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	err := common.DeleteRecord(email, common.UserDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
