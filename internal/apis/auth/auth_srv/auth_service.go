package auth_srv

import (
	"github.com/prcryx/monster_mall/internal/common/types"
	"github.com/prcryx/monster_mall/internal/models"
	"gorm.io/gorm"
)

type AuthService struct {
	db        *gorm.DB
	twilioApp *TwilioApp
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db:        db,
		twilioApp: NewTwilioApp(),
	}
}

func (authService *AuthService) SendOtp(otpReq *models.OtpReqBody) (*models.OtpResBody, error) {
	return &models.OtpResBody{
		PhoneNo:            otpReq.PhoneNo,
		VerificationStatus: types.Pending,
	}, nil
}
