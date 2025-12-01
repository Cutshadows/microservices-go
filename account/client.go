package account

import (
	"github.com/Cutshadows/microservices-go/account/pb/github.com/Cutshadows/microservices-go/account/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.AccountServiceClient
}
