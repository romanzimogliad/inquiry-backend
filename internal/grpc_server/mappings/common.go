package mappings

import inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"

type IdName interface {
	GetId() int32
	GetName() string
}

type IdNames []IdName

func ToIdName(T IdName) *inquiry.IdName {
	return &inquiry.IdName{
		Id:   T.GetId(),
		Name: T.GetName(),
	}
}
