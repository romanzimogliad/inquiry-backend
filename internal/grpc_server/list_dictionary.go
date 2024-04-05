package grpc_server

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"

	"github.com/romanzimoglyad/inquiry-backend/internal/grpc_server/mappings"
	inquiry "github.com/romanzimoglyad/inquiry-backend/pb/api_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListDictionary list dictionary
func (i *Implementation) ListDictionary(ctx context.Context, request *inquiry.ListDictionaryRequest) (*inquiry.ListDictionaryResponse, error) {
	dictionary, err := i.inquiryService.ListDictionary(ctx, domain.DictionaryType(request.GetType()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error in ListSubjects: %v", err)
	}

	resp := make([]*inquiry.IdName, len(dictionary))

	for k, v := range dictionary {
		resp[k] = mappings.ToIdName(v)
	}
	return &inquiry.ListDictionaryResponse{
		Pairs: resp,
	}, nil
}
