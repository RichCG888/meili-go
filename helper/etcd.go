package helper

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/client/v3"
	"meili-api/common"
	"time"
)

var (
	con, cancel = context.WithTimeout(context.Background(), 3*time.Second)
)

func init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Info("连接etcd失败, err:%v\n", err)
		return
	}
	defer func(cli *clientv3.Client) {
		err := cli.Close()
		if err != nil {
			log.Info("断开etcd失败, err:%v\n", err)
			panic(err)
		}
	}(cli)
	url, err := cli.KV.Get(con, "url")
	key, err := cli.KV.Get(con, "key")
	cancel()
	if err != nil {
		log.Info("获取值失败, err:%v\n", err)
	}
	log.Info("etcd初始化配置成功")
	meiliConfig := common.MeiliConfig{
		Url:    string(url.Kvs[0].Value),
		ApiKey: string(key.Kvs[0].Value),
	}
	common.Init(&meiliConfig)

}
