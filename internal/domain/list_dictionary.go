package domain

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

// ListDictionary list dictionary
func (i *InquiryService) ListDictionary(ctx context.Context, dictionaryType domain.DictionaryType) ([]*domain.Dictionary, error) {
	if dictionaryType != domain.AllDictionaryType {
		pairs, err := i.database.ListDictionary(ctx, dictionaryType)
		if err != nil {
			return nil, err
		}
		return []*domain.Dictionary{{
			Type:  dictionaryType,
			Pairs: pairs,
		}}, nil
	}
	response := make([]*domain.Dictionary, 0, 5)
	for j := 1; j <= 5; j++ {
		pairs, err := i.database.ListDictionary(ctx, domain.DictionaryType(j))
		if err != nil {
			return nil, err
		}
		response = append(response, &domain.Dictionary{
			Type:  domain.DictionaryType(j),
			Pairs: pairs,
		})
	}

	return response, nil
}
