// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: invite.sql

package sqlc

import (
	"context"
	"time"
)

const createInvite = `-- name: CreateInvite :one
INSERT INTO member_invite (shared_collection_id, collection_shared_by_name, collection_shared_by_email, collection_shared_with, member_access_level, invite_expiry, invite_token)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT ON CONSTRAINT member_invite_shared_collection_id_collection_shared_with_key DO UPDATE SET  collection_shared_by_name = EXCLUDED.collection_shared_by_name, collection_shared_by_email = EXCLUDED.collection_shared_by_email, member_access_level = EXCLUDED.member_access_level, invite_expiry = EXCLUDED.invite_expiry, invite_token = EXCLUDED.invite_token 
RETURNING invite_id, invite_token, shared_collection_id, collection_shared_by_name, collection_shared_by_email, collection_shared_with, invite_expiry, member_access_level
`

type CreateInviteParams struct {
	SharedCollectionID      string      `json:"shared_collection_id"`
	CollectionSharedByName  string      `json:"collection_shared_by_name"`
	CollectionSharedByEmail string      `json:"collection_shared_by_email"`
	CollectionSharedWith    string      `json:"collection_shared_with"`
	MemberAccessLevel       AccessLevel `json:"member_access_level"`
	InviteExpiry            time.Time   `json:"invite_expiry"`
	InviteToken             string      `json:"invite_token"`
}

func (q *Queries) CreateInvite(ctx context.Context, arg CreateInviteParams) (MemberInvite, error) {
	row := q.db.QueryRowContext(ctx, createInvite,
		arg.SharedCollectionID,
		arg.CollectionSharedByName,
		arg.CollectionSharedByEmail,
		arg.CollectionSharedWith,
		arg.MemberAccessLevel,
		arg.InviteExpiry,
		arg.InviteToken,
	)
	var i MemberInvite
	err := row.Scan(
		&i.InviteID,
		&i.InviteToken,
		&i.SharedCollectionID,
		&i.CollectionSharedByName,
		&i.CollectionSharedByEmail,
		&i.CollectionSharedWith,
		&i.InviteExpiry,
		&i.MemberAccessLevel,
	)
	return i, err
}

const deleteInvite = `-- name: DeleteInvite :exec
DELETE FROM member_invite WHERE invite_token = $1
`

func (q *Queries) DeleteInvite(ctx context.Context, inviteToken string) error {
	_, err := q.db.ExecContext(ctx, deleteInvite, inviteToken)
	return err
}

const getInviteByToken = `-- name: GetInviteByToken :one
SELECT invite_id, invite_token, shared_collection_id, collection_shared_by_name, collection_shared_by_email, collection_shared_with, invite_expiry, member_access_level FROM member_invite WHERE invite_token = $1 LIMIT 1
`

func (q *Queries) GetInviteByToken(ctx context.Context, inviteToken string) (MemberInvite, error) {
	row := q.db.QueryRowContext(ctx, getInviteByToken, inviteToken)
	var i MemberInvite
	err := row.Scan(
		&i.InviteID,
		&i.InviteToken,
		&i.SharedCollectionID,
		&i.CollectionSharedByName,
		&i.CollectionSharedByEmail,
		&i.CollectionSharedWith,
		&i.InviteExpiry,
		&i.MemberAccessLevel,
	)
	return i, err
}
