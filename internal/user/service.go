package user

import (
	//"context"
	"todo/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

//func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (User, error) {
//	return nil, nil
//}