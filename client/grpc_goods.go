package client

import (
	goods "github.com/lizhenghan-cn/goods_service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func GoodsGrpc(address string) goods.GoodsClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return goods.NewGoodsClient(conn)

}
