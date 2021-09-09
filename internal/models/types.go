package models

import (
	"time"
)

type Profile struct {
	Id            string    `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	AvatarLink    string    `json:"avatar_link"`
	LastActivity  time.Time `json:"last_activity"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedFromIp string    `json:"created_from_ip"`
	DeletedAt     time.Time `json:"deleted_at"`
	IsActivate    bool      `json:"is_activate"`
}
