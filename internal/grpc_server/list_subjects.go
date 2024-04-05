package grpc_server

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/grpc_server/mappings"
	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListSubjects list subjects
func (i *Implementation) ListSubjects(ctx context.Context, request *inquiry.Empty) (*inquiry.ListSubjectsResponse, error) {
	lessons, err := i.inquiryService.ListSubjects(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error in ListSubjects: %v", err)
	}

	resp := make([]*inquiry.IdName, len(lessons))

	for k, v := range lessons {
		resp[k] = mappings.ToIdName(v)
	}
	return &inquiry.ListSubjectsResponse{
		Subjects: resp,
	}, nil
}
