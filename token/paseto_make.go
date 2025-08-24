package token

import (
	"fmt"
	"time"
)

type PasetoMaker struct {
	secretKey string
}

func NewPasetoMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("")
	}
	return &JWTMaker{secretKey}, nil
}
func (maker PasetoMaker) CreateToken(username string, duration time.Duration) (string, *Payload) {
	return "", nil
}
func (maker PasetoMaker) VerifyToken(token string) (*Payload, error) {
	return nil, nil
}
