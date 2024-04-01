package domain

import "context"

func (i *InquiryService) Ping(ctx context.Context) (string, error) {
	return i.database.Ping(ctx)
}
