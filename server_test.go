package visglb

import (
	"github.com/d-ion/isglb"
	"github.com/d-ion/isglb/algorithms/random"
	pb "github.com/d-ion/isglb/proto"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

const sleep = 1000

type TestTransmissionReporter struct {
	random.RandTransmissionReport
}

func (t TestTransmissionReporter) Bind(ch chan<- *pb.TransmissionReport) {
	go func(ch chan<- *pb.TransmissionReport) {
		for {
			<-time.After(time.Duration(rand.Int31n(sleep)) * time.Millisecond)
			ch <- t.RandReport()
		}
	}(ch)
}

type TestComputationReporter struct {
	random.RandComputationReport
}

func (t TestComputationReporter) Bind(ch chan<- *pb.ComputationReport) {
	go func(ch chan<- *pb.ComputationReport) {
		for {
			<-time.After(time.Duration(rand.Int31n(sleep)) * time.Millisecond)
			ch <- t.RandReport()
		}
	}(ch)
}

type TestSessionTracker struct {
}

func (t TestSessionTracker) FetchSessionEvent() *isglb.SessionEvent {
	<-time.After(time.Duration(rand.Int31n(sleep)) * time.Millisecond)
	return &isglb.SessionEvent{
		Session: &pb.ClientNeededSession{
			Session: "",
			User:    "",
		}, State: isglb.SessionEvent_State(rand.Intn(2)),
	}
}

func TestServer(t *testing.T) {
	pba := Visualize(&random.Random{RandomTrack: true})
	syncer := isglb.NewSFUStatusSyncer(
		isglb.NewChanClientStreamFactory(
			isglb.NewService[*pb.SFUStatus](
				pb.AlgorithmWrapper{ProtobufAlgorithm: pba},
			),
		),
		&pb.Node{Id: "test"},
		isglb.ToolBox{
			TransmissionReporter: TestTransmissionReporter{random.RandTransmissionReport{}},
			ComputationReporter:  TestComputationReporter{random.RandComputationReport{}},
			SessionTracker:       TestSessionTracker{},
		})
	syncer.Start()
	defer syncer.Stop()

	server := NewServer(pba)
	http.HandleFunc("/GetStatusListJSON", server.HandleGetStatusListJSON)
	http.HandleFunc("/FetchStatusListJSON", server.HandleFetchStatusListJSON)
	http.HandleFunc("/FetchReportListJSON", server.HandleFetchReportListJSON)
	_ = http.ListenAndServe(":80", nil)
}
