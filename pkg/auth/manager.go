package auth

import "errors"

type TokenManager interface {
}

type manager struct {
	signinKey string
}

func NewManager(signinKey string) (TokenManager, error) {
	if signinKey == "" {
		return nil, errors.New("emoty signing key")
	}
	return &manager{
		signinKey: signinKey,
	}, nil
}
