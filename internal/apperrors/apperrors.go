package apperrors

import "errors"

var ErrWrongTTL = errors.New("ttl must not be 0")
var ErrWrongPort = errors.New("wrong server port value")
var ErrWrongCleanFreq = errors.New("wrong server port value")
