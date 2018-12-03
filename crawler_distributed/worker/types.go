package worker

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"richie.com/richie/learn_golang/crawler/engine"
	"richie.com/richie/learn_golang/crawler/zhenai/parser"
	"richie.com/richie/learn_golang/crawler_distributed/config"
)

type SerializedParse struct {
	Name  string
	Args  interface{}
	Args2 interface{}
}

// {"ParseCityList",nil}
// {"ParseProfile",userName,userGender}

type Request struct {
	Url    string
	Parser SerializedParse
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SeriazeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParse{
			Name: name,
			Args: args,
		},
	}
}

func SeriazeResult(
	r engine.ParseResult) ParseResult {

	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(
			result.Requests,
			SeriazeRequest(req))
	}
	return result
}

func DeserializeRequest(
	r Request) (engine.Request, error) {
	parser, err := deserializeParse(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		enginReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserizlizing "+
				"request: %v", err)
			continue
		}
		result.Requests = append(
			result.Requests,
			enginReq, )
	}
	return result
}

func deserializeParse(p SerializedParse) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(
			parser.ParseCityList, config.ParseCityList), nil

	case config.ParseCity:
		return engine.NewFuncParser(
			parser.ParseCity, config.ParseCity), nil

	case config.NilParser:
		return engine.NilParser{}, nil

	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {

			return parser.NewProfileParser(userName, "男士"), nil
		} else {
			return nil, fmt.Errorf("invalid "+
				"arg: %v", p.Args)
		}

	default:
		return nil, errors.New(
			"unknown parser name")

	}
}
