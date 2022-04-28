package initialize

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"liuqichat/api"
	"liuqichat/global"
	"liuqichat/pkg/proto"
	"net"
)

func InitGRPC() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v",
			global.GRPCInfo.Port,
		))
	if err != nil {
		panic(fmt.Errorf("启动 grpc 服务失败 err: %v", err))
	}

	grpcServer := grpc.NewServer()
	// 注册 grpc 服务到 redis 中
	global.RedisDB.SAdd(context.Background(), "grpcIp", global.GRPCInfo.Host+":"+global.GRPCInfo.Port)
	fmt.Println("grpc 启动成功 注册信息：grpcIp->", global.GRPCInfo.Host+":"+global.GRPCInfo.Port)
	proto.RegisterChatServer(grpcServer, &api.Chat{})
	if err = grpcServer.Serve(lis); err != nil {
		panic(fmt.Errorf("启动 grpc 服务失败 err: %v", err))
	}
}
