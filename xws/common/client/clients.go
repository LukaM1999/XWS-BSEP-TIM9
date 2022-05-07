package client

import (
	"context"
	"dislinkt/common/auth"
	pbComment "dislinkt/common/proto/comment_service"
	pbConnection "dislinkt/common/proto/connection_service"
	pbPost "dislinkt/common/proto/post_service"
	pbProfile "dislinkt/common/proto/profile_service"
	pbReaction "dislinkt/common/proto/reaction_service"
	pbSecurity "dislinkt/common/proto/security_service"
	"time"

	"google.golang.org/grpc"
)

func NewPostClient(address string) (pbPost.PostServiceClient, error) {
	tlsCredentials, err := auth.LoadTLSClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbPost.NewPostServiceClient(conn)
	return client, nil
}

func NewProfileClient(address string) (pbProfile.ProfileServiceClient, error) {
	tlsCredentials, err := auth.LoadTLSClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbProfile.NewProfileServiceClient(conn)
	return client, nil
}

func NewConnectionClient(address string) (pbConnection.ConnectionServiceClient, error) {
	tlsCredentials, err := auth.LoadTLSClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbConnection.NewConnectionServiceClient(conn)
	return client, nil
}

func NewSecurityClient(address string) (pbSecurity.SecurityServiceClient, error) {
	tlsCredentials, err := auth.LoadTLSClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbSecurity.NewSecurityServiceClient(conn)
	return client, nil
}

func NewReactionClient(address string) (pbReaction.ReactionServiceClient, error) {
	tlsCredentials, err := auth.LoadTLSClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbReaction.NewReactionServiceClient(conn)
	return client, nil
}

func NewCommentClient(address string) (pbComment.CommentServiceClient, error) {
	tlsCredentials, err := auth.LoadTLSClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbComment.NewCommentServiceClient(conn)
	return client, nil
}
