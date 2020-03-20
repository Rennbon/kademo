package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.etcd.io/etcd/clientv3"

	"google.golang.org/grpc/grpclog"
)

func main() {
	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: time.Minute,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close() // make sure to close the client

	_, err = cli.Put(context.TODO(), "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}
	v, err := cli.Get(context.Background(), "foo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v.Count, v.Header, v.Kvs, v.More, v.OpResponse())
}
