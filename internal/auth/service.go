package auth

import (
	"errors"
	"fmt"
	"links/internal/user"
	"links/pkg/jwt"
	"links/pkg/secure"
)

type AuthHService struct {
	UserRepository *user.UserRepository
	JWT            *jwt.JWT
}

func NewAuthService(rep *user.UserRepository, j *jwt.JWT) *AuthHService {
	return &AuthHService{
		UserRepository: rep,
		JWT:            j,
	}
}

func (service *AuthHService) Login(phone string) (*user.User, error) {
	u, err := service.UserRepository.Find(phone)
	if err != nil {
		u = &user.User{
			Phone: phone,
		}
	}
	sessid, err := secure.GenerateSecureID(16)
	if err != nil {
		return nil, fmt.Errorf("session generate error: %s", err.Error())
	}
	code, err := secure.GenerateSecureDigits(4)
	if err != nil {
		return nil, fmt.Errorf("code generate error: %s", err.Error())
	}
	u.SessionId = sessid
	u.ConfirmCode = code
	service.UserRepository.Database.DB.Save(&u)
	return u, nil
}

func (service *AuthHService) Confirm(SessionId string, ConfirmCode string) (string, error) {
	u, err := service.UserRepository.FindBySession(SessionId)
	if err != nil {
		return "", fmt.Errorf("Confirm error: %s", err.Error())
	}
	if u.ConfirmCode != ConfirmCode {
		return "", errors.New("confirm error")
	}
	token, err := service.JWT.Create(u.Phone)
	if err != nil {
		return "", fmt.Errorf("jwt error: %s", err.Error())
	}
	return token, nil
}
