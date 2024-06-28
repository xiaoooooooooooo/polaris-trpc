package main

import (
	"context"
	"polaris-trpc/edu/imau/pb"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/errs"
	"trpc.group/trpc-go/trpc-go/log"
	_ "trpc.group/trpc-go/trpc-naming-polarismesh/registry"
)

func main() {
	trpcServer := trpc.NewServer()
	pb.RegisterUserServiceService(trpcServer, &UserServiceImpl{})
	if err := trpcServer.Serve(); err != nil {
		log.Fatalf("rpc服务启动失败 err: %v", err)
	}
}

type UserServiceImpl struct{}

func (*UserServiceImpl) GetUserById(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	log.Infof("9090 收到请求")

	if req.GetId() == 1 {
		return &pb.UserResponse{
			Id:   1,
			Name: "张三",
		}, nil
	} else {
		// 返回自定义错误
		return nil, errs.New(10000, "用户不存在")
	}
}
