package domain

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

// ListDictionary list dictionary
func (i *InquiryService) ListDictionary(ctx context.Context, dictionaryType domain.DictionaryType) ([]*domain.IdName, error) {
	return i.database.ListDictionary(ctx, dictionaryType)
}
