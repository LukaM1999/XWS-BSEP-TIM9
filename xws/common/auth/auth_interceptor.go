package auth

import (
	"context"
	auth "dislinkt/common/domain"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor is a server interceptor for authentication and authorization
type AuthInterceptor struct {
	jwtManager      *JWTManager
	accessibleRoles map[string]string
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(jwtManager *JWTManager, accessibleRoles map[string]string) *AuthInterceptor {
	//log.SetLevel(log.InfoLevel)
	//log.SetReportCaller(true)
	//multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
	//	Filename:   "../../logs/xws.log",
	//	MaxSize:    1,
	//	MaxBackups: 3,
	//	MaxAge:     28,
	//	Compress:   true,
	//})
	//log.SetOutput(multiWriter)
	return &AuthInterceptor{jwtManager, accessibleRoles}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		err, userId, username := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, "userId", userId)
		ctx = context.WithValue(ctx, "username", username)
		return handler(ctx, req)
	}
}

// Stream returns a server interceptor function to authenticate and authorize stream RPC
func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)

		err, _, _ := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func getSecurityDatabaseClient(host, port string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s/", host, port)
	options := options.Client().ApplyURI(uri)
	return mongo.Connect(context.TODO(), options)
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) (error, string, string) {
	logger := log.WithFields(log.Fields{
		"method": method,
	})
	logger.Info("Authorizing")
	permission, ok := interceptor.accessibleRoles[method]
	if !ok {
		// everyone can access
		return nil, "", ""
	}

	securityClient, err := getSecurityDatabaseClient("localhost", "27017")
	if err != nil {
		logger.Error("Error getting security database client")
		return status.Errorf(codes.Internal, "could not connect to security database: %v", err), "", ""
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Error("Error getting metadata from incoming context")
		return status.Errorf(codes.Unauthenticated, "metadata is not provided"), "", ""
	}

	values := md["authorization"]
	if len(values) == 0 {
		logger.Error("Error getting authorization from metadata")
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided"), "", ""
	}

	accessToken := values[0]
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		logger.Error("Error verifying access token")
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err), "", ""
	}

	filter := bson.D{{"role", claims.Role}}
	rolePermission := &auth.RolePermission{}
	err = securityClient.Database("security_service").Collection("rolePermission").FindOne(context.TODO(), filter).Decode(rolePermission)
	if err != nil {
		logger.WithField("role", claims.Role).Error("Error getting role permissions")
		return status.Errorf(codes.Internal, "could not find role permissions: %v", err), "", ""
	}
	for _, p := range rolePermission.Permissions {
		if p == permission || p == "*" {
			logger.Info("Authorized")
			return nil, claims.UserId, claims.Username
		}
	}
	logger.WithField("role", claims.Role).Error("Error authorizing")
	return status.Error(codes.PermissionDenied, "no permission to access this RPC"), "", ""
}
