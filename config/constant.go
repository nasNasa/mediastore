package config

import "time"

const (
	JwtSignKey                 = "jwt_sign_key"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
	AuthMiddlewareContextKey   = "claims"
)
