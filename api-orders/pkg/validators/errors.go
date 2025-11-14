package validators

import "errors"

var ErrEmptyPhone = errors.New("empty phone")
var ErrInvalidPhoneFormat = errors.New("invalid phone format")