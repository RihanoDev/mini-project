package users

import "api-mini-project/app/middlewares"

type userUsecase struct {
	userRepository Repository
	jwtAuth        *middlewares.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &userUsecase{
		userRepository: ur,
		jwtAuth:        jwtAuth,
	}
}

func (uu *userUsecase) Register(userDomain *Domain) Domain {
	return uu.userRepository.Register(userDomain)
}

func (uu *userUsecase) Login(userDomain *Domain) string {
	user := uu.userRepository.GetByEmail(userDomain)

	if user.ID == 0 {
		return ""
	}

	token := uu.jwtAuth.GenerateToken(int(user.ID))

	return token
}

func (uu *userUsecase) CheckData(userDomain *Domain) Domain {
	return uu.userRepository.CheckData(userDomain)
}
