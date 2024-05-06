package http_msg

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/errmsg"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
	"net/http"
)

func Error(err error) (message string, code int) {
	switch err.(type) {
	case richerror.RichError:
		re := err.(richerror.RichError)
		msg := re.Message()

		code := mapKindToHTTPStatusCode(re.Kind())

		// we should not expose unexpected error messages
		if code >= 500 {
			msg = errmsg.ErrorMsgSomethingWentWrong
		}

		return msg, code
	default:
		return err.Error(), http.StatusBadRequest
	}
}

func mapKindToHTTPStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.Invalid:
		return http.StatusUnprocessableEntity
	case richerror.NotFound:
		return http.StatusNotFound
	case richerror.Forbidden:
		return http.StatusForbidden
	case richerror.Unexpected:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
