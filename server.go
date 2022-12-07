package visglb

import (
	pb "github.com/d-ion/isglb/proto"
	"net/http"
)

type Server struct {
	Service
}

func NewServer(algorithm pb.ProtobufAlgorithm) Server {
	return Server{NewService(algorithm)}
}

func (s Server) HandleGetStatusListJSON(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("content-type", "text/json")
	w.WriteHeader(200)
	_, _ = w.Write(s.GetStatusListJSON())
}

func (s Server) HandleFetchStatusListJSON(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("content-type", "text/json")
	w.WriteHeader(200)
	_, _ = w.Write(s.FetchStatusListJSON())
}

func (s Server) HandleFetchReportListJSON(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("content-type", "text/json")
	w.WriteHeader(200)
	_, _ = w.Write(s.FetchReportListJSON())
}
