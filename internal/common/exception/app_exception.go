package exception

import "log"

type AppException struct {
	code int 
	msg  string
}

func newAppException(code int, msg string, debugMsg ...string) AppException {
	if len(debugMsg) != 0 {
		log.Printf("Error Occurred due to : %v", debugMsg[0])
	}
	return AppException{
		code: code,
		msg:  msg,
	}
}

var _ error = (*AppException)(nil)

func (e AppException) GetErrorCode() int {
	return e.code
}
func (e AppException) Error() string {
	return e.msg
}
