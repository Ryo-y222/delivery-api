package service

import "errors"

var ErrInvalidCredentials = errors.New("invalid credentials")

var ErrEmailAlreadyExists = errors.New("Email already exists")
