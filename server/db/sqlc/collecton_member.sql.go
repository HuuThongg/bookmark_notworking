// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: collecton_member.sql

package sqlc

import (
	"context"
)

const addNewCollectionMember = `-- name: AddNewCollectionMember :one
INSERT INTO collection_member (collection_id, member_id, collection_access_level) VALUES ($1, $2, $3) RETURNING collection_id, member_id, join_date, collection_access_level
`

type AddNewCollectionMemberParams struct {
	CollectionID          string                `json:"collection_id"`
	MemberID              int64                 `json:"member_id"`
	CollectionAccessLevel CollectionAccessLevel `json:"collection_access_level"`
}

func (q *Queries) AddNewCollectionMember(ctx context.Context, arg AddNewCollectionMemberParams) (CollectionMember, error) {
	row := q.db.QueryRowContext(ctx, addNewCollectionMember, arg.CollectionID, arg.MemberID, arg.CollectionAccessLevel)
	var i CollectionMember
	err := row.Scan(
		&i.CollectionID,
		&i.MemberID,
		&i.JoinDate,
		&i.CollectionAccessLevel,
	)
	return i, err
}

const checkIfCollectionMemberWithCollectionAndMemberIdsExists = `-- name: CheckIfCollectionMemberWithCollectionAndMemberIdsExists :one
SELECT EXISTS (SELECT collection_id, member_id, join_date, collection_access_level FROM collection_member WHERE collection_id = $1 AND member_id = $2 LIMIT 1)
`

type CheckIfCollectionMemberWithCollectionAndMemberIdsExistsParams struct {
	CollectionID string `json:"collection_id"`
	MemberID     int64  `json:"member_id"`
}

func (q *Queries) CheckIfCollectionMemberWithCollectionAndMemberIdsExists(ctx context.Context, arg CheckIfCollectionMemberWithCollectionAndMemberIdsExistsParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkIfCollectionMemberWithCollectionAndMemberIdsExists, arg.CollectionID, arg.MemberID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getCollectionMemberByCollectionAndMemberIDs = `-- name: GetCollectionMemberByCollectionAndMemberIDs :one
SELECT collection_id, member_id, join_date, collection_access_level FROM collection_member WHERE collection_id = $1 AND member_id = $2 LIMIT 1
`

type GetCollectionMemberByCollectionAndMemberIDsParams struct {
	CollectionID string `json:"collection_id"`
	MemberID     int64  `json:"member_id"`
}

func (q *Queries) GetCollectionMemberByCollectionAndMemberIDs(ctx context.Context, arg GetCollectionMemberByCollectionAndMemberIDsParams) (CollectionMember, error) {
	row := q.db.QueryRowContext(ctx, getCollectionMemberByCollectionAndMemberIDs, arg.CollectionID, arg.MemberID)
	var i CollectionMember
	err := row.Scan(
		&i.CollectionID,
		&i.MemberID,
		&i.JoinDate,
		&i.CollectionAccessLevel,
	)
	return i, err
}

const getcollectionsSharedWithUser = `-- name: GetcollectionsSharedWithUser :many
SELECT collection_id, member_id, join_date, collection_access_level FROM  collection_member WHERE member_id = $1
`

func (q *Queries) GetcollectionsSharedWithUser(ctx context.Context, memberID int64) ([]CollectionMember, error) {
	rows, err := q.db.QueryContext(ctx, getcollectionsSharedWithUser, memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CollectionMember
	for rows.Next() {
		var i CollectionMember
		if err := rows.Scan(
			&i.CollectionID,
			&i.MemberID,
			&i.JoinDate,
			&i.CollectionAccessLevel,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
