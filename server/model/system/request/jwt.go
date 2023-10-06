package request

import (
	"github.com/gofrs/uuid/v5"
	jwt "github.com/golang-jwt/jwt/v4"
)

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}
