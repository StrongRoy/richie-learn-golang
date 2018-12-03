package main

import (
	"richie.com/richie/learn_golang/crawler/engine"
	"richie.com/richie/learn_golang/crawler/persist"
	"richie.com/richie/learn_golang/crawler/scheduler"
	"richie.com/richie/learn_golang/crawler/zhenai/parser"
	"richie.com/richie/learn_golang/crawler_distributed/config"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(itemChan)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
