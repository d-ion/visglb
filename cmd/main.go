package main

import (
	"github.com/d-ion/visglb"
	"net/http"
)

func main() {
	server := visglb.NewServer(nil)
	http.HandleFunc("/GetStatusListJSON", server.HandleGetStatusListJSON)
	http.HandleFunc("/FetchStatusListJSON", server.HandleFetchStatusListJSON)
	http.HandleFunc("/FetchReportListJSON", server.HandleFetchReportListJSON)
	_ = http.ListenAndServe(":80", nil)
}
