package exception

import (
	"net/http"
)

func EmptyError() AppException {
	return newAppException(0, "")
}
func InternalServerError(debugMsg ...string) AppException {
	return newAppException(http.StatusInternalServerError, InternalServerErrorMsg, debugMsg...)
}

func InvalidBodyRequestException() AppException {
	return newAppException(http.StatusBadRequest, InvalidReqBodyMsg)
}

func UnacceptedSigningError() AppException {
	return newAppException(http.StatusForbidden, SigningErrorMsg)
}

func TooShortPasswordException() AppException {
	return newAppException(http.StatusBadRequest, TooShortPasswordMsg)
}
func SpecialCharRequiredInPasswordException() AppException {
	return newAppException(http.StatusBadRequest, RequireSpecialCharInPasswordMsg)
}
func NumRequiredInPasswordException() AppException {
	return newAppException(http.StatusBadRequest, RequiredNumInPasswordMsg)
}
