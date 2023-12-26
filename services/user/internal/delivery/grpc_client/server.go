package grpc_client

import "google.golang.org/grpc"

type ServerClient struct {
	AuthClient *grpc.ClientConn
}
