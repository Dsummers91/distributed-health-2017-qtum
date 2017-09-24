package main

import (
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type DeployContractRequest struct {
	isSponsor   bool
	Sponsors    []string
	Individuals []string
}

type GetContractInfoResponse struct {
	Address   string `json:"address"`
	Balance   string `json:"balance"`
	Result    bool   `json:"finished"`
	StartDate uint   `json:"startDate"`
	EndDate   uint   `json:"endDate"`
}

func startAPI() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/createContract", deployContractAPI).Methods("GET")
	router.HandleFunc("/contractInfo", getContractInfo).Methods("GET")
	router.HandleFunc("/joinContract/sponsor", joinContractSponsor).Methods("GET")
	router.HandleFunc("/joinContract/participant", joinContractParticipant).Methods("GET")
	router.HandleFunc("/joinContract/company", joinContractCompany).Methods("GET")
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
	generateBlocks()
}

func deployContractAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	response := DeployContract()
	w.Write(response)
	generateBlocks()
}

func getContractInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var response GetContractInfoResponse
	response.Address = mostRecentAddress

	responseAsBytes, _ := json.Marshal(response)
	w.Write(responseAsBytes)
	generateBlocks()

}

func joinContractCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	response := joinContractBlockchain(3)
	w.Write(response)
	generateBlocks()
}

func joinContractParticipant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	response := joinContractBlockchain(2)
	w.Write(response)
	generateBlocks()
}

func joinContractSponsor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	response := joinContractBlockchain(1)
	w.Write(response)
	generateBlocks()
}

func participateContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	response := participateContractBlockchain()
	w.Write(response)
	generateBlocks()
}

func finishContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	response := finishContractBlockchain()
	generateBlocks()
}

func refundContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	response := finishContractBlockchain()
	generateBlocks()
}
