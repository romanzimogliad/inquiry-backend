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

	dictionaries := make([]*inquiry.Dictionary, 0, len(dictionary))
	for _, v := range dictionary {
		pairs := make([]*inquiry.IdName, 0, len(v.Pairs))
		for _, pair := range v.Pairs {
			pairs = append(pairs, mappings.ToIdName(pair))
		}
		dictionaries = append(dictionaries, &inquiry.Dictionary{
			Type:  inquiry.DictionaryType(v.Type),
			Pairs: pairs,
		})
	}

	return &inquiry.ListDictionaryResponse{
		Dictionaries: dictionaries,
	}, nil
}
