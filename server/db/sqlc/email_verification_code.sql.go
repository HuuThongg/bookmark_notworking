// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: email_verification_code.sql

package sqlc

import (
	"context"
	"time"
)

const deleteEmailVerificationCode = `-- name: DeleteEmailVerificationCode :exec
DELETE FROM email_verification WHERE email = $1
`

func (q *Queries) DeleteEmailVerificationCode(ctx context.Context, email string) error {
	_, err := q.db.ExecContext(ctx, deleteEmailVerificationCode, email)
	return err
}

const getOtp = `-- name: GetOtp :one
SELECT code, email, expiry FROM email_verification WHERE email = $1 LIMIT 1
`

func (q *Queries) GetOtp(ctx context.Context, email string) (EmailVerification, error) {
	row := q.db.QueryRowContext(ctx, getOtp, email)
	var i EmailVerification
	err := row.Scan(&i.Code, &i.Email, &i.Expiry)
	return i, err
}

const newEmailVerificationCode = `-- name: NewEmailVerificationCode :one
INSERT INTO email_verification (code, email, expiry)
VALUES ($1, $2, $3)
ON CONFLICT (email) DO UPDATE SET code = EXCLUDED.code, expiry = EXCLUDED.expiry
RETURNING code, email, expiry
`

type NewEmailVerificationCodeParams struct {
	Code   string    `json:"code"`
	Email  string    `json:"email"`
	Expiry time.Time `json:"expiry"`
}

func (q *Queries) NewEmailVerificationCode(ctx context.Context, arg NewEmailVerificationCodeParams) (EmailVerification, error) {
	row := q.db.QueryRowContext(ctx, newEmailVerificationCode, arg.Code, arg.Email, arg.Expiry)
	var i EmailVerification
	err := row.Scan(&i.Code, &i.Email, &i.Expiry)
	return i, err
}
