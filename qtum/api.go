package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func startAPI() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", TestFunc).Methods("GET")
	go http.ListenAndServe(":8888", router)
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// value, _ := hex.DecodeString(string(bytes))
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("just works fam"))
}
