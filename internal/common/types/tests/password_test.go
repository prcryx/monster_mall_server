package tests

import (
	"regexp"
	"testing"

	"github.com/prcryx/monster_mall/internal/common/exception"
	"github.com/prcryx/monster_mall/internal/common/types"
)

func TestUrlEncodedPassword(t *testing.T) {
	password1 := "Password@12"
	want := regexp.MustCompile("Password%4012")

	passPhase, err := types.NewPassword(types.PasswordConfig{
		PassStr:       password1,
		IsUrlEncoding: true,
		MinLength:     7,
	})

	if !want.MatchString(passPhase.ToString()) || err != nil {
		t.Fatalf(`NewPassword(PassConfig{}) = %q, %v, want match for %#q, nil`, passPhase.ToString(), err, want)
	}

}

func TestNonEmptyPassword(t *testing.T) {
	password := ""
	passConf := types.PasswordConfig{
		PassStr:       password,
		IsUrlEncoding: true,
		MinLength:     8,
	}

	want := exception.TooShortPasswordException()

	passPhase, err := types.NewPassword(passConf)
	if err == nil || err != want {
		t.Fatalf(`TestEmptyUrlEncodedPassword(passConf) = %v, %v, want err = %v`, passPhase, err, want)
	}

}
func TestSpecialCharRequiredPassword(t *testing.T) {
	password := "earth616"
	passConf := types.PasswordConfig{
		PassStr:       password,
		IsUrlEncoding: true,
		MinLength:     6,
	}

	want := exception.SpecialCharRequiredInPasswordException()

	passPhase, err := types.NewPassword(passConf)
	if err == nil || err != want {
		t.Fatalf(`TestEmptyUrlEncodedPassword(passConf) = %v, %v, want err = %v`, passPhase, err, want)
	}

}
