package api

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const AppID = "kratos-risk"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (DemoClient, error) {
	client := warden.NewClient(cfg, opts...)
	fmt.Println("------------------------", fmt.Sprintf("etcd://default/kratos_etcd/%s/", AppID))
	cc, err := client.Dial(context.Background(), fmt.Sprintf("etcd://default/kratos_etcd/%s/", AppID))
	if err != nil {
		return nil, err
	}
	return NewDemoClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm api.proto
