package email

import (
	"github.com/dstgo/tracker/internal/types/errs"
)

var (
	ErrCodeExpired = errs.NewI18nError("email.code.expired").FallBack("email code exceeded")
	ErrSendFailed  = errs.NewI18nError("email.send.failed").FallBack("email send failed")
)
