package auth_ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	service "github.com/prcryx/monster_mall/internal/apis/auth/auth_srv"
	"github.com/prcryx/monster_mall/internal/common/exception"
	"github.com/prcryx/monster_mall/internal/common/utils"
	"github.com/prcryx/monster_mall/internal/models"
	"gorm.io/gorm"
)

type IAuthController interface {
	SendOtp(http.ResponseWriter, *http.Request)
}

type AuthController struct {
	service  *service.AuthService
	validate *validator.Validate
}

var _ IAuthController = (*AuthController)(nil)

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		service:  service.NewAuthService(db),
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (ac *AuthController) SendOtp(w http.ResponseWriter, request *http.Request) {

	//validating request
	otpReq := new(models.OtpReqBody)
	err := json.NewDecoder(request.Body).Decode(&otpReq)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, exception.InternalServerError().Error())
		return
	}
	//checking the dto
	if err := ac.validate.Struct(otpReq); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, exception.InvalidBodyRequestException().Error())
		return
	}

	verificationRes, err := ac.service.SendOtp(otpReq)
	if err != nil {
		if appEx, ok := err.(exception.AppException); ok {
			utils.ResponseWithError(w, appEx.GetErrorCode(), appEx.Error())
			return
		}
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJSONData(w, http.StatusOK, verificationRes)

}
