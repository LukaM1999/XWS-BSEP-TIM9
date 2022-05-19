package startup

import (
	"context"
	cfg "dislinkt/api_gateway/startup/config"
	"dislinkt/common/auth"
	commentGw "dislinkt/common/proto/comment_service"
	connectionGw "dislinkt/common/proto/connection_service"
	postGw "dislinkt/common/proto/post_service"
	profileGw "dislinkt/common/proto/profile_service"
	reactionGw "dislinkt/common/proto/reaction_service"
	securityGw "dislinkt/common/proto/security_service"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
	"os"
	"strings"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(
		//runtime.WithForwardResponseOption(append),
		),
	}
	server.initHandlers()
	return server
}

func append(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return nil
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()

		// TODO: be careful!
		o := r.Header.Get("Origin")
		h.Set("Access-Control-Allow-Origin", o)

		if r.Method == http.MethodOptions {
			h.Set("Access-Control-Allow-Methods", strings.Join(
				[]string{
					http.MethodOptions,
					http.MethodGet,
					http.MethodPut,
					http.MethodHead,
					http.MethodPost,
					http.MethodDelete,
					http.MethodPatch,
					http.MethodTrace,
				}, ", ",
			))

			h.Set("Access-Control-Allow-Headers", strings.Join(
				[]string{
					"Access-Control-Allow-Headers",
					"Origin",
					"X-Requested-With",
					"Content-Type",
					"Accept",
					"Authorization",
				}, ", ",
			))

			return
		}

		next.ServeHTTP(w, r)
	})
}

func (server *Server) initHandlers() {
	tlsCredentials, err := auth.LoadTLSClientCredentials()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}
	securityEndpoint := fmt.Sprintf("%s:%s", server.config.SecurityHost, server.config.SecurityPort)
	err = securityGw.RegisterSecurityServiceHandlerFromEndpoint(context.TODO(), server.mux, securityEndpoint, opts)
	if err != nil {
		panic(err)
	}

	profileEndpoint := fmt.Sprintf("%s:%s", server.config.ProfileHost, server.config.ProfilePort)
	err = profileGw.RegisterProfileServiceHandlerFromEndpoint(context.TODO(), server.mux, profileEndpoint, opts)
	if err != nil {
		panic(err)
	}

	commentEndpoint := fmt.Sprintf("%s:%s", server.config.CommentHost, server.config.CommentPort)
	err = commentGw.RegisterCommentServiceHandlerFromEndpoint(context.TODO(), server.mux, commentEndpoint, opts)
	if err != nil {
		panic(err)
	}

	reactionEndpoint := fmt.Sprintf("%s:%s", server.config.ReactionHost, server.config.ReactionPort)
	err = reactionGw.RegisterReactionServiceHandlerFromEndpoint(context.TODO(), server.mux, reactionEndpoint, opts)
	if err != nil {
		panic(err)
	}

	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.ConnectionHost, server.config.ConnectionPort)
	err = connectionGw.RegisterConnectionServiceHandlerFromEndpoint(context.TODO(), server.mux, connectionEndpoint, opts)
	if err != nil {
		panic(err)
	}

	postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	err = postGw.RegisterPostServiceHandlerFromEndpoint(context.TODO(), server.mux, postEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) Start() {
	serverCertFile := getCertPath() + "cert/server-cert.pem"
	serverKeyFile := getCertPath() + "cert/server-key.pem"
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", server.config.Port),
		serverCertFile, serverKeyFile, cors(server.mux)))
}

func getCertPath() string {
	if os.Getenv("OS_ENV") != "docker" {
		return "../../"
	}
	return ""
}
