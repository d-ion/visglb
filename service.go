package visglb

import (
	pb "github.com/d-ion/isglb/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func marshalStatusMapToJSONList(sm map[string]*pb.SFUStatus) []byte {
	sb := []byte{'['}
	for _, status := range sm {
		if b, err := protojson.Marshal(status); err == nil {
			sb = append(sb, b...)
			sb = append(sb, ',')
		}
	}
	if len(sb) <= 1 {
		sb = append(sb, ']')
	} else {
		sb[len(sb)-1] = ']'
	}
	return sb
}

type Service struct {
	*VisualizedAlgorithm
}

func NewService(algorithm *VisualizedAlgorithm) Service {
	return Service{VisualizedAlgorithm: algorithm}
}

func (s Service) GetStatusListJSON() []byte {
	return marshalStatusMapToJSONList(s.GetStatusMap())
}

func (s Service) FetchStatusListJSON() []byte {
	return marshalStatusMapToJSONList(s.FetchStatusMap())
}

func marshalReportListToJSONList(lr []*pb.QualityReport) []byte {
	sb := []byte{'['}
	for _, report := range lr {
		if b, err := protojson.Marshal(report); err == nil {
			sb = append(sb, b...)
			sb = append(sb, ',')
		}
	}
	if len(sb) <= 1 {
		sb = append(sb, ']')
	} else {
		sb[len(sb)-1] = ']'
	}
	return sb
}

func (s Service) FetchReportListJSON() []byte {
	return marshalReportListToJSONList(s.FetchReportList())
}
