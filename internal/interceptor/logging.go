package interceptor

import (
	"context"

	"github.com/romanzimoglyad/inquiry-backend/internal/logger"

	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	//logger.Info().Msgf("%s : %s", info.FullMethod, req)
	logger.Debug().Msgf("%s", info.FullMethod)
	res, err := handler(ctx, req)
	if err != nil {
		logger.Error().Msgf("%s : %s", info.FullMethod, err)
		return nil, err
	}
	if info.FullMethod != "/inquiry.Inquiry/ListLessons" && info.FullMethod != "/inquiry.Inquiry/ListDictionary" {
		logger.Debug().Msgf("%s : %s", info.FullMethod, res)
	}

	return res, nil
}
