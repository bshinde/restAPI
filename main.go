package main

import (
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/gorilla/mux"
	"net/http"
	"restAPI/handler"
)

func main() {
	xray.Configure(xray.Config{LogLevel: "info"})
	router := mux.NewRouter()
	router.HandleFunc("/hello", handler.HelloHandler).Methods("GET")
	router.HandleFunc("/getCustData", handler.GetCustData).Methods("GET")
	router.HandleFunc("/path", handler.Path).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)

}
