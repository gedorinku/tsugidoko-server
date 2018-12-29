package server

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrInvalidSession is returned when session is invalid
var ErrInvalidSession = status.Error(codes.Unauthenticated, "Unauthenticated")
