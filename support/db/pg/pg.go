package pg

import (
	"github.com/lib/pq"
	"github.com/AnneNamuli/go-stellar/support/errors"
)

func IsUniqueViolation(err error) bool {
	switch pgerr := errors.Cause(err).(type) {
	case *pq.Error:
		return string(pgerr.Code) == "23505"
	default:
		return false
	}
}
