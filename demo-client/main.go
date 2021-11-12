package main

import (
	"context"
	"database/sql"
	pb "example/demo-server/proto"
	source "github.com/asim/go-micro/plugins/config/source/consul/v4"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/config"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"strconv"
	"time"
)

var (
	service = "demo-server"
	version = "latest"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
/*	dsn := "host=192.168.100.80 user=postgres password=Zdz2020! dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
		return
	}
	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	err = db.Create(&users).Error
	if err != nil {
		log.Fatal(err)
		return
	}
	consulConfig, err := GetConsulConfig("192.168.100.80", 8500, "")
	if err != nil {
		log.Fatal(err)
		return
	}
	mysqlInfo := GetMysqlFromConsul(consulConfig, "mysql")
	log.Info(mysqlInfo)*/

	//注册中心
	consul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.100.80:8500",
		}
	})
	// Create service
	srv := micro.NewService(
		micro.Name("test-client"),
		micro.Registry(consul),
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)
	srv.Init()

	// Create client
	c := pb.NewExampleService(service, srv.Client())

	for {
		// Call service
		rsp, err := c.Call(context.Background(), &pb.CallRequest{Name: "John"})
		if err != nil {
			log.Fatal(err)
		}

		log.Info(rsp)

		time.Sleep(1 * time.Second)
	}
}

//设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := source.NewSource(
		// 设置配置中心地址
		source.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		// 设置前缀 不设置默认前缀 /micro/config
		source.WithPrefix(prefix),
		//是否移除前缀，这里是设置为true，表示可以不带前缀直接获取对应配置
		source.StripPrefix(true),
	)

	// 配置初始化
	config, err := config.NewConfig()
	if err != nil {
		return config, err
	}
	// 加载配置文件
	err = config.Load(consulSource)
	return config, err
}

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

// 获取mysql配置
func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}
