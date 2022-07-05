package api

import (
	pb "dislinkt/common/proto/post_service"
	"dislinkt/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapPostToPb(post *domain.Post) *pb.Post {
	pbProfile := &pb.Profile{
		Id:        post.Profile.Id.Hex(),
		FirstName: post.Profile.FirstName,
		LastName:  post.Profile.LastName,
	}
	pbContent := &pb.Content{
		Text:  post.Content.Text,
		Image: post.Content.Image,
		Links: make([]string, 0),
	}
	pbPost := &pb.Post{
		Id:        post.Id.Hex(),
		Profile:   pbProfile,
		CreatedAt: timestamppb.New(post.CreatedAt),
		Content:   pbContent,
	}

	for _, link := range post.Content.Links {
		pbContent.Links = append(pbContent.Links, link)
	}

	return pbPost
}

func mapPbToPost(pbPost *pb.Post) *domain.Post {
	profile := &domain.Profile{
		Id:        getObjectId(pbPost.Profile.Id),
		FirstName: pbPost.Profile.FirstName,
		LastName:  pbPost.Profile.LastName,
	}
	content := &domain.Content{
		Text:  pbPost.Content.Text,
		Image: pbPost.Content.Image,
		Links: make([]string, 0),
	}
	for _, link := range pbPost.Content.Links {
		content.Links = append(content.Links, link)
	}
	post := &domain.Post{
		Id:        getObjectId(pbPost.Id),
		Profile:   *profile,
		CreatedAt: pbPost.CreatedAt.AsTime(),
		Content:   *content,
	}

	return post
}

func mapConnectionToPb(connection *domain.Connection) *pb.Connection {
	return &pb.Connection{
		Id:        connection.Id.Hex(),
		IssuerId:  connection.IssuerId.Hex(),
		SubjectId: connection.SubjectId.Hex(),
	}
}

func mapPbToConnection(pbConnection *pb.Connection) *domain.Connection {
	return &domain.Connection{
		Id:        getObjectId(pbConnection.Id),
		IssuerId:  getObjectId(pbConnection.IssuerId),
		SubjectId: getObjectId(pbConnection.SubjectId),
	}
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
