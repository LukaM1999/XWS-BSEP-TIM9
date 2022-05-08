package api

import (
	"context"
	pb "dislinkt/common/proto/connection_service"
	pbPost "dislinkt/common/proto/post_service"
	"dislinkt/connection_service/application"
)

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
		handler.service.Delete(newConnection.Id.Hex())
		return nil, err
	}
	if newConnection.IsApproved {
		_, err = handler.postClient.CreateConnection(context.TODO(), &pbPost.CreateConnectionRequest{
			Connection: mapConnectionToPostConnectionPb(newConnection),
		})
		if err != nil {
			handler.service.Delete(newConnection.Id.Hex())
			return nil, err
		}
	}
	return &pb.CreateResponse{
		Connection: mapConnectionToPb(newConnection),
	}, nil
}

func (handler *ConnectionHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := handler.service.Delete(request.Id)
	if err != nil {
		return nil, err
	}
	handler.postClient.DeleteConnection(context.TODO(), &pbPost.DeleteRequest{Id: request.Id})
	return &pb.DeleteResponse{}, nil
}

func (handler *ConnectionHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	connection, err := handler.service.Update(request.Id)
	if err != nil {
		return nil, err
	}
	if connection.IsApproved {
		_, err = handler.postClient.CreateConnection(context.TODO(), &pbPost.CreateConnectionRequest{
			Connection: mapConnectionToPostConnectionPb(connection),
		})
		if err != nil {
			return nil, err
		}
	}
	return &pb.UpdateResponse{
		Connection: mapConnectionToPb(connection),
	}, nil
}
