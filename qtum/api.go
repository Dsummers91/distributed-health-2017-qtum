package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func startAPI() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", TestFunc).Methods("GET")
	router.HandleFunc("/createContract", DeployContractAPI).Methods("GET")
	go http.ListenAndServe(":8888", router)
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// value, _ := hex.DecodeString(string(bytes))
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("just works fam"))
}

func DeployContractAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	// DeployContract()
}
