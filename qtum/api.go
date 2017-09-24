package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type DeployContractRequest struct {
	isSponsor   bool
	Sponsors    []string
	Individuals []string
}

type GetContractInfoResponse struct {
	
}
func startAPI() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/createContract", deployContractAPI).Methods("GET")
	router.HandleFunc("/contractInfo", getContractInfo).Methods("GET")
	router.HandleFunc("/joinContract", joinContract).Methods("GET")
	router.HandleFunc("/participateContract", participateContract).Methods("GET")
	router.HandleFunc("/finishContract", finishContract).Methods("GET")
	router.HandleFunc("/refundContract", refundContract).Methods("GET")
	go http.ListenAndServe(":8888", router)
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// value, _ := hex.DecodeString(string(bytes))
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("just works fam"))
}

func deployContractAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("{success: true}"))
}

func getContractInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var 
}

func joinContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}

func participateContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}

func finishContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}

func refundContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}
