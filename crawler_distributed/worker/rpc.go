package worker

import "richie.com/richie/learn_golang/crawler/engine"

type CrawlService struct{}

func (CrawlService) Process(
	req Request, result *ParseResult) error {
	engineReq,err := DeserializeRequest(req)
	if err != nil{
		return  err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil{
		return  err
	}
	*result = SeriazeResult(engineResult)
	return nil

}
