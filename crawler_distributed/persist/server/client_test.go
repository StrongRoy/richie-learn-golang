package main

import (
	"richie.com/richie/learn_golang/crawler/engine"
	"richie.com/richie/learn_golang/crawler/model"
	"richie.com/richie/learn_golang/crawler_distributed/config"
	"richie.com/richie/learn_golang/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	// start ItemSaverServer
	go serverRpc(host, "test1")

	time.Sleep(time.Second)
	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call Save
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1361568533",
		Type: "zhenai",
		Id:   "1361568533",
		Payload: model.Profile{
			Age:        29,
			Height:     163,
			Weight:     60,
			Name:       "Need阳光",
			Gender:     "女",
			Income:     "3-5千",
			Marriage:   "未婚",
			Education:  "大学本科",
			WorkPlace:  "阿坝汶川",
			Occupation: "政府机构",
			Xinzuo:     "魔羯座(12.22-01.19)",
		},
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
