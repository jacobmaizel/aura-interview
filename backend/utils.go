package main

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

func PgxErrHelper(err error) (string, string) {
	var pgErr *pgconn.PgError
	fmt.Println("error:", err.Error())
	if errors.As(err, &pgErr) {
		return pgErr.Code, pgErr.Message

	}
	return "", ""
}
