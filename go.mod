module example

go 1.16

require (
	github.com/asim/go-micro/plugins/config/source/consul/v4 v4.0.0-20211111140334-799b8d6a6559
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20211111140334-799b8d6a6559
	github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v4 v4.0.0-20211111140334-799b8d6a6559
	go-micro.dev/v4 v4.2.1
	google.golang.org/protobuf v1.26.0
	gorm.io/driver/postgres v1.2.1
	gorm.io/gorm v1.22.2
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
