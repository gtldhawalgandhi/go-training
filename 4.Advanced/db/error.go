package db

import "github.com/jackc/pgconn"

func PGErrorCode(err error) (string, bool) {
	pgerr, ok := err.(*pgconn.PgError)
	if !ok {
		return "", false
	}
	return pgerr.Code, true
}
