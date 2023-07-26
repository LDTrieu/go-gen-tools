package middleware

import (
	"go-gen-tools/internal/repositories"
	"net/http"
)

type Middleware func(next http.Handler) http.Handler

func Authentication(userRepo repositories.UserRepository)
