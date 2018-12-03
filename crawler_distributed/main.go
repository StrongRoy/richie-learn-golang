package main

import (
	"flag"
	"fmt"
	"net/rpc"
	"richie.com/richie/learn_golang/crawler/engine"
	"richie.com/richie/learn_golang/crawler/scheduler"
	"richie.com/richie/learn_golang/crawler/zhenai/parser"
	"richie.com/richie/learn_golang/crawler_distributed/config"
	itemsaver "richie.com/richie/learn_golang/crawler_distributed/persist/client"
	"richie.com/richie/learn_golang/crawler_distributed/rpcsupport"
	worker "richie.com/richie/learn_golang/crawler_distributed/worker/client"
	"strings"
)

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "item saver host")
	workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	if *itemSaverHost == "" || *workerHosts == "" {
		fmt.Println("must specify itemsaver_host and worker_hosts")
		return
	}

	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(
		*workerHosts, ","))
	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, host := range hosts {
		client, err := rpcsupport.NewClient(host)
		if err == nil {
			clients = append(clients, client)
			fmt.Printf("Connected to %s.", host)
		} else {
			fmt.Printf("Error connecting to %s: %v.", host, err)
		}
	}

	out := make(chan *rpc.Client)

	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out

}
