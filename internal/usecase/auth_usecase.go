package usecase

type AuthUsecase interface {
}

type authUsecaseImpl struct {
}

func NewAuthUsecase() AuthUsecase {
	return &authUsecaseImpl{}
}
