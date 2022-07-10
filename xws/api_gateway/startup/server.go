package startup

import (
	"context"
	"crypto/tls"
	cfg "dislinkt/api_gateway/startup/config"
	auth "dislinkt/common/domain"
	"dislinkt/common/loggers"
	commentGw "dislinkt/common/proto/comment_service"
	connectionGw "dislinkt/common/proto/connection_service"
	jobGw "dislinkt/common/proto/job_offer_service"
	postGw "dislinkt/common/proto/post_service"
	profileGw "dislinkt/common/proto/profile_service"
	reactionGw "dislinkt/common/proto/reaction_service"
	securityGw "dislinkt/common/proto/security_service"
	"dislinkt/common/tracer"
	_ "dislinkt/common/tracer"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	muxprom "gitlab.com/msvechla/mux-prometheus/pkg/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var log = loggers.NewGatewayLogger()

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
	tracer otgo.Tracer
	closer io.Closer
}

func (server *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewServer(config *cfg.Config) *Server {
	tracer, closer := tracer.Init("api-gateway")
	otgo.SetGlobalTracer(tracer)
	server := &Server{
		config: config,
		mux: runtime.NewServeMux(
			//runtime.WithForwardResponseOption(append),
			runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{&runtime.JSONPb{}}),
		),
		tracer: tracer,
		closer: closer,
	}
	server.initHandlers()
	return server
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()

		log.WithFields(logrus.Fields{
			"method":     r.Method,
			"url":        r.URL.String(),
			"origin":     r.Header.Get("Origin"),
			"user-agent": r.Header.Get("User-Agent"),
		}).Info("CORS filter")

		if r.Header.Get("Origin") != "" {
			h.Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		}
		//h.Set("Access-Control-Allow-Origin", "https://localhost:7777")

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
					"Location",
				}, ", ",
			))

			return
		}

		next.ServeHTTP(w, r)
	})
}

func prom(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/metrics" {
			promhttp.Handler().ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

var (
	totalReq = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dislinkt_total_req",
		Help: "The total number of requests",
	})
	successfulReq = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dislinkt_successful_req",
		Help: "The number of successful requests",
	})
	failedReq = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dislinkt_failed_req",
		Help: "The total number of failed requests",
	})
	notFoundReq = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "dislinkt_not_found_req",
		Help: "The total number of 404 requests with endpoint",
	}, []string{"code", "method"})
	visitor = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "dislinkt_visitor_req",
		Help: "Visitor from request",
	}, []string{"ip", "browser", "timestamp"})
	dataFlowFromReq = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dislinkt_data_flow_req",
		Help: "Data flow from request",
	})
)

//var (
//	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
//		Name: "dislinkt_http_duration_seconds",
//		Help: "Duration of HTTP requests.",
//	}, []string{"path"})
//)
//
//func prometheusMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		route := mux.CurrentRoute(r)
//		path, _ := route.GetPathTemplate()
//		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
//		next.ServeHTTP(w, r)
//		timer.ObserveDuration()
//	})
//}

func (server *Server) initHandlers() {
	//tlsCredentials, err := auth.LoadTLSClientCredentials()
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(config)),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor(
			grpc_opentracing.WithTracer(otgo.GlobalTracer()))),
	}
	securityEndpoint := fmt.Sprintf("%s:%s", server.config.SecurityHost, server.config.SecurityPort)
	err := securityGw.RegisterSecurityServiceHandlerFromEndpoint(context.TODO(), server.mux, securityEndpoint, opts)
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

	jobOfferEndpoint := fmt.Sprintf("%s:%s", server.config.JobOfferHost, server.config.JobOfferPort)
	err = jobGw.RegisterJobOfferServiceHandlerFromEndpoint(context.TODO(), server.mux, jobOfferEndpoint, opts)
	if err != nil {
		panic(err)
	}

	//err = server.mux.HandlePath("GET", "/metrics", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//	promhttp.Handler().ServeHTTP(w, r)
	//})
	//if err != nil {
	//	return
	//}
	//err = server.mux.HandlePath("GET", "/security/metrics", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//	promhttp.Handler().ServeHTTP(w, r)
	//})
	//if err != nil {
	//	return
	//}

	err = registerGatewayLogs(server)
	if err != nil {
		panic(err)
	}

	err = registerInterceptorLogs(server)
	if err != nil {
		panic(err)
	}
}

