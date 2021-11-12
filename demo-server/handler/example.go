package handler

import (
	"context"
	"example/demo-server/domain/model"
	"example/demo-server/domain/service"
	"io"
	"time"

	log "go-micro.dev/v4/logger"

	pb "example/demo-server/proto"
)

type Example struct{
	UserService service.IUserService
}

func (e *Example) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received Example.Call request: %v", req)
	e.UserService.AddUser(&model.User{
		Name:         "Jinzhi",
	})
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Example) ClientStream(ctx context.Context, stream pb.Example_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *Example) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Example_ServerStreamStream) error {
	log.Infof("Received Example.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *Example) BidiStream(ctx context.Context, stream pb.Example_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
