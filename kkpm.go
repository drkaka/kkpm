package kkpm

import "github.com/jackc/pgx"

// dbPool the pgx database pool.
var dbPool *pgx.ConnPool

// MessageInfo for private message.
type MessageInfo struct {
	MessageID string `json:"messageid" binding:"required"`
	Message   string `json:"message" binding:"required"`
	FromUser  int    `json:"fromuser" binding:"required"`
	ToUser    int    `json:"touser" binding:"required"`
	At        int    `json:"at" binding:"required"`
}
