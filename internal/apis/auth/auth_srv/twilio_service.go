package auth_srv

import (
	"github.com/prcryx/monster_mall/internal/models"
)

type ITwilioApp interface {
	SendOtp(string) (*models.OtpResBody, error)
}

type TwilioApp struct {
}

func NewTwilioApp() *TwilioApp {
	// create twilio app
	return &TwilioApp{}
}

var _ ITwilioApp = (*TwilioApp)(nil)

func (app *TwilioApp) SendOtp(receiver string) (*models.OtpResBody, error) {
	return &models.OtpResBody{}, nil
}
