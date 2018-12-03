package main

import (
	"fmt"
	"richie.com/richie/learn_golang/crawler_distributed/config"
	"richie.com/richie/learn_golang/crawler_distributed/rpcsupport"
	"richie.com/richie/learn_golang/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRPC(host, worker.CrawlService{})

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1361568533",
		Parser: worker.SerializedParse{
			Name: config.ParseProfile,
			Args: "领悟",
		},
	}
	var result worker.ParseResult

	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil{
		t.Error(err)
	}else {
		fmt.Println(result)
	}
}
