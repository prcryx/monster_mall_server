package exception

import (
	"net/http"
)

func EmptyError() AppException {
	return newAppException(0, "")
}
func InternalServerError(debugMsg ...string) AppException {
	return newAppException(http.StatusInternalServerError, "internal server error", debugMsg...)
}

func InvalidBodyRequestException() AppException {
	return newAppException(http.StatusBadRequest, "invalid body request")
}

func UnacceptedSigningError() AppException {
	return newAppException(http.StatusForbidden, "signing error")
}
