package main

//import (
//	"context"
//	"fmt"
//	"github.com/go-kratos/kratos/pkg/log"
//	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
//	xtime "github.com/go-kratos/kratos/pkg/time"
//	"kratos-risk/api"
//	"time"
//)
//
//func main() {
//	cli ,err := api.NewClient(&warden.ClientConfig{Timeout: xtime.Duration(5*time.Second)})
//	if err != nil {
//		panic(err)
//	}
//	req := &api.HelloReq{Name: "zhangsan"}
//	resp,err := cli.SayHelloURL(context.Background(),req)
//	if err!= nil {
//		log.Debug("err；",err)
//	}
//
//	fmt.Println("999999999",resp.Content)
//}
