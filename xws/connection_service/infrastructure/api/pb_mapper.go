package api

import (
	pb "dislinkt/common/proto/connection_service"
	pbPost "dislinkt/common/proto/post_service"
	"dislinkt/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

func mapConnectionToPb(connection *domain.Connection) *pb.Connection {
	return &pb.Connection{
		Id:         strconv.Itoa(connection.Id),
		IssuerId:   connection.IssuerId,
		SubjectId:  connection.SubjectId,
		Date:       timestamppb.New(connection.Date),
		IsApproved: connection.IsApproved,
	}
}

func mapConnectionToPostConnectionPb(connection *domain.Connection) *pbPost.Connection {
	return &pbPost.Connection{
		Id:        strconv.Itoa(connection.Id),
		IssuerId:  connection.IssuerId,
		SubjectId: connection.SubjectId,
	}
}

func mapPbToConnection(pbConnection *pb.Connection) *domain.Connection {
	id, err := strconv.Atoi(pbConnection.Id)
	if err != nil {
		id = 0
	}
	return &domain.Connection{
		Id:         id,
		IssuerId:   pbConnection.IssuerId,
		SubjectId:  pbConnection.SubjectId,
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
