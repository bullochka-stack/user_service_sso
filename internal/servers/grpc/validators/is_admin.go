package validators

import (
	ssov1 "github.com/bullochka-stack/protos/gen/go/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	emptyUserId = 0
)

func ValidateIsAdminRequest(req *ssov1.IsAdminRequest) error {
	// TODO: подключить библиотеку валидации
	if req.GetUserId() == emptyUserId {
		return status.Error(codes.InvalidArgument, "missing user_id")
	}

	return nil
}
