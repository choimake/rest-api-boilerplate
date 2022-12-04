package login

import (
	"context"
	"rest-api-boilerplate/internal/entity/user"
)

var _ UseCase = (*usecase)(nil)

type UseCase interface {
	Handle(ctx context.Context, input *Input) (*Output, error)
}

type usecase struct {
	userRepository user.Repository
}

func NewUseCase(repository user.Repository) UseCase {
	return &usecase{
		userRepository: repository,
	}
}

func (u usecase) Handle(ctx context.Context, input *Input) (*Output, error) {

	usr, err := u.userRepository.FetchByIdAndPassword(ctx, input.Id, input.Password)
	if err != nil {
		return nil, err
	}

	o := outputTranslator(usr)

	return o, nil
}
