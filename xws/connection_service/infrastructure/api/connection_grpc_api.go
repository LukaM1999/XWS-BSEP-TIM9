package api

import (
	"context"
	"dislinkt/common/loggers"
	pb "dislinkt/common/proto/connection_service"
	pbPost "dislinkt/common/proto/post_service"
	"dislinkt/common/tracer"
	"dislinkt/connection_service/application"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

var log = loggers.NewConnectionLogger()

type ConnectionHandler struct {
	pb.UnimplementedConnectionServiceServer
	service    *application.ConnectionService
	postClient pbPost.PostServiceClient
}

func NewConnectionHandler(service *application.ConnectionService, postClient pbPost.PostServiceClient) *ConnectionHandler {
	return &ConnectionHandler{
		service:    service,
		postClient: postClient,
	}
}

func (handler *ConnectionHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "Get Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	Connections, err := handler.service.Get(ctx, request.UserId)
	if err != nil {
		log.WithField("userId", request.UserId).Errorf("Cannot get connections: %v", err)
		return nil, err
	}
	response := &pb.GetResponse{
		Connections: []*pb.Connection{},
	}
	for _, Connection := range Connections {
		current := mapConnectionToPb(Connection)
		response.Connections = append(response.Connections, current)
	}
	return response, nil
}

func (handler *ConnectionHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "Create Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	newConnection, err := handler.service.Create(ctx, request.Connection.IssuerId, request.Connection.SubjectId)
	if err != nil {
		log.Errorf("Cannot create connection: %v", err)
		return nil, err
	}
	log.Info("Connection created")
	return &pb.CreateResponse{
		Connection: mapConnectionToPb(newConnection),
	}, nil
}

func (handler *ConnectionHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "Delete Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	id, err := strconv.Atoi(request.Id)
	if err != nil {
		log.WithField("connectionId", request.Id).Errorf("Cannot convert connection id to int: %v", err)
		return nil, err
	}
	err = handler.service.Delete(ctx, id)
	if err != nil {
		log.Errorf("Cannot delete connection: %v", err)
		return nil, err
	}
	handler.postClient.DeleteConnection(ctx, &pbPost.DeleteRequest{Id: request.Id})
	log.Info("Connection deleted")
	return &pb.DeleteResponse{}, nil
}

func (handler *ConnectionHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "Update Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	id, err := strconv.Atoi(request.Id)
	if err != nil {
		log.WithField("connectionId", request.Id).Errorf("Cannot convert connection id to int: %v", err)
		return nil, err
	}
	connection, err := handler.service.UpdateConnection(ctx, id)
	if err != nil {
		log.WithField("connectionId", request.Id).Errorf("Cannot update connection: %v", err)
		return nil, err
	}
	if connection.IsApproved {
		_, err = handler.postClient.CreateConnection(ctx, &pbPost.CreateConnectionRequest{
			Connection: mapConnectionToPostConnectionPb(connection),
		})
		if err != nil {
			log.Errorf("Cannot create connection: %v", err)
			return nil, err
		}
	}
	log.WithField("connectionId", request.Id).Infof("Connection updated")
	return &pb.UpdateResponse{
		Connection: mapConnectionToPb(connection),
	}, nil
}

func (handler *ConnectionHandler) GetRecommendations(ctx context.Context, request *pb.GetRecommendationsRequest) (*pb.GetRecommendationsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetRecommendations Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	Connections, err := handler.service.GetRecommendations(ctx, request.UserId)
	if err != nil {
		log.WithField("userId", request.UserId).Errorf("Cannot get connections: %v", err)
		return nil, err
	}
	response := &pb.GetRecommendationsResponse{
		Recommendations: make([]string, 0),
	}
	for _, Connection := range Connections {
		response.Recommendations = append(response.Recommendations, Connection)
	}
	return response, nil
}

func (handler *ConnectionHandler) BlockUser(ctx context.Context, request *pb.BlockUserRequest) (*pb.BlockUserResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "BlockUser Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	success, err := handler.service.BlockUser(ctx, ctx.Value("userId").(string), request.UserId)
	if err != nil {
		log.WithField("userId", request.UserId).Errorf("Cannot block user: %v", err)
		return nil, err
	}
	log.WithField("userId", request.UserId).Infof("User blocked")
	return &pb.BlockUserResponse{Success: success}, nil
}

func (handler *ConnectionHandler) GetBlockedUsers(ctx context.Context, request *pb.GetBlockedUsersRequest) (*pb.GetBlockedUsersResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetBlockedUsers Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	users, err := handler.service.GetBlockedUsers(ctx, request.UserId)
	if err != nil {
		log.WithField("userId", request.UserId).Errorf("Cannot get blocked users: %v", err)
		return nil, err
	}
	response := &pb.GetBlockedUsersResponse{
		BlockedUsers: make([]string, 0),
	}
	for _, user := range users {
		response.BlockedUsers = append(response.BlockedUsers, user)
	}
	return response, nil
}

func (handler *ConnectionHandler) GetBlockers(ctx context.Context, request *pb.GetBlockersRequest) (*pb.GetBlockersResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetBlockers Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	users, err := handler.service.GetBlockers(ctx, request.UserId)
	if err != nil {
		log.WithField("userId", request.UserId).Errorf("Cannot get blockers: %v", err)
		return nil, err
	}
	response := &pb.GetBlockersResponse{
		Blockers: make([]string, 0),
	}
	for _, user := range users {
		response.Blockers = append(response.Blockers, user)
	}
	return response, nil
}

func (handler *ConnectionHandler) UnblockUser(ctx context.Context, request *pb.UnblockUserRequest) (*pb.UnblockUserResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "UnblockUser Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	success, err := handler.service.UnblockUser(ctx, ctx.Value("userId").(string), request.UserId)
	if err != nil {
		log.WithField("userId", request.UserId).Errorf("Cannot unblock user: %v", err)
		return nil, err
	}
	log.WithField("userId", request.UserId).Infof("User unblocked")
	return &pb.UnblockUserResponse{
		Success: success,
	}, nil
}

func (handler *ConnectionHandler) GetConnection(ctx context.Context, request *pb.GetConnectionRequest) (*pb.GetConnectionResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetConnection Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	connection, err := handler.service.GetConnection(ctx, request.User1Id, request.User2Id)
	if err != nil {
		log.WithField("userId", request.User1Id).Errorf("Cannot get connection: %v", err)
		return nil, err
	}
	if connection == nil {
		return &pb.GetConnectionResponse{Connection: nil}, nil
	}
	response := &pb.GetConnectionResponse{
		Connection: mapConnectionToPb(connection),
	}

	return response, nil
}

func (handler *ConnectionHandler) GetLogs(ctx context.Context, request *pb.GetLogsRequest) (*pb.GetLogsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetLogs Handler")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	logs, err := handler.service.GetLogs(ctx)
	if err != nil {
		log.Errorf("GLF")
		return nil, err
	}
	pbLogs := make([]*pb.Log, len(logs))
	for i, log := range logs {
		pbLogs[i] = &pb.Log{
			Time:        timestamppb.New(log.Time),
			Level:       log.Level,
			Service:     "Connection service",
			Msg:         log.Msg,
			FullContent: log.FullContent,
		}
	}
	log.Info("GLD")
	return &pb.GetLogsResponse{Logs: pbLogs}, nil
}
