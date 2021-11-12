package main

import (
	"example/demo-server/domain/repository"
	service2 "example/demo-server/domain/service"
	"example/demo-server/handler"
	pb "example/demo-server/proto"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	service = "demo-server"
	version = "latest"
)

func main() {
	//注册中心
	consul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.100.80:8500",
		}
	})
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		// 添加 consul  注册中心
		micro.Registry(consul),
		//这里设置地址和需要暴露的端口
		micro.Address("0.0.0.0:10000"),
	)
	srv.Init()

	dsn := "host=192.168.100.80 user=postgres password=Zdz2020! dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	rp := repository.NewUserRepository(db)
	rp.InitTable()
	//Initialise service


	userService:=service2.NewUserService(repository.NewUserRepository(db))
	// Register handler
	pb.RegisterExampleHandler(srv.Server(), &handler.Example{UserService: userService})

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
