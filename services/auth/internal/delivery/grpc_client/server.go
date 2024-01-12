package grpc_client

import "google.golang.org/grpc"

type ServerClient struct {
	UserClient         *grpc.ClientConn
	NotificationClient *grpc.ClientConn
}
