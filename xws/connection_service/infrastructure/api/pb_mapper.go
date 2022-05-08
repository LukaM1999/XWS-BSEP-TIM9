package api

import (
	pb "dislinkt/common/proto/connection_service"
	pbPost "dislinkt/common/proto/post_service"
	"dislinkt/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapConnectionToPb(connection *domain.Connection) *pb.Connection {
	return &pb.Connection{
		Id:         connection.Id.Hex(),
		IssuerId:   connection.IssuerId.Hex(),
		SubjectId:  connection.SubjectId.Hex(),
		Date:       timestamppb.New(connection.Date),
		IsApproved: connection.IsApproved,
	}
}

func mapConnectionToPostConnectionPb(connection *domain.Connection) *pbPost.Connection {
	return &pbPost.Connection{
		Id:        connection.Id.Hex(),
		IssuerId:  connection.IssuerId.Hex(),
		SubjectId: connection.SubjectId.Hex(),
	}
}

func mapPbToConnection(pbConnection *pb.Connection) *domain.Connection {
	return &domain.Connection{
		Id:         getObjectId(pbConnection.Id),
		IssuerId:   getObjectId(pbConnection.IssuerId),
		SubjectId:  getObjectId(pbConnection.SubjectId),
		Date:       pbConnection.Date.AsTime(),
		IsApproved: pbConnection.IsApproved,
	}
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
