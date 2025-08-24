package grpc_service

import (
	goods "github.com/lizhenghan-cn/goods_grpc/proto"
	"github.com/lizhenghans/grpc_service/client"
	"log"
)

type GrpcClient struct {
	Address string
}

func NewGrpcClient(address string) *GrpcClient {
	return &GrpcClient{Address: address}
}

func (g *GrpcClient) NewGrpcApi() goods.GoodsClient {
	log.Printf("GrpcAPi服务启动")
	return client.GoodsGrpcApi(g.Address)
}

func (g *GrpcClient) NewGrpcService() client.Server {
	log.Println("GrpcService启动")
	return client.GoodsGrpcService(g.Address)
}
