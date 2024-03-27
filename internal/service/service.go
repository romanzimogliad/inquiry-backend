package service

import inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"

type Implementation struct {
	inquiry.UnimplementedInquiryServer
}

func NewInquiryV1() *Implementation {
	return &Implementation{}
}
