package domain

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (i *InquiryService) GetMethodToRole(ctx context.Context) ([]*domain.MethodToRole, error) {
	return i.database.GetMethodToRole(ctx)
}
