package model

import "time"

type AccountDTO struct {
	Owner     string    `json:"owner" binding:"required"`
	Currency  string    `json:"currency" binding:"required,currency"`
	Balance   int64     `json:"balance" binding:"required,min=0"`
	CreatedAt time.Time `json:"created_at" binding:"required,datetime"`
}
