package grpc

import "google.golang.org/grpc"

func NewClient(conn *grpc.ClientConn) XephonkClient {
	return NewXephonkClient(conn)
}
