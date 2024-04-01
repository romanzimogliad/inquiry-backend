package grpc_service

import (
	"github.com/romanzimoglyad/inquiry-backend/internal/domain"
	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
)

type Implementation struct {
	inquiry.UnimplementedInquiryServer
	inquiryService *domain.InquiryService
}

func NewInquiryV1(inquiryService *domain.InquiryService) *Implementation {
	return &Implementation{inquiryService: inquiryService}
}
