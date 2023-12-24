package adder

import (
	"context"

	trainv1 "github.com/Shuv1Wolf/train_protos/gen/go"
)

type GRPCServer struct {
	trainv1.UnimplementedAdderServer
}

func (s *GRPCServer) Add(ctx context.Context, req *trainv1.AddRequest) (*trainv1.AddResponse, error) {
	return &trainv1.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
