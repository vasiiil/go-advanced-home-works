package auth

import "errors"

var ErrWrongVerificationCode = errors.New("wrong verification code")
var ErrInvalidVerificationCode = errors.New("invalid verification code")
