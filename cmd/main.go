package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/pkg/naming"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/go-kratos/kratos/pkg/naming/etcd"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden/resolver"

	"go.etcd.io/etcd/clientv3"
	"kratos-risk/internal/di"
)

func main() {
	flag.Parse()
	log.Init(&log.Config{Stdout: true}) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("kratos-risk start")
	paladin.Init()
	discoveryInit()
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("kratos-risk exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func discoveryInit() {
	builder, err := etcd.New(&clientv3.Config{Endpoints: []string{"http://localhost:2379", "http://localhost:2380"}})
	if err != nil {
		panic(err)
	}
	instance := &naming.Instance{
		AppID:    "kratos-risk",
		Hostname: "",
		Addrs:    []string{"http://localhost:9000"},
	}
	cancelFunc, err := builder.Register(context.Background(), instance)
	if err != nil {
		cancelFunc()
		panic(err)
	}
	resolver.Register(builder)
}
