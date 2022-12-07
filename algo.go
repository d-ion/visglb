package visglb

import pb "github.com/d-ion/isglb/proto"

type ProtobufAlgorithm interface {
	UpdateSFUStatus(current []*pb.SFUStatus, reports []*pb.QualityReport) (expected []*pb.SFUStatus)
}

type VisualizedAlgorithm struct {
	pb.ProtobufAlgorithm

	lastStatusMap map[string]*pb.SFUStatus
	statusEvent   chan map[string]*pb.SFUStatus
	reportEvent   chan []*pb.QualityReport
}

func (a VisualizedAlgorithm) UpdateSFUStatus(current []*pb.SFUStatus, reports []*pb.QualityReport) (expected []*pb.SFUStatus) {
	statusMap := make(map[string]*pb.SFUStatus)
	for _, s := range current {
		statusMap[s.SFU.Id] = s // save new status
	}
	expected = a.ProtobufAlgorithm.UpdateSFUStatus(current, reports)
	for _, s := range expected {
		statusMap[s.SFU.Id] = s // save new status
	}
	statusChange := false
	for id, newStatus := range statusMap {
		if oldStatus, ok := a.lastStatusMap[id]; ok {
			if !oldStatus.Compare(newStatus) { // compare with old status
				statusChange = true // if change then tag it
			}
		}
	}
	a.lastStatusMap = statusMap
	if statusChange { // non-block
		select {
		case <-a.statusEvent:
		default:
		}
		select {
		case a.statusEvent <- statusMap:
		default:
		}
	}
	return expected
}

// GetStatusMap get the current status map
func (a VisualizedAlgorithm) GetStatusMap() map[string]*pb.SFUStatus {
	return a.lastStatusMap
}

// FetchStatusUpdate fetch the status update and get the latest status
// TODO: implement pubsub and implement fetch for multi user
func (a VisualizedAlgorithm) FetchStatusUpdate() map[string]*pb.SFUStatus {
	return <-a.statusEvent
}

// FetchReport fetch the reports
// TODO: implement pubsub and implement fetch for multi user
func (a VisualizedAlgorithm) FetchReport() []*pb.QualityReport {
	return <-a.reportEvent
}
