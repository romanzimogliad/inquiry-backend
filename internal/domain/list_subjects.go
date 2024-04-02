package domain

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (i *InquiryService) ListSubjects(ctx context.Context) ([]*domain.Subject, error) {
	return i.database.ListSubjects(ctx)
}
