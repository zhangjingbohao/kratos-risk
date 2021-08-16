package test

import (
	"context"
	"fmt"
	"kratos-risk/api"
	"testing"
)

func TestHello(t *testing.T) {
	client, err := api.NewClient(nil)
	if err != nil {
		fmt.Println("9999999999999999999", err)
		panic(err)
	}
	req := &api.HelloReq{Name: "张三"}
	resp, err := client.SayHelloURL(context.Background(), req)
	if err != nil {
		fmt.Println("------------------------", err)
		panic(err)
	}
	fmt.Println("+++++++++++++++++++++++++", resp)
}
