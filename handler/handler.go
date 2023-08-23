package handler

import (
	"encoding/json"
	"github.com/aws/aws-xray-sdk-go/xray"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, seg := xray.BeginSegment(r.Context(), "HelloSegment")
	defer seg.Close(nil)

	msg := Message{Text: "Hello World"}
	response, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-type", "application/Json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetCustData(w http.ResponseWriter, r *http.Request) {
	_, seg := xray.BeginSegment(r.Context(), "HelloSegment")
	defer seg.Close(nil)

	msg := Message{Text: "in get cust method"}
	resp, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-type", "application/Json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func Path(w http.ResponseWriter, r *http.Request) {
	_, seg := xray.BeginSegment(r.Context(), "HelloSegment")
	defer seg.Close(nil)
	ms := Message{Text: "hello"}
	res, err := json.Marshal(ms)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/Json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
