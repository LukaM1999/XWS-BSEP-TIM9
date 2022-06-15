package api

import (
	"context"
	"dislinkt/common/loggers"
	pb "dislinkt/common/proto/connection_service"
	pbPost "dislinkt/common/proto/post_service"
	"dislinkt/connection_service/application"
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
	Connections, err := handler.service.Get(request.UserId)
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
	connection := mapPbToConnection(request.Connection)
	newConnection, err := handler.service.Create(connection)
	if err != nil {
		log.Errorf("Cannot create connection: %v", err)
		return nil, err
	}
	if newConnection.IsApproved {
		_, err = handler.postClient.CreateConnection(context.TODO(), &pbPost.CreateConnectionRequest{
			Connection: mapConnectionToPostConnectionPb(newConnection),
		})
		if err != nil {
			log.Errorf("Cannot create connection: %v", err)
			handler.service.Delete(newConnection.Id.Hex())
			return nil, err
		}
	}
	log.Info("Connection created")
	return &pb.CreateResponse{
		Connection: mapConnectionToPb(newConnection),
	}, nil
}

func (handler *ConnectionHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := handler.service.Delete(request.Id)
	if err != nil {
		log.Errorf("Cannot delete connection: %v", err)
		return nil, err
	}
	handler.postClient.DeleteConnection(context.TODO(), &pbPost.DeleteRequest{Id: request.Id})
	log.Info("Connection deleted")
	return &pb.DeleteResponse{}, nil
}

func (handler *ConnectionHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	connection, err := handler.service.Update(request.Id)
	if err != nil {
		log.WithField("connectionId", request.Id).Errorf("Cannot update connection: %v", err)
		return nil, err
	}
	if connection.IsApproved {
		_, err = handler.postClient.CreateConnection(context.TODO(), &pbPost.CreateConnectionRequest{
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