func registerGatewayLogs(server *Server) error {
	return server.mux.HandlePath("GET", "/logs", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Header().Set("Content-Type", "application/json")
		logPathPrefix := "../../logs/"
		if os.Getenv("OS_ENV") == "docker" {
			logPathPrefix = "./logs/"
		}
		resp := make(map[string][]auth.Log)
		content, err := os.ReadFile(logPathPrefix + "api_gateway/api_gateway.log")
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		lines := strings.Split(string(content), "\n")
		logs := make([]auth.Log, 0)
		for _, line := range lines {
			if line == "" {
				continue
			}
			var log auth.Log
			splitBySpace := strings.Split(line, " ")
			log.Time, err = time.Parse("2006-01-02T15:04:05.000Z", strings.Trim(strings.Split(splitBySpace[0], "=")[1], "\""))
			if err != nil {
				log.Time = time.Time{}
			}
			log.Level = strings.Split(splitBySpace[1], "=")[1]
			re := regexp.MustCompile(`msg="[/\\=!?'"\.a-zA-Z0-9_\s:-]*"`)
			msg := re.FindString(line)
			if msg != "" {
				log.Msg = strings.Trim(strings.Split(msg, "=")[1], "\"")
			}
			if msg == "" {
				re = regexp.MustCompile(`msg=[a-zA-Z]*`)
				msg = re.FindString(line)
				if msg != "" {
					log.Msg = strings.Split(msg, "=")[1]
				}
			}
			log.Service = "API gateway"
			log.FullContent = line
			logs = append(logs, log)
		}
		resp["logs"] = logs
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})
}

func registerInterceptorLogs(server *Server) error {
	return server.mux.HandlePath("GET", "/interceptor/logs", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Header().Set("Content-Type", "application/json")
		logPathPrefix := "../../logs/"
		if os.Getenv("OS_ENV") == "docker" {
			logPathPrefix = "./logs/"
		}
		resp := make(map[string][]auth.Log)
		content, err := os.ReadFile(logPathPrefix + "auth_interceptor/auth_interceptor.log")
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		lines := strings.Split(string(content), "\n")
		logs := make([]auth.Log, 0)
		for _, line := range lines {
			if line == "" {
				continue
			}
			var log auth.Log
			splitBySpace := strings.Split(line, " ")
			log.Time, err = time.Parse("2006-01-02T15:04:05.000Z", strings.Trim(strings.Split(splitBySpace[0], "=")[1], "\""))
			if err != nil {
				log.Time = time.Time{}
			}
			log.Level = strings.Split(splitBySpace[1], "=")[1]
			re := regexp.MustCompile(`msg="[/\\=!?'"\.a-zA-Z0-9_\s:-]*"`)
			msg := re.FindString(line)
			if msg != "" {
				log.Msg = strings.Trim(strings.Split(msg, "=")[1], "\"")
			}
			if msg == "" {
				re = regexp.MustCompile(`msg=[a-zA-Z]*`)
				msg = re.FindString(line)
				if msg != "" {
					log.Msg = strings.Split(msg, "=")[1]
				}
			}
			log.Service = "Auth interceptor"
			log.FullContent = line
			logs = append(logs, log)
		}
		resp["logs"] = logs
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})
}

func (server *Server) Start() {
	r := mux.NewRouter()
	instrumentation := muxprom.NewDefaultInstrumentation()
	r.Use(instrumentation.Middleware)
	r.Path("/metrics").Handler(promhttp.Handler())
	r.PathPrefix("/").Handler(cors(muxMiddleware(server)))
	serverCertFile := getCertPath() + "cert/server-cert.pem"
	serverKeyFile := getCertPath() + "cert/server-key.pem"
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", server.config.Port), serverCertFile, serverKeyFile, r))
}

var grpcGatewayTag = opentracing.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

func muxMiddleware(server *Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		endpointName := r.Method + " " + r.URL.Path

		parentSpanContext, err2 := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err2 == nil || err2 == opentracing.ErrSpanContextNotFound {
			serverSpan := opentracing.GlobalTracer().StartSpan(
				endpointName,
				ext.RPCServerOption(parentSpanContext),
				grpcGatewayTag,
			)
			r = r.WithContext(opentracing.ContextWithSpan(r.Context(), serverSpan))
			defer serverSpan.Finish()
		}
		lrw := negroni.NewResponseWriter(w)
		server.mux.ServeHTTP(lrw, r)

		statusCode := lrw.Status()
		ipAddr := r.RemoteAddr
		browser := r.UserAgent()
		t := time.Now()
		visitorLabel := prometheus.Labels{
			"ip":        ipAddr,
			"browser":   browser,
			"timestamp": t.Format("2006-01-02T15:04:05.000Z"),
		}
		visitor.With(visitorLabel).Inc()

		gb := r.ContentLength
		fmt.Println(gb)
		dataFlowFromReq.Add(float64(gb))
		fmt.Println(dataFlowFromReq)

		totalReq.Inc()
		if statusCode >= 200 && statusCode <= 399 {
			successfulReq.Inc()
		} else if statusCode >= 400 && statusCode <= 599 {
			if statusCode == 404 {
				labels := prometheus.Labels{
					"code":   "404",
					"method": endpointName,
				}
				notFoundReq.With(labels).Inc()
			}
			failedReq.Add(3)
		}
	})
}

func getCertPath() string {
	if os.Getenv("OS_ENV") != "docker" {
		return "../../"
	}
	return ""
}
