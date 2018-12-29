package interceptor

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/gedorinku/tsugidoko-server/app/di"
	"github.com/gedorinku/tsugidoko-server/infra/record"
)

const (
	// AuthorizationKey is metadata key
	AuthorizationKey = "authorization"
	// SessionAuthorizationType is the type in authorization header
	SessionAuthorizationType = "session"
)

var (
	// ErrMetadataNotFound is returned when metadata not found in context
	ErrMetadataNotFound = errors.New("metadata not found in context")
	// ErrInvalidAuthorizationMetadata is returned when authorization metadata is invalid
	ErrInvalidAuthorizationMetadata = status.Error(codes.InvalidArgument, "Invalid authorization metadata")
)

type currentSessionKey struct{}

// GetCurrentSession returns the current session from context
func GetCurrentSession(ctx context.Context) (s *record.Session, ok bool) {
	v := ctx.Value(currentSessionKey{})
	s, ok = v.(*record.Session)
	return
}

// Authorizator provide the authorization interceptor
type Authorizator struct {
	di.StoreComponent
}

// NewAuthorizator returns new Authorizator
func NewAuthorizator(store di.StoreComponent) *Authorizator {
	return &Authorizator{
		StoreComponent: store,
	}
}

// UnaryServerInterceptor returns authorization UnaryServerInterceptor
func (a *Authorizator) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return a.authorization
}

func (a *Authorizator) authorization(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	d, err := fromMeta(ctx, AuthorizationKey)
	if err != nil {
		grpclog.Error(err)
		return handler(ctx, req)
	}

	if strings.Index(d, SessionAuthorizationType) != 0 {
		return nil, ErrInvalidAuthorizationMetadata
	}

	secret := strings.TrimSpace(d[len(SessionAuthorizationType):])
	s, err := a.SessionStore(ctx).GetSession(secret)
	if err != nil {
		grpclog.Error(err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	newCtx := context.WithValue(ctx, currentSessionKey{}, s)
	return handler(newCtx, req)
}

func fromMeta(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.WithStack(ErrMetadataNotFound)
	}

	h := md.Get(key)
	if len(h) == 0 {
		return "", errors.WithStack(errors.Errorf("metadata not found: key = %s", key))
	}
	return h[0], nil
}
