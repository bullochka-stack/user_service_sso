package validators

import (
	ssov1 "github.com/bullochka-stack/protos/gen/go/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateRegisterRequest(req *ssov1.RegisterRequest) error {
	// TODO: подключить библиотеку валидации
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "missing email")
	}

	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "missing password")
	}

	return nil
}
