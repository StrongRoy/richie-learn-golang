package client

import (
	"net/rpc"
	"richie.com/richie/learn_golang/crawler/engine"
	"richie.com/richie/learn_golang/crawler_distributed/config"
	"richie.com/richie/learn_golang/crawler_distributed/worker"
)

func CreateProcessor(clients chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (
		engine.ParseResult, error) {

		sReq := worker.SeriazeRequest(req)

		var sResult worker.ParseResult
		client := <-clients
		err := client.Call(config.CrawlServiceRpc,
			sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil

	}
}
