package types

import (
	"net/url"
	"regexp"

	"github.com/prcryx/monster_mall/internal/common/exception"
)

type PasswordConfig struct {
	PassStr       string
	IsUrlEncoding bool
	MinLength     int
}

type Password string

type IPassWord interface {
	ToString() string
}

var _ IPassWord = (*Password)(nil)

func NewPassword(passConfig PasswordConfig) (Password, error) {
	var minLen int = 0
	if passConfig.MinLength > 0 {
		minLen = passConfig.MinLength
	}
	s := passConfig.PassStr
	hasSpecialChar := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString
	hasNum := regexp.MustCompile(`\d`).MatchString

	// Check constraints
	if len(s) < minLen {
		return "", exception.TooShortPasswordException()
	}
	if !hasSpecialChar(s) {
		return "", exception.SpecialCharRequiredInPasswordException()

	}
	if !hasNum(s) {
		return "", exception.NumRequiredInPasswordException()
	}

	// If constraints are met, create the Password
	if passConfig.IsUrlEncoding {
		return Password(url.QueryEscape(s)), nil
	}
	return Password(s), nil
}

func (p Password) ToString() string {
	return string(p)
}
