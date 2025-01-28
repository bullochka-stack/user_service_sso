package validators

import (
	ssov1 "github.com/bullochka-stack/protos/gen/go/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	emptyAppId = 0
)

func ValidateLoginRequest(req *ssov1.LoginRequest) error {
	// TODO: подключить библиотеку валидации
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "missing email")
	}

	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "missing password")
	}

	if req.GetAppId() == emptyAppId {
		return status.Error(codes.InvalidArgument, "missing app_id")
	}

	return nil
}
