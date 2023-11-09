package models

import "github.com/prcryx/monster_mall/internal/common/types"

type OtpReqBody struct {
	PhoneNo string `json:"phoneNo" validate:"required,e164"`
}

type OtpResBody struct {
	PhoneNo            string       `json:"phoneNo"`
	VerificationStatus types.Status `json:"verificationStatus"`
}

type OtpVerificationReqBody struct {
	PhoneNo string `json:"phoneNo"`
	Otp     string `json:"otp"`
}
type OtpVerificationResBody struct {
	PhoneNo            string       `json:"phoneNo"`
	VerificationStatus types.Status `json:"status"`
}

type AuthenticationResponse struct {
	UID         uint   `json:"uid"`
	AccessToken string `json:"accessToken"`
	PhoneNo     string `json:"phoneNo"`
	CreatedAt   int64  `json:"createdAt"`
}

type TokenRefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}
