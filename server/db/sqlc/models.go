// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type AccessLevel string

const (
	AccessLevelView  AccessLevel = "view"
	AccessLevelEdit  AccessLevel = "edit"
	AccessLevelAdmin AccessLevel = "admin"
)

func (e *AccessLevel) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AccessLevel(s)
	case string:
		*e = AccessLevel(s)
	default:
		return fmt.Errorf("unsupported scan type for AccessLevel: %T", src)
	}
	return nil
}

type NullAccessLevel struct {
	AccessLevel AccessLevel `json:"access_level"`
	Valid       bool        `json:"valid"` // Valid is true if AccessLevel is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAccessLevel) Scan(value interface{}) error {
	if value == nil {
		ns.AccessLevel, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AccessLevel.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAccessLevel) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AccessLevel), nil
}

type CollectionAccessLevel string

const (
	CollectionAccessLevelView  CollectionAccessLevel = "view"
	CollectionAccessLevelEdit  CollectionAccessLevel = "edit"
	CollectionAccessLevelAdmin CollectionAccessLevel = "admin"
)

func (e *CollectionAccessLevel) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CollectionAccessLevel(s)
	case string:
		*e = CollectionAccessLevel(s)
	default:
		return fmt.Errorf("unsupported scan type for CollectionAccessLevel: %T", src)
	}
	return nil
}

type NullCollectionAccessLevel struct {
	CollectionAccessLevel CollectionAccessLevel `json:"collection_access_level"`
	Valid                 bool                  `json:"valid"` // Valid is true if CollectionAccessLevel is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCollectionAccessLevel) Scan(value interface{}) error {
	if value == nil {
		ns.CollectionAccessLevel, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CollectionAccessLevel.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCollectionAccessLevel) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CollectionAccessLevel), nil
}

type Account struct {
	ID              int64          `json:"id"`
	Fullname        string         `json:"fullname"`
	Email           string         `json:"email"`
	EmailVerified   bool           `json:"email_verified"`
	Picture         string         `json:"picture"`
	AccountPassword string         `json:"account_password"`
	CreateAt        time.Time      `json:"create_at"`
	Intention       sql.NullString `json:"intention"`
	LastLogin       time.Time      `json:"last_login"`
}

type AccountSession struct {
	RefreshTokenID string    `json:"refresh_token_id"`
	AccountID      int64     `json:"account_id"`
	IssuedAt       time.Time `json:"issued_at"`
	Expiry         time.Time `json:"expiry"`
	UserAgent      string    `json:"user_agent"`
	ClientIp       string    `json:"client_ip"`
}

type CollectionMember struct {
	CollectionID          string                `json:"collection_id"`
	MemberID              int64                 `json:"member_id"`
	JoinDate              time.Time             `json:"join_date"`
	CollectionAccessLevel CollectionAccessLevel `json:"collection_access_level"`
}

type Contact struct {
	ID          int64  `json:"id"`
	Account     int64  `json:"account"`
	MessageBody string `json:"message_body"`
}

type EmailVerification struct {
	Code   string    `json:"code"`
	Email  string    `json:"email"`
	Expiry time.Time `json:"expiry"`
}

type Folder struct {
	FolderID               string         `json:"folder_id"`
	AccountID              int64          `json:"account_id"`
	FolderName             string         `json:"folder_name"`
	Path                   string         `json:"path"`
	Label                  string         `json:"label"`
	Starred                bool           `json:"starred"`
	FolderCreatedAt        time.Time      `json:"folder_created_at"`
	FolderUpdatedAt        time.Time      `json:"folder_updated_at"`
	SubfolderOf            sql.NullString `json:"subfolder_of"`
	FolderDeletedAt        sql.NullTime   `json:"folder_deleted_at"`
	TextsearchableIndexCol interface{}    `json:"textsearchable_index_col"`
}

type Link struct {
	LinkID                 string         `json:"link_id"`
	LinkTitle              string         `json:"link_title"`
	LinkThumbnail          string         `json:"link_thumbnail"`
	LinkFavicon            string         `json:"link_favicon"`
	LinkHostname           string         `json:"link_hostname"`
	LinkUrl                string         `json:"link_url"`
	LinkNotes              string         `json:"link_notes"`
	AccountID              int64          `json:"account_id"`
	FolderID               sql.NullString `json:"folder_id"`
	AddedAt                time.Time      `json:"added_at"`
	UpdatedAt              time.Time      `json:"updated_at"`
	DeletedAt              sql.NullTime   `json:"deleted_at"`
	TextsearchableIndexCol interface{}    `json:"textsearchable_index_col"`
}

type MemberInvite struct {
	InviteID                int64       `json:"invite_id"`
	InviteToken             string      `json:"invite_token"`
	SharedCollectionID      string      `json:"shared_collection_id"`
	CollectionSharedByName  string      `json:"collection_shared_by_name"`
	CollectionSharedByEmail string      `json:"collection_shared_by_email"`
	CollectionSharedWith    string      `json:"collection_shared_with"`
	InviteExpiry            time.Time   `json:"invite_expiry"`
	MemberAccessLevel       AccessLevel `json:"member_access_level"`
}

type PasswordResetToken struct {
	ID          sql.NullInt64 `json:"id"`
	AccountID   int64         `json:"account_id"`
	TokenHash   string        `json:"token_hash"`
	TokenExpiry time.Time     `json:"token_expiry"`
}

type PublicSharedCollection struct {
	CollectionID          string                `json:"collection_id"`
	CollectionPassward    string                `json:"collection_passward"`
	CollectionSharedBy    int64                 `json:"collection_shared_by"`
	CollectionShareAt     time.Time             `json:"collection_share_at"`
	CollectionShareExpiry sql.NullTime          `json:"collection_share_expiry"`
	CollectionAccessLevel CollectionAccessLevel `json:"collection_access_level"`
}
