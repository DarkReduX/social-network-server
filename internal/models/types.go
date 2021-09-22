package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Profile struct {
	UUID          string     `json:"-" example:"1"`
	Username      string     `json:"username" example:"darkredux"`
	Password      string     `json:"password" example:"2"`
	AvatarLink    *string    `json:"avatar_link" example:"https://images/2.com"`
	LastActivity  *time.Time `json:"last_activity" example:"2021-09-14T17:45:31.025716Z"`
	CreatedAt     *time.Time `json:"created_at" example:"2021-09-14T17:45:31.025716Z"`
	CreatedFromIp *string    `json:"-" example:"192.168.255.255"`
	DeletedAt     *time.Time `json:"deleted_at" example:"2021-09-14T17:45:31.025716Z"`
	IsActivate    bool       `json:"is_activate" example:"false"`
}

type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
