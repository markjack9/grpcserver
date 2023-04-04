package trip

import (
	"context"
	trippb "server/proto/gen/go"
)

// Service implementation trip service.
type Service struct {
}

func (*Service) GetTrip(c context.Context,
	req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "abc",
			End:         "edb",
			DurationSec: 3600,
			FeeCent:     1000,
			//调用新的添加的数据结构
			StartPos: &trippb.Location{
				Latitude:  30,
				Longitude: 20,
			},
			EndPos: &trippb.Location{
				Latitude:  32,
				Longitude: 25,
			},
			Status: trippb.TripStatus_FINISHED,
			PathLocations: []*trippb.Location{
				//创建一个切片存储trippb的经纬度数据
				//数据值进行引用
				{
					Latitude:  31,
					Longitude: 22,
				},
				{
					Latitude:  32,
					Longitude: 23,
				},
			},
		},
	}, nil
}
