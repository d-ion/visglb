package visglb

import (
	pb "github.com/d-ion/isglb/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func marshalStatusMapToJSONList(sm map[string]*pb.SFUStatus) []byte {
	if len(sm) <= 0 {
		return []byte{'[', ']'}
	}
	sb := make([]byte, 0)
	sb = append(sb, '[')
	for _, status := range sm {
		if b, err := protojson.Marshal(status); err == nil {
			sb = append(sb, b...)
		}
		sb = append(sb, ',')
	}
	sb[len(sb)-1] = ']'
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
	if len(lr) <= 0 {
		return []byte{'[', ']'}
	}
	sb := make([]byte, 0)
	sb = append(sb, '[')
	for _, status := range lr {
		if b, err := protojson.Marshal(status); err == nil {
			sb = append(sb, b...)
		}
		sb = append(sb, ',')
	}
	sb[len(sb)-1] = ']'
	return sb
}

func (s Service) FetchReportListJSON() []byte {
	return marshalReportListToJSONList(s.FetchReportList())
}
