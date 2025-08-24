package client

import (
	"flag"
	goods "github.com/lizhenghan-cn/goods_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

func GoodsGrpcApi(address string) goods.GoodsClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return goods.NewGoodsClient(conn)

}

type Server struct {
	goods.UnimplementedGoodsServer
}

func GoodsGrpcService(address string) Server {
	flag.Parse()
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	goods.RegisterGoodsServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return Server{}
}
