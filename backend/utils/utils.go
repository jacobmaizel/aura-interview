

import (
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
)

func PgxErrHelper(err error) (string, string) {
	var pgErr *pgconn.PgError
	fmt.Println("error:", err.Error())
	if errors.As(err, &pgErr) {
		return pgErr.Code, pgErr.Message

	}
	return "", ""
}
