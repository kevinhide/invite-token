package models

import "time"

type InviteToken struct {
	ID        string    `json:"id"`
	Active    bool      `json:"active"`
	ExpiresAt time.Time `json:"expires_at"`
}

type TokenReq struct {
	Token string `json:"token"`
}
